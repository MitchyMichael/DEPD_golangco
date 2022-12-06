package main

import (
	"DEPD_golangco/db"
	"DEPD_golangco/routes"
)

func main() {
	db.Init()

	e := routes.Init()
	e.Logger.Fatal(e.Start(":9999"))
}
