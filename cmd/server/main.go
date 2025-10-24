package main

import (
	"github.com/namdang-fdp/go-ecommerce/internal/routers"
)

func main() {
	r := routers.NewRouter()
	r.Run()
}
