// @title BackendTestGolang API
// @version 1.0
// @description This is a sample cart/order API.
// @host localhost:8082
// @BasePath /
// @schemes http
package main

import (
	_ "backnedTestGolang/docs"
	"backnedTestGolang/internal/app"
)

func main() {

	app.Run()
}
