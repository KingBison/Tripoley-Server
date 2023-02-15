package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"tripoley-server/helpers"
	"tripoley-server/models"
)

func AddPlayer(game *models.ServerData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Error Reading Body"))
			return
		}
		var req models.Player
		err = json.Unmarshal(reqData, &req)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Error Unmarshaling Body"))
			return
		}

		helpers.ProcessAddPlayer(game, req)

	}
}
