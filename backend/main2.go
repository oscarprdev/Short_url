// main.go
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"short_url/api"
	auth "short_url/pkg/features/auth/usecases"
	"short_url/storage"
)

func main() {
	port := os.Getenv("PORT")
	dbUrl := os.Getenv("DB_URL")

	listenAddr := flag.String("listenadrr", port, "the server address")
	store := storage.NewPostgresStorage(dbUrl)
	server := api.NewServer(store)

	auth.NewAuthProvider()

	fmt.Println("Server running on port:", *listenAddr)
	log.Fatal(server.Start(*listenAddr))
}
