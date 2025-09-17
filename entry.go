package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

// UserData 结构体
type UserData struct {
	PhoneNumber      string
	VerificationCode string
	SendTime         time.Time
	DailyCount       int
	LastSendDate     time.Time
}

var Database = make(map[string]UserData)

func init() {
	Database = make(map[string]UserData)
	rand.Seed(time.Now().UnixNano())
}

func validatePhoneNumber(number string) bool {
	if len(number) != 11 || number[0] != '1' {
		return false
	}
	for _, r := range number {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func sendVerificationCode(number string) {
	data := Database[number]

	if !data.LastSendDate.IsZero() && time.Now().Day() != data.LastSendDate.Day() {
		data.DailyCount = 0
	}
	if data.DailyCount >= 5 {
		fmt.Println("您今天获取验证码次数已达上限，请明天再试")
		return
	}
	if !data.SendTime.IsZero() && time.Since(data.SendTime) < 60*time.Second {
		fmt.Println("一分钟内已获取验证码，无法重新获取")
		return
	}

	randomNumber := rand.Intn(900000) + 100000
	codeString := fmt.Sprintf("%d", randomNumber)

	newData := UserData{
		PhoneNumber:      number,
		VerificationCode: codeString,
		SendTime:         time.Now(),
		DailyCount:       data.DailyCount + 1,
		LastSendDate:     time.Now(),
	}
	Database[number] = newData

	fmt.Printf("验证码为：%s\n", codeString)
}

func main() {
	var phoneNumber string
	for {
		fmt.Println("请输入电话号码。")
		_, err := fmt.Scanln(&phoneNumber)
		if err != nil {
			fmt.Println("输入错误，请重试")
			continue
		}
		phoneNumber = strings.TrimSpace(phoneNumber)

		if !validatePhoneNumber(phoneNumber) {
			fmt.Println("登录失败，手机号格式不正确")
			continue
		}
		break
	}

	for {
		var choice string
		fmt.Println("1：输入验证码来登录哦 2：获取验证码")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("出错啦，再输一遍！")
			continue
		}
		choice = strings.TrimSpace(choice)

		if choice == "2" {
			sendVerificationCode(phoneNumber)
		} else if choice == "1" {
			var inputCode string
			fmt.Println("请输入你收到的验证码")
			_, err := fmt.Scanln(&inputCode)
			if err != nil {
				fmt.Println("错错错，请重试")
				continue
			}
			inputCode = strings.TrimSpace(inputCode)

			data := Database[phoneNumber]
			if data.SendTime.IsZero() {
				fmt.Println("输入无效，先去获取验证码！！！")
				continue
			}

			if time.Since(data.SendTime) > 5*time.Minute {
				delete(Database, phoneNumber)
				fmt.Println("验证码已过期，请重新获取")
				continue
			}

			if inputCode == data.VerificationCode {
				delete(Database, phoneNumber)
				fmt.Println("登录成功！完结撒花！")
				return
			} else {
				fmt.Println("验证码不对哦，请重试")
				continue
			}
		} else {
			fmt.Println("无效！输入1或2！！！")
			continue
		}
	}
}
