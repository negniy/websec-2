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
    // Формируем URL для запроса к бэкенду
    const url = `http://localhost:8080/api/trains/through/?station=${station}&date=${date}`;

    // Отправляем GET-запрос
    fetch(url, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
        },
    })
        .then(response => {
            if (!response.ok) {
                throw new Error('Ошибка при получении данных');
            }
            return response.json();
        })
        .then(data => {
            console.log(data);
            displayTrains(data);
        })
        .catch(error => {
            console.error('Ошибка:', error);
            alert('Не удалось загрузить данные');
        });
}


function fetchTrainRoute(fromStation, toStation, date) {
    const url = `http://localhost:8080/api/trains/route/?from=${fromStation}&to=${toStation}&date=${date}`;
    
    fetch(url, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
        },
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Ошибка запроса: ' + response.status);
        }
        return response.json();
    })
    .then(data => displayTrains(data))
    .catch(error => console.error('Ошибка при выполнении запроса:', error));
}

function formatTime(dateString) {
    const date = new Date(dateString);
    const hours = date.getHours().toString().padStart(2, '0');
    const minutes = date.getMinutes().toString().padStart(2, '0');
    return `${hours}:${minutes}`;
}

function displayTrains(trains) {
    const scheduleDiv = document.getElementById('schedule');
    scheduleDiv.innerHTML = '<h2>Поезда:</h2>';
    
    if (trains.length > 0) {
        let scheduleHTML = '<ul>';
        trains.forEach(train => {
            const stationTo = train.station_to ? train.station_to : train.station_from;
            const arrival = train.arrival_time ? `Прибытие: ${stationTo} - ${formatTime(train.arrival_time)}` : '';
            const departure = train.departure_time ? `Отправление: ${train.station_from} - ${formatTime(train.departure_time)}` : '';
        
            scheduleHTML += `
                <li>
                    <p>Поезд №${train.number} (${train.title})</p>
                    ${arrival ? `<p>${arrival}</p>` : ''}
                    ${departure ? `<p>${departure}</p>` : ''}
                </li>
            `;
        });
        scheduleHTML += '</ul>';
        scheduleDiv.innerHTML += scheduleHTML;
    } else {
        scheduleDiv.innerHTML += '<p>Поезда не найдены для данного маршрута.</p>';
    }
}

document.getElementById('swapStationsBtn').addEventListener('click', function() {
    const fromStation = document.getElementById('fromStation');
    const toStation = document.getElementById('toStation');
    
    if (fromStation.value && toStation.value) {
        const temp = fromStation.value;
        fromStation.value = toStation.value;
        toStation.value = temp;
    } else {
        alert("Пожалуйста, введите обе станции перед тем, как менять их местами.");
    }
});