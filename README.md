# test-task-questions-go

Сервис для работы со сущностями "вопрос" и "ответ".

## Deployment

Команда для разворачивания проекта:

```bash
  docker-compose up -d
```
После разворачивания будут слушаться порты:

:8080 - сервис

:5432 - Postgres

## Примечания
При старте работы сервиса выполняется миграция по файлу из папки "migrations"

Сделан тест для http handler, который расположен в файле "internal/hadnlers/question_handler_test.go"

## API Reference

#### Получение вопроса

```http
  GET http://localhost:8080/questions/${id}
```

#### Получение всех вопросов

```http
  GET http://localhost:8080/questions/
```

#### Создание вопроса

```http
  POST http://localhost:8080/questions/
```

#### Удаление вопроса вместе с его ответами

```http
  DELETE http://localhost:8080/answers/${id}
```

#### Получение ответа

```http
  GET http://localhost:8080/answers/${id}
```

#### Создание ответа для вопроса

```http
  POST http://localhost:8080/questions/${id}/answers/
```

#### Удаление ответа

```http
  DELETE http://localhost:8080/answers/${id}
```






