package server

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log"
	"net/http"
	database "nproxy/internal/database/generated"
	"nproxy/internal/database/migrator"
	"os"
	"strconv"
	"time"

	_ "modernc.org/sqlite"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port    int
	queries *database.Queries
	ctx     context.Context
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	ctx := context.Background()

	db, err := sql.Open("sqlite", "sqlite3.db")
	if err != nil {
		log.Fatal(err)
	}

	migrator.Migrate(db)

	NewServer := &Server{
		port:    port,
		queries: database.New(db),
		ctx:     ctx,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
