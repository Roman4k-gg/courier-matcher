Courier Matcher
Поиск ближайших курьеров по координатам — пет-проект на Go + PostgreSQL.

🚀 Запуск
bash
git clone https://github.com/Roman4k-gg/courier-matcher
cd courier-matcher

# 1. Создай БД courier_matcher и таблицу couriers (см. README полную)
# 2. Заполни тестовыми данными (100 курьеров Томска)
go mod tidy
go run cmd/matcher/main.go
📋 API
# Работает?
http://localhost:8080/ping

# Поиск курьеров
curl "http://localhost:8080/find-couriers?lat=56.48&lon=84.95&radius=3000&vehicle=bike"
Ответ:

json
{"couriers": [{"id": "...", "distance": 1234, "rating": 4.8, "vehicle": "bike"}]}
🏗️ Стек
Go 1.22

PostgreSQL

net/http

Сделано за 3 дня