package server

import (
	"RECStatGo/pkg/database"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	connector *database.Connector
}

func (s *Server) Serve(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	metaData := database.MetaData{
		Timestamp:      time.Now().Unix(),
		ProgramName:    query.Get("program"),
		Version:        query.Get("version"),
		Platform:       query.Get("platform"),
		InstallationId: query.Get("uuid"),
		Mac:            query.Get("mac"),
		UserName:       query.Get("uname"),
	}

	err := metaData.IsValid()
	if err != nil {
		handleError(&w, err)
		return
	}

	err = (*s.connector).AddEntry(metaData)
	if err != nil {
		handleError(&w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200 - Success"))
}

func (s *Server) Close() {
	(*s.connector).Close()
}

func handleError(w *http.ResponseWriter, err error) {
	fmt.Println(err)
	(*w).WriteHeader(http.StatusBadRequest)
	message := fmt.Sprintf("400 - bad request; %s", err)
	(*w).Write([]byte(message))
}

func NewServer(connectorName, connectorUri string) (Server, error) {
	connector, err := database.NewConnector(connectorName, connectorUri)
	if err != nil {
		return Server{}, err
	}
	return Server{connector: &connector}, nil
}
