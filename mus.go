package main

import (
	"github.com/breezechen/mus/app"
	"github.com/breezechen/mus/config"
)

func main() {
	app.Serve(config.REDIS_SERVER, config.REDIS_PASSWORD)
}
