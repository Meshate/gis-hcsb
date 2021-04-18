package main

import (
	gis_hcsb "gis-hcsb"
	"gis-hcsb/config"
	"gis-hcsb/router"
)

func main() {
	config.Init()
	gis_hcsb.Setup()
	app := router.App()
	_ = app.Listen(":11452")
}
