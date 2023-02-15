package serverutils

import (
	"net/http"
	"time"
	"tripoley-server/models"

	"github.com/gorilla/mux"
)

func GetServerConfig(router *mux.Router) (*http.Server, *models.ServerData) {

	server := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	serverData := models.ServerData{
		Config: models.ServerConfig{
			Addr:      server.Addr,
			StartTime: time.Now().String(),
		},
	}

	return server, &serverData

}
