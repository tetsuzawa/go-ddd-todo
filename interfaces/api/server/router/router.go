package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/rs/cors"

	"github.com/tetsuzawa/godddtodo/interfaces/api/server/handler"
	"github.com/tetsuzawa/godddtodo/interfaces/api/server/httputil"
	"github.com/tetsuzawa/godddtodo/interfaces/api/server/middleware"
)

type AppHandler struct {
	h func(http.ResponseWriter, *http.Request) (int, interface{}, error)
}

func (a AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, res, err := a.h(w, r)
	if err != nil {
		httputil.RespondErrorJson(w, status, err)
		return
	}
	httputil.RespondJSON(w, status, res)
	return
}

func NewRouter(h handler.AppHandler) *mux.Router {
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Authorization"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
	})

	commonChain := alice.New(
		middleware.RecoverMiddleware,
		corsMiddleware.Handler,
	)

	r := mux.NewRouter()
	r.Methods(http.MethodGet).Path("/ping").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	r.Methods(http.MethodGet).Path("/todo/{todo_id:[0-9]+}").Handler(commonChain.Then(AppHandler{h.ShowTodo}))
	r.Methods(http.MethodPost).Path("/todo").Handler(commonChain.Then(AppHandler{h.CreateTodo}))

	return r
}
