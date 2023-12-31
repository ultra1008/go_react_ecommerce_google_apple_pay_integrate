package auth

import (
	"database/sql"

	"github.com/bkielbasa/go-ecommerce/backend/auth/adapter"
	"github.com/bkielbasa/go-ecommerce/backend/auth/app"
	"github.com/bkielbasa/go-ecommerce/backend/auth/port"
	"github.com/bkielbasa/go-ecommerce/backend/internal/application"
	"github.com/bkielbasa/go-ecommerce/backend/internal/https"
	"github.com/gorilla/mux"
)

func New(db *sql.DB) application.BoundedContext {
	authStorage := adapter.NewPostgresAuthStorage(db)
	sessStorage := adapter.NewPostgresSessionStorage(db)
	appServ := app.NewAuth(authStorage, sessStorage)

	return &boundedContext{
		httpHandler: port.NewHTTP(appServ),
	}
}

type boundedContext struct {
	httpHandler port.HTTP
}

func (m boundedContext) MuxRegister(r *mux.Router) {
	r.HandleFunc("/api/v1/auth/login", m.httpHandler.Login).Methods("POST")
	r.HandleFunc("/api/v1/auth/login", https.EmptyHandler).Methods("OPTIONS")
	r.HandleFunc("/api/v1/auth/register", m.httpHandler.Register).Methods("POST")
	r.HandleFunc("/api/v1/auth/register", https.EmptyHandler).Methods("OPTIONS")
	r.HandleFunc("/api/v1/auth/me", m.httpHandler.Me).Methods("GET")
	r.HandleFunc("/api/v1/auth/me", https.EmptyHandler).Methods("OPTIONS")
	r.HandleFunc("/api/v1/auth/logout", m.httpHandler.Logout).Methods("DELETE")
	r.HandleFunc("/api/v1/auth/logout", https.EmptyHandler).Methods("OPTIONS")
}
