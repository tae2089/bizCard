package main

import (
	"bizCard/router"
	biztrace "bizCard/trace"
	"context"
	"go.opentelemetry.io/otel"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var APP_NAME = "bizCard"

func main() {
	f, err := os.Create("traces.txt")
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	tp := biztrace.Ttt(f)
	otel.SetTracerProvider(tp)
	router.SetupService()
	engine := router.SetupRouter()

	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			panic(err)
		}
	}()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server started listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown: %v", err)
	}

	select {
	case <-ctx.Done():
		log.Println("Timeout of 5 seconds.")
	}

	log.Println("Server exiting..")

}
