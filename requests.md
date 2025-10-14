curl -i http://localhost:8080/health

curl -i -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Купить молоко"}'

curl -i -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Kupit hleb"}'

  curl -i http://localhost:8080/tasks

  curl -i http://localhost:8080/tasks/2
  
  curl -i -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"ok"}'