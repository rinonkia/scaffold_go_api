package main

import (
	"database/sql"
	"fmt"
	"github.com/rinonkia/go_api_tutorial/controllers"
	"github.com/rinonkia/go_api_tutorial/routers"
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

	// service層生成
	s := services.NewMyAppService(db)

	// controller層生成
	con := controllers.NewMyAppController(s)

	// router層生成
	r := routers.NewRouter(con)

	log.Println("server start at port: 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
