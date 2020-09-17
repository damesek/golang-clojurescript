package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func serve(ctx context.Context) (err error) {

	mux := http.NewServeMux()
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	mux.Handle("/", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			t, _ := template.ParseFiles("public/index.html")

			http.SetCookie(w,
				&http.Cookie{
					Domain:   "localhost",
					Expires:  time.Date(2020, 11, 23, 1, 5, 3, 0, time.UTC),
					HttpOnly: true,
					Name:     "main",
					Path:     "/",
					Value:    "this-is-my-cookie-value-not-hard-to-decrypt",
					Secure:   false,
				})

			t.Execute(w, nil)
		},
	))

	srv := &http.Server{
		Addr:    ":6969",
		Handler: mux,
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%+s\n", err)
		}
	}()

	log.Printf("server started")

	<-ctx.Done()

	log.Printf("server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err = srv.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shutdown Failed:%+s", err)
	}

	log.Printf("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}

	return
}

func main() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-c
		log.Printf("system call:%+v", oscall)
		cancel()
	}()

	if err := serve(ctx); err != nil {
		log.Printf("failed to serve:+%v\n", err)
	}
}
