package main

import (
	"fmt"
	"gin-example/config"
	"gin-example/router"

)

func init(){
	fmt.Println("main包 init")
}
func main() {
	var router = router.SetRouter();

	router.Run( fmt.Sprintf("%s:%d", config.Main.HttpHost,config.Main.HttpPort))
}


