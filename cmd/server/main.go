package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/amaan287/zorvyn-assignment/constants"
)

func main() {
	env, err := constants.GetEnv()
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("server is running on http://localhost%s", env.PORT)
	if err = http.ListenAndServe(env.PORT, nil); err != nil {
		log.Panic(err)
	}
}
