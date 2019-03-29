package main

import (
	"github.com/micro-plat/hydra/hydra"
)

type assistantapi struct {
	*hydra.MicroApp
}

func main() {

	app := &assistantapi{
		hydra.NewApp(
			hydra.WithPlatName("gsms"),
			hydra.WithSystemName("merchant-api"),
			hydra.WithServerTypes("api"),
			hydra.WithDebug()),
	}

	app.init()
	app.install()
	app.handling()

	app.Start()
}
