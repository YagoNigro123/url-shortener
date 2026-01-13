package main 

import (
	"fmt"
	"log"

	"github.com/YagoNigro123/url-shortener/internal/core"
	"github.com/YagoNigro123/url-shortener/internal/store"
)

func main(){
	connStr := "postgres://user:password@localhost:5432/shortener_db?sslmode=disable"

	postgressStore, err := store.NewPostgresStore(connStr)
	if err != nil {
		log.Fatalf("error with conect to db: %v", err)
	}

	//defer postgresStore().Close()

	service := core.NewService(postgressStore)

	//test 
	linkTest := "https://google.com"
	fmt.Println("Shortener "+ linkTest +"...")

	link, err := service.Shorten(linkTest)

	if err != nil{
		log.Fatalf("Fail in shortener: %v", err)
	}

	fmt.Printf("Sucessfuly! \n")
	fmt.Printf("Original URL: %s\n", link.Original)
	fmt.Printf("Generated ID:  %s\n", link.ID)
	fmt.Println("Check the database, it shuld be here")
}