
Простой REST API для управления цитатами: добавление, получение всех, случайной и по автору, а также удаление по ID.

## 📦 Структура проекта

- `main.go` — запускает HTTP-сервер и настраивает маршруты.
- `handler.go` — содержит логику обработки запросов и хранилище данных (в памяти).
  
## 🚀 Запуск проекта

### Требования

- Go 1.18 или новее
- Установленный [Go Modules](https://go.dev/doc/modules/setup)

### Установка зависимостей

В корне проекта выполните:

```
go mod tidy
```

```Запуск
go run main.go
```
Сервер стартует по адресу: http://localhost:8080

🔗 Эндпоинты
Добавить цитату
POST /quotes
Тело запроса (JSON):

{
  "author": "Albert Einstein",
  "quote": "Life is like riding a bicycle. To keep your balance, you must keep moving."
}
Получить все цитаты
GET /quotes

Получить случайную цитату
GET /quotes/random

Получить цитаты по автору
GET /quotes?author=имя

Пример: /quotes?author=Albert Einstein

Удалить цитату по ID
DELETE /quotes/{id}
Пример: /quotes/1

🧪 Пример использования через curl
```Добавить цитату:

curl -X POST http://localhost:8080/quotes \
  -H "Content-Type: application/json" \
  -d '{"author": "Albert Einstein", "quote": "Imagination is more important than knowledge."}'
```
```Получить случайную:
curl http://localhost:8080/quotes/random
```
```Удалить цитату:
curl -X DELETE http://localhost:8080/quotes/1
```
📎 Зависимости
gorilla/mux — маршрутизация HTTP-запросов

```Установить:
go get github.com/gorilla/mux
```
🧼 Очистка
Все данные хранятся в оперативной памяти и теряются при перезапуске сервера.
