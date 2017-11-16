package main

import "fmt"

func main() {
	r := initRouter()
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
	fmt.Printf("8080")
}
