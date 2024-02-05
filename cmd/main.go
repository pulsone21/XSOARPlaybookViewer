package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"os/user"

	"github.com/go-git/go-git/v5"
	"github.com/gorilla/mux"
	"github.com/pulsone21/PlaybookViewer/internal/dataloader"
	"github.com/robfig/cron/v3"
)

type Server struct {
	*http.Server
	Port string
	Repo *git.Repository
	Cron *cron.Cron
}

func main() {
	fmt.Println("Hello World")
	srv := createServer()

	fmt.Println(user.Current())

	go func() {
		fmt.Printf("Service started on http://localhost:%s \n", srv.Server.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(context.TODO())
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	fmt.Println("\nshutting down")
}

func createServer() *Server {
	r := mux.NewRouter()

	r.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("ok")
	})

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))

	port := os.Getenv("PORT")
	fmt.Println("Port: ", port)

	srv := Server{
		Server: &http.Server{
			Addr:    port,
			Handler: r,
		},
	}
	fmt.Println(srv)
	url := os.Getenv("GIT_REPO_URL")
	username := os.Getenv("GIT_REPO_USERNAME")
	token := os.Getenv("GIT_REPO_TOKEN")
	repo, err := dataloader.CloneRepo(url, username, token)
	if err != nil {
		log.Fatal(err)
	}
	srv.Repo = repo

	srv.Cron = cron.New()
	srv.Cron.AddFunc("* 10 4 * * *", srv.ReprocessData)
	srv.Cron.Start()

	return &srv
}

func (s *Server) ReprocessData() {
	err := dataloader.PullRepo(s.Repo, os.Getenv("GIT_REPO_USERNAME"), os.Getenv("GIT_REPO_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
	//reprocess all the playbooks and create the files which are used for dispaly
	err = dataloader.ProcessData()
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Server) ListenAndServe() error {
	return s.Server.ListenAndServe()
}
