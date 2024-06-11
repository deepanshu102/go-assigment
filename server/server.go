package server

import (
	"github.com/deepanshu102/go-assigment/repo"
	"github.com/deepanshu102/go-assigment/routers"
)

func StartServer(repo *repo.EmployeeRepo) {
	r := routers.SetupRouter(repo)
	r.Run(":8080") // Start the server on port 8080
}
