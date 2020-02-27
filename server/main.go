package main

import (
	"fmt"
	"github.com/Planck1858/todo_list_cli/server/config"
	"github.com/Planck1858/todo_list_cli/server/router"
	"log"
	"net/http"
)

func main() {
	log.Println("Launching server ...")
	conf := config.GetConfig("./config/config.json")
	fmt.Println(conf.Port)
	r := router.NewRouter()

	err := http.ListenAndServe(conf.Port, r)
	if err != nil {
		panic(err)
	}
}
