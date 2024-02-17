package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/gorilla/mux"
	"github.com/pulsone21/PlaybookViewer/internal/dataloader"
	"github.com/robfig/cron"
)

type Server struct {
	*http.Server
	Port string
	Repo *git.Repository
	Cron *cron.Cron
	Data dataloader.DataHolder
}

func CreateServer() *Server {
	r := mux.NewRouter()

	port := os.Getenv("PORT")
	fmt.Println("Port: ", port)

	srv := Server{
		Server: &http.Server{
			Addr:    port,
			Handler: r,
		},
	}
	fmt.Println(srv)

	r.HandleFunc("/api/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("ok")
	})

	r.HandleFunc("/api/playbooks", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(srv.Data.GetPlaybooks())
	})

	r.HandleFunc("/api/playbooks/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		fmt.Println("id: ", id)
		pb := srv.Data.GetPlaybook(id)
		if pb == nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(fmt.Sprintf("Playbook with %s not found.", id))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(pb)
	})

	r.HandleFunc("/api/playbooks/view/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		gC := srv.Data.GetPlaybookView(id)

		if gC == nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(fmt.Sprintf("Playbook with %s not found.", id))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(gC)
	})
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))

	url := os.Getenv("GIT_REPO_URL")
	username := os.Getenv("GIT_REPO_USERNAME")
	token := os.Getenv("GIT_REPO_TOKEN")
	repo, err := dataloader.CloneRepo(url, username, token)
	if err != nil {
		log.Fatal(err)
	}
	srv.Repo = repo
	srv.ReprocessData()
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
	// reprocess all the playbooks and create the files which are used for dispaly
	dH, err := dataloader.ProcessData()
	if err != nil {
		log.Fatal(err)
	}
	s.Data = *dH
}

func (s *Server) ListenAndServe() error {
	return s.Server.ListenAndServe()
}
