Courier Matcher

Это мой учебный проект, я сделал его за 2 вечера

В чем суть ? 

Этот проект иммитирует сервис по подбору курьеров для заказа
Пользователь отправляет запрос со своей долготой, шириной и опцианально вид транспорта курьера(foot/bike/car)

сервис возращает топ 5 курьеров по удаленности с нужным видом транспорта

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
