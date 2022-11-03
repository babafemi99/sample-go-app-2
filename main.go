package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"task-1/internal/service"
	"time"
)

func main() {
	arthSrv := service.NewArthSrv()
	app1 := NewApp(arthSrv)

	router := chi.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)
	router.Get("/data", app1.GetData)
	router.Post("/data", app1.PostData)

	port := ":5101"
	srv := http.Server{
		Addr:        port,
		Handler:     handler,
		IdleTimeout: 120 * time.Second,
	}
	log.Printf("SERVER STARTING ON PORT %v \n\n", port)
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Printf("ERROR STARTING SERVER: %v", err)
			os.Exit(1)
		}
	}()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Printf("Closing now, We've gotten signal: %v", sig)

	ctx := context.TODO()
	srv.Shutdown(ctx)
}
