package main

import (
	"net/http"
	"github.com/kamikazestar/selfdevelopment-ps-go-gettingstarted/controllers"
)

func main() {
	controllers.registerControllers()
	http.ListenAndServe(":3000", nil)
}
