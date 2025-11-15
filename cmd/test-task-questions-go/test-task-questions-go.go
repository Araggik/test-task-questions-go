package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Araggik/test-task-questions-go/internal/handlers"
	"github.com/Araggik/test-task-questions-go/internal/repositories"
	"github.com/Araggik/test-task-questions-go/internal/services"
	"github.com/Araggik/test-task-questions-go/migrations"
	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//psql -U postgres -d postgres -c 'SELECT * FROM questions'

func receiveConnStr() string {
	user := os.Getenv("DB_USER")
	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, name)
}

func checkError(err error, message string) {
	if err != nil {
		log.Fatal(message, err)
		panic(err)
	}
}

func main() {
	connStr := receiveConnStr()

	db, err := gorm.Open(postgres.Open(connStr))
	checkError(err, "Не удалось подключиться к БД")

	sqlDB, err := db.DB()
	checkError(err, "Не удалось получить sqlDB")

	goose.SetBaseFS(migrations.EmbedMigrations)

	err = goose.SetDialect("postgres")
	checkError(err, "Не удалось сменить диалект у goose")

	err = goose.Up(sqlDB, ".")
	checkError(err, "Не удалось прмиенить миграции")

	questionRepository := repositories.NewQuestionRepository(db)
	answerRepository := repositories.NewAnswerRepository(db)

	questionService := services.NewQuestionService(questionRepository)
	answerService := services.NewAnswerService(answerRepository)

	questionHandler := handlers.NewQuestionHandler(questionService)
	answerHandler := handlers.NewAnswerHandler(answerService)

	http.HandleFunc("GET /questions/{id}", questionHandler.GetQuestion)
	http.HandleFunc("GET /questions/", questionHandler.GetAllQuestion)
	http.HandleFunc("POST /questions/", questionHandler.CreateQuestion)
	http.HandleFunc("DELETE /questions/{id}", questionHandler.DeleteQuestion)

	http.HandleFunc("GET /answers/{id}", answerHandler.GetAnswer)
	http.HandleFunc("POST /questions/{id}/answers/", answerHandler.CreateAnswer)
	http.HandleFunc("DELETE /answers/{id}", answerHandler.DeleteAnswer)

	err = http.ListenAndServe(":8080", nil)
	checkError(err, "Не удалось развернуть сервер на порту 8080")
}
