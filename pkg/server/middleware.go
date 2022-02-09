package server

import (
	"context"
	"log"
	"net/http"
)

type Context struct {
	ID  string
	Ctx context.Context
	Err *Error
	Srv *Server
}

type Handler struct {
	f   func(c *Context, w http.ResponseWriter, req *http.Request)
	srv *Server
}

func ApiHandler(srv *Server, f func(c *Context, w http.ResponseWriter, req *http.Request)) http.Handler {
	return &Handler{
		f:   f,
		srv: srv,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := &Context{
		ID:  NewID(),
		Ctx: context.Background(),
		Srv: h.srv,
	}

	log.Printf("starting http request request_id=%s\n", c.ID)
	defer log.Printf("ending http request request_id=%s\n", c.ID)

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("X-request-id", c.ID)

	h.f(c, w, req)

	if c.Err != nil {
		c.Err.RequestID = c.ID
		log.Println(c.Err.ToJSON())
		if c.Err.StatusCode == http.StatusInternalServerError {
			c.Err.DetailedError = "Internal Server Error"
			c.Err.Params = nil
			c.Err.ID = ""
		}
		w.WriteHeader(c.Err.StatusCode)
		w.Write([]byte(c.Err.ToJSON()))
	}
}
