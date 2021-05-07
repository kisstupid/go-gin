package server

import "github.com/kisstupid/go-gin/configs"

func Init() {
	config := configs.GetConfig()
	r := NewRouter()

	r.Run(config.GetString("server.port"))
}
