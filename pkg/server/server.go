package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/joshuaAllday/perkbox/pkg/placeholder"
)

var (
	filePath string = "./input/data.csv"
	apiURL   string = "https://jsonplaceholder.typicode.com/todos"
)

type Server struct {
	api        *Api
	rootRouter *mux.Router
	server     *http.Server
	service    *placeholder.Placeholder
	filePath   string
}

func NewServer() (*Server, error) {
	server := &Server{
		rootRouter: mux.NewRouter(),
	}
	server.api = InitApi(server)
	todos, err := placeholder.NewToDos(apiURL)
	if err != nil {
		return nil, err
	}

	server.filePath = filePath
	server.service = todos
	return server, nil
}

func (s *Server) Start() error {
	s.server = &http.Server{
		Handler: s.rootRouter,
		Addr:    ":8080",
	}

	fmt.Println("Starting server running on :8080")
	go s.server.ListenAndServe()
	return nil
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	fmt.Println("Shutting down server")
	return s.server.Shutdown(ctx)
}

func (s *Server) GetTodos(ctx context.Context, csv [][]string) ([]*placeholder.TODO, *Error) {
	todos, err := s.service.GetTodos(ctx, csv)
	if err != nil {
		return nil, NewError("server.GetTodos", "server.todos.get.error", err.Error(), nil, http.StatusInternalServerError)
	}
	return todos, nil
}
