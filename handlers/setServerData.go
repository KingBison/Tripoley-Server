package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"tripoley-server/models"
)

func SetServerData(game *models.ServerData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Error Reading Body"))
			return
		}
		var req models.ServerData
		err = json.Unmarshal(reqData, &req)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Error Unmarshaling Body"))
			return
		}

		*game = req
		w.WriteHeader(201)
	}
}
