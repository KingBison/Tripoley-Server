package models

type ServerData struct {
	Players []Player
	Game    Game
	Config  ServerConfig
}

type ServerConfig struct {
	Addr      string
	StartTime string
}
