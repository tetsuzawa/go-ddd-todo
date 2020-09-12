package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"

	"github.com/tetsuzawa/godddtodo/interfaces/api/server/dbutil"
	"github.com/tetsuzawa/godddtodo/interfaces/api/server/router"
	"github.com/tetsuzawa/godddtodo/registry"
)

type Server struct {
	db     *sql.DB
	router *mux.Router
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init(datasource string) error {
	db, err := dbutil.Init(datasource)
	if err != nil {
		return err
	}
	s.db = db
	r := registry.NewRegistry(db)
	s.router = router.NewRouter(r.NewAppHandler())
	return nil
}

func (s *Server) Run(port int) {
	defer s.Close()
	log.Printf("Listening on port %d", port)
	err := http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		handlers.CombinedLoggingHandler(os.Stdout, s.router),
	)
	if err != nil {
		panic(err)
	}
}
func (s *Server) Close() {
	s.db.Close()
}
