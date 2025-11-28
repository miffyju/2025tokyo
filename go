<!DOCTYPE html>
<html lang="zh-Hant-TW">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>我的旅遊行程表</title>
    <!-- Tailwind CSS -->
    <script src="https://cdn.tailwindcss.com"></script>
    <!-- FontAwesome Icons -->
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.4.0/css/all.min.css">
    <!-- Google Fonts: Noto Sans TC -->
    <link href="https://fonts.googleapis.com/css2?family=Noto+Sans+TC:wght@300;400;500;700&display=swap" rel="stylesheet">
    
    <style>
        body {
            font-family: 'Noto Sans TC', sans-serif;
            background-color: #f3f4f6;
        }
        .hide-scrollbar::-webkit-scrollbar {
            display: none;
        }
        .hide-scrollbar {
            -ms-overflow-style: none;
            scrollbar-width: none;
        }
        .glass-nav {
            background: rgba(255, 255, 255, 0.95);
            backdrop-filter: blur(10px);
            border-bottom: 1px solid rgba(0,0,0,0.05);
        }
        /* Modal 動畫 */
        .modal {
            transition: opacity 0.3s ease, visibility 0.3s ease;
        }
        .modal.hidden {
            opacity: 0;
            visibility: hidden;
            pointer-events: none;
        }
        .modal:not(.hidden) {
            opacity: 1;
            visibility: visible;
            pointer-events: auto;
        }
    </style>
    <script>
        tailwind.config = {
            theme: {
                extend: {
                    colors: {
                        travel: {
                            50: '#f0f9ff',
                            100: '#e0f2fe',
                            500: '#0ea5e9',
                            600: '#0284c7',
                            900: '#0c4a6e',
                        }
                    }
                }
            }
        }
    </script>
</head>
<body class="text-gray-800 pb-24">

    <!-- Header 區域 -->
    <header class="bg-gradient-to-br from-travel-600 to-travel-900 text-white pt-8 pb-16 px-6 relative overflow-hidden shadow-lg">
        <div class="absolute top-0 right-0 opacity-10 transform translate-x-10 -translate-y-10">
            <i class="fa-solid fa-plane-departure text-9xl"></i>
        </div>
        <div class="relative z-10 max-w-4xl mx-auto">
            <div class="flex justify-between items-start">
                <div id="headerInfo">
                    <h1 class="text-3xl font-bold mb-2">我的旅程 ✈️</h1>
                    <p class="text-travel-100 text-sm"><i class="fa-regular fa-calendar mr-2"></i>隨時出發</p>
                </div>
                <!-- 設定按鈕 -->
                <button onclick="toggleSettingsModal()" class="bg-white/20 backdrop-blur-sm rounded-full p-2 w-10 h-10 flex items-center justify-center hover:bg-white/30 transition-colors">
                    <i class="fa-solid fa-gear text-white"></i>
                </button>
            </div>
            
            <!-- 搜尋框 -->
            <div class="mt-6 relative">
                <input type="text" id="searchInput" placeholder="搜尋景點、美食..." class="w-full py-3 pl-10 pr-4 rounded-full text-gray-800 focus:outline-none focus:ring-2 focus:ring-travel-500 shadow-lg">
                <i class="fa-solid fa-search absolute left-4 top-3.5 text-gray-400"></i>
            </div>
        </div>
    </header>

    <!-- 天數選擇器 (Sticky) -->
    <nav class="glass-nav sticky top-0 z-40 px-4 py-3 shadow-sm overflow-x-auto hide-scrollbar">
        <div class="max-w-4xl mx-auto flex space-x-3" id="dayTabs">
            <!-- JS 動態生成按鈕 -->
        </div>
    </nav>

    <!-- 主要內容區 -->
    <main class="max-w-4xl mx-auto px-4 -mt-8 relative z-20" id="itineraryContainer">
        <!-- JS 動態生成行程卡片 -->
    </main>

    <!-- 設定 Modal -->
    <div id="settingsModal" class="modal hidden fixed inset-0 z-50 flex items-center justify-center bg-black/50 px-4">
        <div class="bg-white rounded-2xl w-full max-w-md p-6 shadow-2xl transform transition-all scale-100">
            <div class="flex justify-between items-center mb-4">
                <h2 class="text-xl font-bold text-gray-800">行程設定</h2>
                <button onclick="toggleSettingsModal()" class="text-gray-400 hover:text-gray-600">
                    <i class="fa-solid fa-times text-xl"></i>
                </button>
            </div>

            <div class="space-y-4">
                <div class="bg-blue-50 p-4 rounded-xl border border-blue-100">
                    <h3 class="font-bold text-blue-800 mb-2"><i class="fa-solid fa-file-csv mr-2"></i>匯入 Google Sheet</h3>
                    <p class="text-sm text-blue-600 mb-3">請先將您的 Google Sheet 下載為 <strong>CSV</strong> 格式，然後在此上傳。</p>
                    
                    <label class="block w-full">
                        <span class="sr-only">選擇 CSV 檔案</span>
                        <input type="file" id="csvInput" accept=".csv" class="block w-full text-sm text-gray-500
                            file:mr-4 file:py-2 file:px-4
                            file:rounded-full file:border-0
                            file:text-sm file:font-semibold
                            file:bg-travel-50 file:text-travel-700
                            hover:file:bg-travel-100
                        "/>
                    </label>
                    <div class="mt-2 text-xs text-gray-400">
                        欄位順序：天數 | 日期 | 時間 | 標題 | 類型 | 地點 | 價格 | 備註 | 連結
                    </div>
                </div>

                <div class="border-t border-gray-100 pt-4">
                    <button onclick="resetData()" class="w-full py-2 text-red-500 hover:bg-red-50 rounded-lg transition-colors text-sm font-medium">
                        <i class="fa-solid fa-trash-can mr-2"></i>清除資料並恢復範例
                    </button>
                </div>
            </div>
        </div>
    </div>

    <!-- 底部功能列 -->
    <div class="fixed bottom-0 left-0 w-full bg-white border-t border-gray-200 py-3 px-6 flex justify-around items-center z-40 text-xs text-gray-500 shadow-[0_-2px_10px_rgba(0,0,0,0.05)] md:hidden">
        <button class="flex flex-col items-center text-travel-600 font-bold" onclick="window.scrollTo({top:0, behavior:'smooth'})">
            <i class="fa-solid fa-list-ul text-xl mb-1"></i>
            行程
        </button>
        <button class="flex flex-col items-center hover:text-travel-600" onclick="alert('請點擊行程中的「導航」按鈕直接開啟 Google Maps')">
            <i class="fa-solid fa-map-location-dot text-xl mb-1"></i>
            地圖
        </button>
        <button class="flex flex-col items-center hover:text-travel-600" onclick="toggleSettingsModal()">
            <i class="fa-solid fa-gear text-xl mb-1"></i>
            設定
        </button>
    </div>

    <!-- JavaScript -->
    <script>
        // --- 1. 預設範例資料 ---
        const defaultData = [
            {
                day: 1,
                date: "Day 1",
                items: [
                    { time: "09:00", title: "範例：抵達機場", type: "transport", location: "機場", price: "", note: "點擊右上角設定匯入您的 CSV", link: "" },
                    { time: "12:00", title: "範例：午餐時間", type: "food", location: "市區", price: "¥1000", note: "", link: "" },
                    { time: "15:00", title: "範例：觀光景點", type: "sightseeing", location: "景點A", price: "免費", note: "記得拍照", link: "" }
                ]
            }
        ];

        // 全域變數
        let travelData = [];
        let currentDay = 1;

        // --- 2. 初始化與資料讀取 ---
        function init() {
            // 嘗試從 localStorage 讀取
            const savedData = localStorage.getItem('myTravelItinerary');
            if (savedData) {
                try {
                    travelData = JSON.parse(savedData);
                } catch (e) {
                    console.error("資料損毀，載入預設值");
                    travelData = defaultData;
                }
            } else {
                travelData = defaultData;
            }

            // 如果資料是空的或是舊格式，重置
            if (!Array.isArray(travelData) || travelData.length === 0) {
                travelData = defaultData;
            }

            // 確保 currentDay 有效
            currentDay = travelData[0].day;

            renderDayTabs();
            renderItinerary();
        }

        // --- 3. 輔助函式 ---
        function getTypeStyle(typeRaw) {
            // 模糊比對類型
            const type = typeRaw ? typeRaw.toLowerCase() : 'other';
            
            if (type.includes('食') || type.includes('餐') || type === 'food') 
                return { icon: 'fa-utensils', color: 'text-orange-500 bg-orange-50', border: 'border-orange-200' };
            
            if (type.includes('車') || type.includes('交通') || type.includes('機') || type === 'transport') 
                return { icon: 'fa-train-subway', color: 'text-blue-500 bg-blue-50', border: 'border-blue-200' };
            
            if (type.includes('景') || type.includes('遊') || type.includes('觀光') || type === 'sightseeing') 
                return { icon: 'fa-camera', color: 'text-emerald-500 bg-emerald-50', border: 'border-emerald-200' };
            
            if (type.includes('住') || type.includes('店') || type.includes('館') || type === 'stay') 
                return { icon: 'fa-bed', color: 'text-purple-500 bg-purple-50', border: 'border-purple-200' };
            
            if (type.includes('購') || type.includes('買') || type === 'shopping') 
                return { icon: 'fa-bag-shopping', color: 'text-pink-500 bg-pink-50', border: 'border-pink-200' };

            return { icon: 'fa-person-walking', color: 'text-gray-500 bg-gray-50', border: 'border-gray-200' };
        }

        // --- 4. 渲染函式 ---
        function renderDayTabs() {
            const container = document.getElementById('dayTabs');
            container.innerHTML = '';

            travelData.forEach(dayData => {
                const isActive = dayData.day == currentDay; // 寬鬆比對
                const btn = document.createElement('button');
                btn.className = `flex-shrink-0 px-5 py-2 rounded-full text-sm font-medium transition-all duration-300 whitespace-nowrap 
                    ${isActive ? 'bg-travel-600 text-white shadow-md transform scale-105' : 'bg-white text-gray-600 hover:bg-gray-100'}`;
                btn.innerHTML = `Day ${dayData.day} <span class="text-xs opacity-80 ml-1">${dayData.date}</span>`;
                btn.onclick = () => {
                    currentDay = dayData.day;
                    renderDayTabs();
                    renderItinerary();
                    window.scrollTo({ top: 0, behavior: 'smooth' });
                };
                container.appendChild(btn);
            });
        }

        function renderItinerary() {
            const container = document.getElementById('itineraryContainer');
            container.innerHTML = '';

            const dayData = travelData.find(d => d.day == currentDay);
            
            if (!dayData) return;

            // 更新 Header 標題日期
            const headerInfo = document.querySelector('#headerInfo p');
            if(headerInfo && travelData.length > 0) {
                const startDate = travelData[0].date;
                const endDate = travelData[travelData.length-1].date;
                headerInfo.innerHTML = `<i class="fa-regular fa-calendar mr-2"></i>${startDate} - ${endDate}`;
            }

            const searchTerm = document.getElementById('searchInput').value.toLowerCase();
            
            let itemsToRender = dayData.items;
            if (searchTerm) {
                itemsToRender = itemsToRender.filter(item => 
                    (item.title && item.title.toLowerCase().includes(searchTerm)) || 
                    (item.location && item.location.toLowerCase().includes(searchTerm)) ||
                    (item.note && item.note.toLowerCase().includes(searchTerm))
                );
            }

            if (itemsToRender.length === 0) {
                container.innerHTML = `
                    <div class="text-center py-20 text-gray-400">
                        <i class="fa-regular fa-face-sad-tear text-4xl mb-3"></i>
                        <p>找不到符合的行程</p>
                    </div>`;
                return;
            }

            const timelineDiv = document.createElement('div');
            timelineDiv.className = "space-y-4 pb-20";

            itemsToRender.forEach((item, index) => {
                const style = getTypeStyle(item.type);
                const mapLink = item.link || `https://www.google.com/maps/search/?api=1&query=${encodeURIComponent(item.location || item.title)}`;

                const html = `
                <div class="relative flex items-start group animate-fade-in" style="animation-delay: ${index * 50}ms">
                    <div class="absolute left-4 top-0 bottom-0 w-0.5 bg-gray-200 group-last:bottom-auto group-last:h-full z-0"></div>
                    
                    <div class="relative z-10 flex flex-col items-center mr-4 min-w-[3.5rem]">
                        <div class="h-8 w-8 rounded-full ${style.color} flex items-center justify-center border-2 border-white shadow-sm mb-1">
                            <i class="fa-solid ${style.icon} text-xs"></i>
                        </div>
                        <span class="text-xs font-bold text-gray-500 bg-white px-1 rounded">${item.time || '--:--'}</span>
                    </div>

                    <div class="flex-1 bg-white rounded-xl p-4 shadow-sm border border-gray-100 hover:shadow-md transition-shadow">
                        <div class="flex justify-between items-start mb-2">
                            <h3 class="font-bold text-gray-800 text-lg leading-tight">${item.title}</h3>
                            ${item.price ? `<span class="text-xs font-medium text-gray-500 bg-gray-100 px-2 py-1 rounded-full whitespace-nowrap ml-2">${item.price}</span>` : ''}
                        </div>
                        
                        ${item.location ? `
                        <div class="flex items-center text-sm text-gray-500 mb-3">
                            <i class="fa-solid fa-location-dot mr-1.5 text-travel-500"></i>
                            <span class="line-clamp-1">${item.location}</span>
                        </div>` : ''}

                        ${item.note ? `
                        <div class="bg-yellow-50 rounded-lg p-2.5 mb-3 text-sm text-gray-700 border-l-2 border-yellow-300">
                            <i class="fa-regular fa-lightbulb mr-1 text-yellow-500"></i> ${item.note}
                        </div>` : ''}

                        <a href="${mapLink}" target="_blank" class="inline-flex items-center justify-center w-full sm:w-auto px-4 py-2 bg-gray-50 hover:bg-travel-50 text-travel-600 text-sm font-medium rounded-lg border border-gray-200 transition-colors">
                            <i class="fa-solid fa-map-location-dot mr-2"></i> 導航
                        </a>
                    </div>
                </div>
                `;
                
                const wrapper = document.createElement('div');
                wrapper.innerHTML = html;
                timelineDiv.appendChild(wrapper.firstElementChild);
            });

            container.appendChild(timelineDiv);
        }

        // --- 5. CSV 處理 ---
        document.getElementById('csvInput').addEventListener('change', function(e) {
            const file = e.target.files[0];
            if (!file) return;

            const reader = new FileReader();
            reader.onload = function(e) {
                const text = e.target.result;
                parseCSV(text);
            };
            reader.readAsText(file);
        });

        function parseCSV(csvText) {
            const lines = csvText.split('\n');
            const newData = [];
            let currentDayObj = null;

            // 簡單的 CSV 解析 (處理逗號，移除引號)
            // 假設第一行是標題，從第二行開始
            for (let i = 1; i < lines.length; i++) {
                let line = lines[i].trim();
                if (!line) continue;

                // 處理引號內的逗號 (簡易版，不處理巢狀引號)
                // 將引號內的文字先隱藏逗號
                let inQuote = false;
                let tempLine = '';
                for(let char of line) {
                    if(char === '"') inQuote = !inQuote;
                    else if(char === ',' && inQuote) tempLine += '###COMMA###';
                    else tempLine += char;
                }
                
                const cols = tempLine.split(',').map(col => {
                    let c = col.trim().replace(/^"|"$/g, ''); // 移除前後引號
                    return c.replace(/###COMMA###/g, ','); // 還原逗號
                });

                // 欄位對應: 0:Day, 1:Date, 2:Time, 3:Title, 4:Type, 5:Location, 6:Price, 7:Note, 8:Link
                if (cols.length < 4) continue; // 至少要有標題

                const day = cols[0];
                const date = cols[1];

                if (!currentDayObj || currentDayObj.day != day) {
                    currentDayObj = {
                        day: day,
                        date: date,
                        items: []
                    };
                    newData.push(currentDayObj);
                }

                currentDayObj.items.push({
                    time: cols[2],
                    title: cols[3],
                    type: cols[4],
                    location: cols[5],
                    price: cols[6],
                    note: cols[7],
                    link: cols[8]
                });
            }

            if (newData.length > 0) {
                travelData = newData;
                localStorage.setItem('myTravelItinerary', JSON.stringify(travelData));
                currentDay = travelData[0].day;
                renderDayTabs();
                renderItinerary();
                toggleSettingsModal();
                alert('匯入成功！您的行程已更新。');
            } else {
                alert('匯入失敗，請檢查 CSV 格式。');
            }
        }

        // --- 6. UI 操作 ---
        function toggleSettingsModal() {
            const modal = document.getElementById('settingsModal');
            modal.classList.toggle('hidden');
        }

        function resetData() {
            if(confirm('確定要清除所有資料並恢復預設值嗎？')) {
                localStorage.removeItem('myTravelItinerary');
                location.reload();
            }
        }

        // 動畫樣式
        const styleSheet = document.createElement("style");
        styleSheet.innerText = `
            @keyframes fadeIn {
                from { opacity: 0; transform: translateY(10px); }
                to { opacity: 1; transform: translateY(0); }
            }
            .animate-fade-in {
                animation: fadeIn 0.5s ease-out forwards;
                opacity: 0;
            }
        `;
        document.head.appendChild(styleSheet);

        // 監聽
        document.getElementById('searchInput').addEventListener('input', renderItinerary);

        // 啟動
        init();

    </script>
</body>
</html>
