// @title           Content Creator Inshight
// @version         1.0
// @description     A Comprehensive Semantic Analysis Web Service.
// @BasePath  /

package main

import (
	"fmt"

	"github.com/SoyebSarkar/content-creator-insight/app"
	_ "github.com/SoyebSarkar/content-creator-insight/docs"
)

func main() {
	fmt.Println("Aplication Start")
	app.StartApplication()
}
