package main

import (
	"fmt"
	"hospital-portal/internal/database"
	"hospital-portal/internal/server"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main(){
	fmt.Println(".")

	closeSignal := make(chan os.Signal, 1)
	signal.Notify(closeSignal, syscall.SIGINT, syscall.SIGTERM)

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	
	go server.StartServer(r)
	log.Println("server started.")

	err:=database.InitDB()
	if err != nil {
		log.Println(err)
		return
	}

	server.RegisterRoutes(r)
	<-closeSignal

	server.StopServer()
	database.CloseDB()
}