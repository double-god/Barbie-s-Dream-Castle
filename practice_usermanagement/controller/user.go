//用户控制器
/*承接路由，协调业务逻辑与数据存取层，
解析HTTP请求（参数，Body),调用service层去干活，从service拿到结果，调用util.Respond返回json相应*/
package controller

import (
	"net/http"
	"practice_usermanagement/model"
	"practice_usermanagement/pkg/database"
	"practice_usermanagement/pkg/e"
	"practice_usermanagement/pkg/util"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Register godoc
// @Summary      用户注册
// @Description  用户注册接口
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        user  body      object{username=string,password=string}  true  "用户名和密码"
// @Success      200   {object}  util.Response
// @Router       /api/register [post]
func Register(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Email    string `json:"email"` //非必须
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		util.RespondError(c, http.StatusBadRequest, e.InvalidParams)
		return
	}
	var user model.User
	if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err == nil {
		util.Respond(c, http.StatusBadRequest, e.ERROR, gin.H{"error": "用户名已存在"})
		return
	}
	newUser := model.User{
		Username: req.Username,
		Email:    req.Email,
	}
	if err := newUser.SetPassword(req.Password); err != nil {
		util.RespondError(c, http.StatusInternalServerError, e.ERROR)
		return
	}
	if err := database.DB.Create(&newUser).Error; err != nil {
		util.RespondError(c, http.StatusInternalServerError, e.ERROR)
		return
	}
	util.RespondSuccess(c, gin.H{})
}

// Login godoc
// @Summary      用户登录
// @Description  用户登录接口
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        user  body      object{username=string,password=string}  true  "用户名和密码"
// @Success      200   {object}  util.Response
// @Router       /api/login [post]
func Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		util.RespondError(c, http.StatusBadRequest, e.InvalidParams)
		return
	}
	var user model.User
	if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			util.Respond(c, http.StatusUnauthorized, e.ERROR, gin.H{"error": "incorrect username or password"})
			return
		}
		util.RespondError(c, http.StatusInternalServerError, e.ERROR)
		return
	}
	if !user.CheckPassword(req.Password) {
		util.Respond(c, http.StatusUnauthorized, e.ERROR, gin.H{"error": "incorrect username or password"})
		return
	}
	token, err := util.GenerateToken(user.ID, user.Username)
	if err != nil {
		util.RespondError(c, http.StatusInternalServerError, e.ErrorAuthToken)
		return
	}
	util.RespondSuccess(c, gin.H{"token": token})
}

// GetProfile godoc
// @Summary      获取用户资料
// @Description  获取用户资料接口
// @Tags         User
// @Produce      json
// @Param        Authorization  header    string  true  "User ID"
// @Success      200   {object}  util.Response{data=model.User}
// @Router       /api/user/{uid} [get]
// @Security     ApiKeyAuth
func GetProfile(c *gin.Context) {
	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		util.RespondError(c, http.StatusBadRequest, e.InvalidParams)
		return
	}
	var user model.User
	if err := database.DB.Select("username", "email", "bio").First(&user, uid).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			util.Respond(c, http.StatusNotFound, e.ERROR, gin.H{"error": "user not found"})
			return
		}
		util.RespondError(c, http.StatusInternalServerError, e.ERROR)
		return
	}
	util.RespondSuccess(c, user)
}

// UpdateUser godoc
// @Summary      更新当前用户信息
// @Description  更新当前登录用户的信息
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        user  body      object{email=string,bio=string}  true  "要更新的信息"
// @Success      200   {object}  util.Response
// @Security     ApiKeyAuth
// @Router       /api/user [put]
func UpdateUser(c *gin.Context) {
	var req struct {
		Email string `json:"email"`
		Bio   string `json:"bio"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		util.RespondError(c, http.StatusBadRequest, e.InvalidParams)
		return
	}
	userID, _ := c.Get("userID")
	var user model.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		util.Respond(c, http.StatusNotFound, e.ERROR, gin.H{"error": "user not found"})
		return
	}

	updates := model.User{
		Email: req.Email,
		Bio:   req.Bio,
	}
	if err := database.DB.Model(&user).Updates(updates).Error; err != nil {
		util.RespondError(c, http.StatusInternalServerError, e.ERROR)
		return
	}
	util.RespondSuccess(c, gin.H{})
}

// GetUser godoc
// @Summary      获取用户信息
// @Description  根据用户ID获取用户信息
// @Tags         User
// @Produce      json
// @Param        uid  path      int  true  "User ID"
// @Success      200  {object}  util.Response{data=model.User}
// @Security     ApiKeyAuth
// @Router       /api/users/{uid} [get]
func GetUser(c *gin.Context) {
	userID := c.Param("uid")
	var user model.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		util.Respond(c, http.StatusNotFound, e.ERROR, gin.H{"error": "user not found"})
		return
	}
	util.RespondSuccess(c, gin.H{"user": user})
}

// SearchUsers godoc
// @Summary      搜索用户
// @Description  根据用户名搜索用户列表
// @Tags         User
// @Produce      json
// @Param        username query     string  false  "Username for searching"
// @Param        page     query     int     false  "Page number" default(1)
// @Param        size     query     int     false  "Page size"   default(10)
// @Success      200      {object}  util.Response{data=[]model.User}
// @Security     ApiKeyAuth
// @Router       /api/user/search [get]
func SearchUsers(c *gin.Context) {
	username := c.Query("username")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	offset := (page - 1) * size
	var users []model.User
	query := database.DB.Select("id", "username", "bio")
	if username != "" {
		query = query.Where("username LIKE ?", "%"+username+"%")
	}
	if err := query.Offset(offset).Limit(size).Find(&users).Error; err != nil {
		util.RespondError(c, http.StatusInternalServerError, e.ERROR)
		return
	}
	util.RespondSuccess(c, users)
}

// controller/user_controller.go

// ... (你其他的 Register, Login 等函数) ...

// GetMe godoc
// @Summary      获取当前用户信息
// @Description  根据 Token 获取当前登录用户的信息
// @Tags         User
// @Produce      json
// @Success      200   {object}  util.Response{data=model.User}
// @Security     ApiKeyAuth
// @Router       /api/user/me [get]
func GetMe(c *gin.Context) {
	// 从 JWTAuth 中间件获取 user_id
	userID_raw, exists := c.Get("user_id")
	if !exists {
		util.RespondError(c, http.StatusUnauthorized, e.ErrorAuthCheckTokenFail)
		return
	}

	// userID 是 uint 类型，需要断言
	userID, ok := userID_raw.(uint)
	if !ok {
		util.RespondError(c, http.StatusUnauthorized, e.ERROR)
		return
	}

	var user model.User
	// 根据 userID 查询用户，不返回 PasswordHash
	if err := database.DB.Select("id", "username", "email", "bio", "created_at").First(&user, userID).Error; err != nil {
		util.Respond(c, http.StatusNotFound, e.ERROR, gin.H{"error": "User not found"})
		return
	}

	util.RespondSuccess(c, user)
}
