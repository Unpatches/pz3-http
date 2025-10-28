## Имя: Дорджиев Виктор
## Группа: ЭФМО-02-25
# Проект pz3-http

Задачи проекта
1. Освоить базовую работу со стандартной библиотекой net/http без сторонних фреймворков
2. Научиться поднимать HTTP-сервер, настраивать маршрутизацию через http.ServeMux.
3. Научиться обрабатывать параметры запроса (query, path), тело запроса (JSON/form-data) и формировать корректные ответы (код статуса, заголовки, JSON).
4. Научиться базовому логированию запросов и обработке ошибок.


---

## Установка и запуск

(Необходимы предустановленные Go и Git)

Клонировать репозиторий:

```
git clone https://github.com/Unpatches/pz3-http
cd pz3-http
```

Команда запуска сервера:

```
go run ./cmd/server
```


------

## Структура проекта

```plaintext
pz3-http/                     
├── cmd/                  
│   └── server/             
│       └── main.go       
├── internal/              
│   └── api/               
│   │   └── handlers.go     
│   │   └── handlers_test.go  
│   │   └── middleware.go
│   │   └── responses.go
│   └── storage/
│       └── memory.go       
├── go.mod                
└── makefile
└── requests.md        
```

## Отчёт о проделанной работе

```
curl -i http://localhost:8080/health
```
![photo_5436264944522884522_x](https://github.com/user-attachments/assets/dff2d903-baad-48b6-98b3-1e7ebde794cc)

```
curl -i http://localhost:8080/tasks
```
![photo_5436264944522884524_x](https://github.com/user-attachments/assets/aa3df659-3a3a-4e1b-89d5-2757fd2e8ae6)

```
curl -i -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Kupit hleb"}'
```
![photo_5436264944522884526_x](https://github.com/user-attachments/assets/95ea89d4-58ee-4575-a7c6-72bf166acff0)

```
curl -i http://localhost:8080/tasks/2
```
![photo_5436264944522884531_x](https://github.com/user-attachments/assets/04e81e73-7ba5-4bb0-b126-fa93bd8cb00f)

------

