package main

import "github.com/ismdeep/swag-crash-demo-001/rest"

func main() {
	if err := rest.Eng.Run("0.0.0.0:8000"); err != nil {
		panic(err)
	}
}
