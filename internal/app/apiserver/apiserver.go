package apiserver

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/sessions"
	"gorest2/internal/app/store/sqlstore"
	"net/http"
)

func Start(config *Config) error {

	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	sessionsStore := sessions.NewCookieStore([]byte(config.SessionKey))
	srv := newServer(store, sessionsStore)

	fmt.Printf("Server start on port: %s", config.BinAddr)
	return http.ListenAndServe(config.BinAddr, srv)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}