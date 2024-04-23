package main

import (
	"fmt"

	"github.com/redianmf/transportation-ticketing-app/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDatabasePg()
}

func main() {
	fmt.Println("Hello")
}
