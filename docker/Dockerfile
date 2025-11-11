FROM golang:1.24.4-alpine

WORKDIR /app

# Копируем исходный код
COPY . .

RUN go mod download

# Собираем приложение
RUN go build -o main ./cmd/test-task-questions-go

EXPOSE 8080

CMD ["./main"]