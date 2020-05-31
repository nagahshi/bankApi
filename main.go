package main

import (
	"github.com/nagahshi/bankApi/api"
	"github.com/nagahshi/bankApi/migrations"
)

func main() {
	migrations.Migrate()

	api.StartupApi()
}
