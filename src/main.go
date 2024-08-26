package main

import(
	"fmt"
	"os"
	"webserver/src/log"
)


// type WebServer struct {
// 	Port		int
// 	Addr	 	utils.IPAddr
// }


func main() {
	logger := log.New(1, "MyLogger")
	ret, err := logger.Err("Hi!")

	if err != nil {
		fmt.Print("ERROR! ")
	}

	fmt.Fprintf(os.Stderr, "Log Printed: %d\n", ret)
}