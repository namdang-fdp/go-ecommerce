package main

import (
	"github.com/namdang-fdp/go-ecommerce/internal/routes"
)

func main() {
	r := routers.NewRouter()
	r.Run()
}
