package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rinonkia/go_api_tutorial/controllers"
	"github.com/rinonkia/go_api_tutorial/services"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("connect to DB")

	articleID := 1
	const sqlStr = `
		select *
		from articles
		where article_id = ?;
	`

	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	s := services.NewMyAppService(db)
	con := controllers.NewMyAppController(s)

	r := mux.NewRouter()
	r.HandleFunc("/hello", con.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", con.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", con.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", con.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", con.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", con.PostCommentHandler).Methods(http.MethodPost)

	log.Println("server start at port: 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
