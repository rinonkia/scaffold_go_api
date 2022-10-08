package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/rinonkia/scaffold_go_api/api"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_USER_PASSWORD")
	dbDatabase := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	log.Println(dbConn)
	if err != nil {
		log.Print(err)
		return
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Print(err)
		return
	}
	log.Println("connect to DB")

	// router生成
	r := api.NewRouter(db)

	log.Println("server start at port: 8080")
	log.Println(http.ListenAndServe(":8080", r))
}
