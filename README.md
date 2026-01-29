Courier Matcher

Это мой учебный проект, я сделал его за 2 вечера

В чем суть ? 

Этот проект иммитирует сервис по подбору курьеров для заказа

Пользователь отправляет запрос со своей долготой, шириной и опцианально вид транспорта курьера(foot/bike/car)

Пример: ```http://localhost:8080/find-couriers?lat=56.48&lon=84.95&radius=5000```

сервис возращает топ 5 курьеров по удаленности с нужным видом транспорта

```json
{

    "couriers": [
    
        {
        
            "id": "0513e77b-0984-4754-8486-f6bf2f5c1772",
            
            "lat": 56.47606754392623,
            
            "lon": 84.96297396192122,
            
            "distance": 908.8157623616361,
            
            "rating": 4.735034201479575,
            
            "vehicle": "bike"
            
        }
    ]    

}
```

Инструкция по запуску:
С Docker:
1. git clone https://github.com/Roman4k_gg/courier-matcher
2. cd courier-matcher
3. docker compose up --build

Бкз Docker:
1. Выполнить init.sql из репозитория
2. go mod tidy
3. go run cmd/matcher/main.go


Стек: Go + PostgreSQL + Docker


