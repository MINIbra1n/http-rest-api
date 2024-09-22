package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/MINIbra1n/http-rest-api/internal/app/store/sqlstore"
	"github.com/gorilla/sessions"
)

// Start
func Start(config *Config) error {
	db, err := newDB(config.DatabaseUrl)
	if err != nil {
		return err
	}
	defer db.Close()
	store := sqlstore.New(db)
	sessionsStore := sessions.NewCookieStore([]byte(config.SessionKey))
	srv := newServer(store, sessionsStore)
	return http.ListenAndServe(config.BildAddr, srv)
}

func newDB(databaseUrl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
