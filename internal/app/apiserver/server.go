package apiserver

import (
	"encoding/json"
	"errors"
	"github.com/Dennikoff/UserTagApi/internal/app/model"
	"github.com/Dennikoff/UserTagApi/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

var (
	errUserNotFound = errors.New("email or password are not correct")
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}
	s.configureRouter()
	return s
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/signup", s.handleUserCreate()).Methods(http.MethodPost)
	s.router.HandleFunc("/login", s.handleUserLogin()).Methods(http.MethodPost)
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) handleUserLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &model.User{}
		if err := json.NewDecoder(r.Body).Decode(user); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		us, err := s.store.User().FindByEmail(user.Email)
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, errUserNotFound)
			return
		}
		if err := us.CheckPassword(user.Password); err != nil {
			s.error(w, r, http.StatusUnauthorized, errUserNotFound)
			return
		}
		s.createSession()

		s.response(w, r, http.StatusOK, us)
	}
}

func (s *server) handleUserCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &model.User{}
		if err := json.NewDecoder(r.Body).Decode(user); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		if err := s.store.User().Create(user); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.response(w, r, http.StatusCreated, user)
	}
}

func (s *server) createSession() {

}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.response(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) response(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		_ = json.NewEncoder(w).Encode(data)
	}
}
