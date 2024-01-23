package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	dbs "short_url/pkg/databases"
	auth "short_url/pkg/features/auth/usecases"
	routes "short_url/pkg/routes"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Databases struct {
	SqlDB *sql.DB
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	dbUrl := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Can not connect to database", err)
	}

	auth.NewAuthProvider()

	dbs := dbs.ProvideDatabase(db)
	router := routes.SetupRoutes(dbs)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Host starting on: %v", port)
	srv.ListenAndServe()
}
