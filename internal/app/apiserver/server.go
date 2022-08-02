package apiserver

import (
	"github.com/Dennikoff/UserTagApi/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type server struct {
	router mux.Router
	logger logrus.Logger
	store  store.Store
}

func newServer(store store.Store) *server {
	return &server{
		store: store,
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
