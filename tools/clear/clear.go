package main

import "os"

func main() {
	if err := os.Remove("gintest.db"); err != nil {
		panic(err)
	}
	println("sqlite database is deleted")
}
