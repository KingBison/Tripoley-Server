package main

import (
	"fmt"
	"log"
	"tripoley-server/handlers"
	"tripoley-server/serverutils"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.Use(serverutils.Middleware)
	server, SERVER_DATA := serverutils.GetServerConfig(router)

	router.HandleFunc("/getServerData", handlers.GetServerData(SERVER_DATA)).Methods("GET")
	router.HandleFunc("/setServerData", handlers.SetServerData(SERVER_DATA)).Methods("POST")

	router.HandleFunc("/addPlayer", handlers.AddPlayer(SERVER_DATA)).Methods("POST")

	fmt.Println("tripoley-server is running on", SERVER_DATA.Config.Addr, "at", SERVER_DATA.Config.StartTime)
	log.Fatal(server.ListenAndServe())
}
