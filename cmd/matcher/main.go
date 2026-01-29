package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"courier-matcher/internal/handler"
	"courier-matcher/internal/repository"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
    	dsn = "postgres://roma:111111@localhost:5432/courier_matcher?sslmode=disable"
	}
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("failed to open db:", err)
	}
	defer db.Close()
	
	if err := db.Ping(); err != nil {
		log.Fatal("failed to ping db:", err)
	}
	
	repo := repository.NewPostgresRepo(db)

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status": "ok"}`))
	})
	
	http.HandleFunc("/find-couriers", handler.FindCouriers(repo))
	
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
