document.getElementById('searchBtn').addEventListener('click', function() {
    const fromStation = document.getElementById('fromStation').value.trim();
    const toStation = document.getElementById('toStation').value.trim();
    const date = document.getElementById('date').value;

    if (fromStation || toStation) {
        if (fromStation && toStation) {
            // Если обе станции введены, ищем расписание между ними
            fetchTrainRoute(fromStation, toStation, date);
        } else {
            // Если введена только одна станция, показываем все поезда через неё
            fetchTrainsThroughStation(fromStation || toStation, date);
        }
    } else {
        alert('Пожалуйста, введите хотя бы одну станцию.');
    }
});

function fetchTrainsThroughStation(station, date) {
    // Здесь будет код для запроса всех поездов через станцию через ваш бэкенд
    console.log(`Поезда через станцию: ${station}`);

    // Симуляция получения данных
    const trains = [
        { number: '1203', from: 'Самара', to: 'Москва', departureTime: '12:30', arrivalTime: '13:45', date: date },
        { number: '1305', from: 'Самара', to: 'Санкт-Петербург', departureTime: '14:45', arrivalTime: '16:00', date: date }
    ];

    displayTrains(trains);
}

function fetchTrainRoute(fromStation, toStation, date) {
    // Здесь будет код для запроса расписания между двумя станциями через ваш бэкенд
    console.log(`Поезда между станциями: ${fromStation} и ${toStation}`);

    // Симуляция получения данных для маршрута
    const trains = [
        { number: '1501', from: fromStation, to: toStation, departureTime: '10:00', arrivalTime: '12:00', date: date },
        { number: '1604', from: fromStation, to: toStation, departureTime: '13:15', arrivalTime: '15:15', date: date }
    ];

    displayTrains(trains);
}

function displayTrains(trains) {
    const scheduleDiv = document.getElementById('schedule');
    scheduleDiv.innerHTML = '<h2>Поезда:</h2>';
    
    if (trains.length > 0) {
        let scheduleHTML = '<ul>';
        trains.forEach(train => {
            scheduleHTML += `
                <li>
                    <p>Поезд №${train.number} (${train.from} - ${train.to})</p>
                    <p>Прибытие: ${train.from} : ${train.departureTime} ${train.date}</p>
                    <p>Отправление: ${train.to} : ${train.arrivalTime} ${train.date}</p>
                </li>
            `;
        });
        scheduleHTML += '</ul>';
        scheduleDiv.innerHTML += scheduleHTML;
    } else {
        scheduleDiv.innerHTML += '<p>Поезда не найдены для данного маршрута.</p>';
    }
}

// Логика для отображения карты
document.getElementById('toggleMapBtnFrom').addEventListener('click', function() {
    const mapDiv = document.getElementById('mapFrom');
    if (mapDiv.style.display === 'none') {
        mapDiv.style.display = 'block';
        this.innerText = 'Свернуть карту';
    } else {
        mapDiv.style.display = 'none';
        this.innerText = 'Развернуть карту';
    }
});

document.getElementById('toggleMapBtnTo').addEventListener('click', function() {
    const mapDiv = document.getElementById('mapTo');
    if (mapDiv.style.display === 'none') {
        mapDiv.style.display = 'block';
        this.innerText = 'Свернуть карту';
    } else {
        mapDiv.style.display = 'none';
        this.innerText = 'Развернуть карту';
    }
});
