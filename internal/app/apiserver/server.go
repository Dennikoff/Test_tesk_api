package apiserver

import (
	"github.com/Dennikoff/UserTagApi/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	router mux.Router
	logger logrus.Logger
	store  store.Store
}

func newServer(store *store.Store) {

}
