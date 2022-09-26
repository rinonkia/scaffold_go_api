package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rinonkia/go_api_tutorial/controllers"
	"github.com/rinonkia/go_api_tutorial/routers"
	"github.com/rinonkia/go_api_tutorial/services"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
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
	fmt.Println(dbConn)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("connect to DB")

	// service層生成
	s := services.NewMyAppService(db)

	// controller層生成
	ArticleController := controllers.NewArticleController(s)
	CommentController := controllers.NewCommentController(s)

	// router層生成
	r := routers.NewRouter(ArticleController, CommentController)

	log.Println("server start at port: 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
