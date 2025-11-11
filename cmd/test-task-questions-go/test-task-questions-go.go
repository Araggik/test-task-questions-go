package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

//psql -U postgres -d postgres -c 'SELECT * FROM questions'

type question struct {
	id        int
	text      string
	createdAt string
}

func receiveConnStr() string {
	user := os.Getenv("DB_USER")
	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, name)
}

func main() {
	connStr := receiveConnStr()

	log.Println(connStr)

	db, err := sql.Open("postgres", connStr)

	if err == nil {
		defer db.Close()

		rows, err := db.Query("select * from questions")

		if err == nil {
			defer rows.Close()

			questions := []question{}

			for rows.Next() {
				q := question{}
				err := rows.Scan(&q.id, &q.text, &q.createdAt)
				if err == nil {
					questions = append(questions, q)

				} else {
					log.Fatal(err)
				}
			}

			for _, v := range questions {
				log.Println(v)
			}

		} else {
			log.Fatal(err)
		}

	} else {
		log.Fatal(err)
	}
}
