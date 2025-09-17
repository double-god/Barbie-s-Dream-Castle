const searchInput = document.getElementById('searchInput');
const searchIcon = document.getElementById('search-icon');
const lastSearchIcon = document.getElementById('last-search-icon');
const randomSearchIcon = document.getElementById('random-search-icon');
const searchHistoryDiv = document.getElementById('search-history');
const clockDiv = document.getElementById('clock');
    
// 预设词语
const presetQuery = "猜你喜欢：《小七和奶龙》"; 
    
// 随机搜索词语列表
const randomQueries = ["南大家园", "云家园", "家园工作室", "小家园传声机"];

// 搜索功能
function performSearch() {
    let query = searchInput.value.trim();
    if (!query) {
        query = presetQuery;
    }
    saveSearchHistory(query);

    const searchURL = `https://www.google.com/search?q=${encodeURIComponent(query)}`;

    // 将跳转放入一个微小的延迟执行要不然我存储来不及好奇怪啊
    setTimeout(() => {
        window.location.href = searchURL;
    }, 100); // 延迟100毫秒
}

// 把搜索记录保存到本地存储
function saveSearchHistory(query) {
    let history = JSON.parse(localStorage.getItem('searchHistory')) || [];
    if (history.indexOf(query) === -1) {
        history.unshift(query);
        if (history.length > 5) {
            history.pop();
        }
        localStorage.setItem('searchHistory', JSON.stringify(history));
        displaySearchHistory();
    }
}
    
// 显示搜索记录
function displaySearchHistory() {
    searchHistoryDiv.innerHTML = '';
    const history = JSON.parse(localStorage.getItem('searchHistory')) || [];
    history.forEach(item => {
        const div = document.createElement('div');
        div.className = 'history-item';
        div.textContent = item;
        div.addEventListener('click', () => {
            const searchURL = `https://www.google.com/search?q=${item}`;
            window.location.href = searchURL;
        });
        searchHistoryDiv.appendChild(div);
    });
}
    
// 时钟
function updateClock() {
    const now = new Date();
    const hours = now.getHours().toString().padStart(2, '0');
    const minutes = now.getMinutes().toString().padStart(2, '0');
    const seconds = now.getSeconds().toString().padStart(2, '0');
    clockDiv.textContent = `${hours}:${minutes}:${seconds}`;
}

// 事件监听
searchIcon.addEventListener('click', performSearch);
searchInput.addEventListener('keydown', (e) => {
    if (e.key === 'Enter') {
        performSearch();
    }
});

lastSearchIcon.addEventListener('click', () => {
    const history = JSON.parse(localStorage.getItem('searchHistory')) || [];
    if (history.length > 0) {
        searchInput.value = history[0];
    }
});

randomSearchIcon.addEventListener('click', () => {
    const randomIndex = Math.floor(Math.random() * randomQueries.length);
    searchInput.value = randomQueries[randomIndex];
});
searchInput.addEventListener('focus', () => {
    // 只有当有历史记录时才显示
    const history = JSON.parse(localStorage.getItem('searchHistory')) || [];
    if (history.length > 0) {
        searchHistoryDiv.classList.add('show');
    }
});

window.addEventListener('click', (e) => {
    // e.target 是用户点击的元素
    // .closest('.search-wrapper') 会查找被点击元素是否在 .search-wrapper 内部
    // 如果点击的地方不在 .search-wrapper 内部，就隐藏历史记录
    if (!e.target.closest('.search-wrapper')) {
        searchHistoryDiv.classList.remove('show');
    }
});

// 页面加载时执行
updateClock();
setInterval(updateClock, 1000);
displaySearchHistory();
