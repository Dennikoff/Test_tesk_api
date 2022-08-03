package apiserver

import (
	"bytes"
	"encoding/json"
	"github.com/Dennikoff/UserTagApi/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	s := newServer(teststore.New())
	testcases := []struct {
		name       string
		payload    interface{}
		statusCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    "d.harke@yandex.ru",
				"password": "12345678",
				"nickname": "def1",
			},
			statusCode: http.StatusCreated,
		},
		{
			name: "invalidEmail",
			payload: map[string]string{
				"email":    "invalid@email",
				"password": "12345678",
				"nickname": "def1",
			},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name: "invalidPassword",
			payload: map[string]string{
				"email":    "valid@email.com",
				"password": "123",
				"nickname": "def1",
			},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name:       "invalidpayload",
			payload:    3,
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/signup", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.statusCode, rec.Code)
		})
	}
}
