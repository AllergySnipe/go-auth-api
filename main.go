package main

import (
	"github.com/AllergySnipe/go-auth-api/initializer"
	"github.com/AllergySnipe/go-auth-api/routes"
)

func init() {
	initializer.EnvVarInit()
	initializer.ConnectToDb()
	initializer.MigrateDb()
}

func main() {
	r := routes.SetupRouter()
	r.Run(":8080")

}
