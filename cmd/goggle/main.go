package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/abiriadev/goggle/pkg/goggle"
)

func pathArg() string {
	flag.Parse()
	f := flag.Arg(0)

	if f == "" {
		return "index.gob"
	}

	return f
}

type QueryRequest struct {
	Q string `json:"q"`
}

type QueryHandler struct {
	Goggle goggle.Goggle
}

func (qh *QueryHandler) Query(query string) (goggle.ResultSet, error) {
	return qh.Goggle.Query(query)
}

func (qh *QueryHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	slog.Info("Request", "method", r.Method, "path", r.URL)

	rw.Header().Set("Access-Control-Allow-Origin", "*")

	var req QueryRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		rw.WriteHeader(400)
		rw.Write([]byte(err.Error()))
		return
	}

	slog.Info("Query", "q", req.Q)

	rs, err := qh.Query(req.Q)
	if err != nil {
		rw.WriteHeader(400)
		rw.Write([]byte(err.Error()))
	}

	slog.Info("ResultSet", "items found", len(rs.Results))

	err = rs.MarshalJsonTo(rw)
	if err != nil {
		rw.WriteHeader(500)
		rw.Write([]byte(err.Error()))
	}
}

func main() {
	port := flag.Int("port", 6099, "port number to bind")
	f := pathArg()

	goggle, err := goggle.NewGoggle(goggle.Config{Limit: 10})
	if err != nil {
		panic(err)
	}

	err = goggle.Load(f)
	if err != nil {
		panic(err)
	}

	qh := QueryHandler{Goggle: goggle}

	http.Handle("/search", &qh)

	slog.Info("server is listening on", "port", strconv.Itoa(*port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
