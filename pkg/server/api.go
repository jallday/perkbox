package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joshuaAllday/perkbox/pkg/utils"
)

type Routes struct {
	root *mux.Router
	todo *mux.Router
}

type Api struct {
	routes *Routes
	srv    *Server
}

func InitApi(server *Server) *Api {
	api := &Api{
		srv: server,
		routes: &Routes{
			root: server.rootRouter,
		},
	}

	api.routes.todo = api.routes.root.PathPrefix("/todo").Subrouter()

	api.initToDos()
	return api
}

func (api *Api) initToDos() {
	api.routes.todo.Handle("", ApiHandler(api.srv, listTodos)).Methods("GET")
}

func listTodos(c *Context, w http.ResponseWriter, req *http.Request) {
	csv, err := utils.ReadCsvFile(c.Srv.filePath)
	if err != nil {
		c.Err = NewError("api.listTodos", "app.file.read.error", err.Error(), map[string]interface{}{
			"path": c.Srv.filePath,
		}, http.StatusInternalServerError)
		return
	}

	todos, nErr := c.Srv.GetTodos(c.Ctx, csv)
	if nErr != nil {
		c.Err = nErr
		return
	}

	if err := json.NewEncoder(w).Encode(&todos); err != nil {
		c.Err = NewError("api.listTodos", "api.response.encoder", err.Error(), nil, http.StatusInternalServerError)
	}

}
