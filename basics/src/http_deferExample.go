package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.google.com/robots.txt") // Get returns a pointer to a Response struct
	if err != nil {
		log.Fatal(err) // Fatal is equivalent to Print() followed by a call to os.Exit(1)
	}
	// The defer statement defers the execution of a function until the surrounding function returns
	defer res.Body.Close()                  // Close closes the Response Body
	robots, err := ioutil.ReadAll(res.Body) // ReadAll reads from r until an error or EOF and returns the data it read
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
