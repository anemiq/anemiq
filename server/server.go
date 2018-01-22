package server

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/anemiq/config"
	"github.com/graphql-go/graphql"
)

type Server struct {
	Config    *config.Config
	GQLSchema graphql.Schema
}

func New(config *config.Config, schema graphql.Schema) *Server {
	return &Server{
		Config:    config,
		GQLSchema: schema,
	}
}

func (s *Server) Start() {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}

		query, err := readReqQueryString(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		params := graphql.Params{
			Schema:        s.GQLSchema,
			RequestString: query.Query,
		}

		result := graphql.Do(params)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})

	http.ListenAndServe(":"+s.Config.Server.Port, nil)
}

func readReqQueryString(r *http.Request) (*queryString, error) {
	queryBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	if len(queryBody) == 0 {
		return nil, errors.New("missing body")
	}
	query := queryString{}
	json.Unmarshal([]byte(queryBody), &query)
	defer r.Body.Close()
	if err != nil {
		return nil, err
	}
	return &query, nil
}

type queryString struct {
	Query string `json:"query"`
}
