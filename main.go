package main

import (
	"github.com/yawn-yawn-yawn/authJWT-go/handler"
	"log"
	"net/http"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/signup", handler.Signup(db))
	http.HandleFunc("/login", handler.Login(db))
	http.HandleFunc("/logout", handler.Logout(db))
	http.HandleFunc("/whoami", handler.WhoAmI(db))

	log.Println("server running...")
	http.ListenAndServe(":8080", nil)
}

func initDB() (*sql.DB, error) {
	return sql.Open("mysql", "root:password@tcp(localhost:3307)/auth_sample")
}
