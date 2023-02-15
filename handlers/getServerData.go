package handlers

import (
	"encoding/json"
	"net/http"
	"tripoley-server/models"
)

func GetServerData(game *models.ServerData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		G, err := json.Marshal(game)
		if err != nil {
			w.Write([]byte("Failed to Get Game Data"))
			w.WriteHeader(500)
			return
		}
		w.WriteHeader(200)
		w.Write(G)
	}
}
