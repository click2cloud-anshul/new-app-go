package main

import (
	"fmt"
	"github.com/new-app-go/router"
	"log"
	"net/http"
)

func main()  {

	r := router.Router()
	fmt.Println("connecting to port 9090...")
	//for _,env2 := range os.Environ(){
	//	fmt.Println(env2)
	//}
	log.Fatal(http.ListenAndServe(":9090", r))
}