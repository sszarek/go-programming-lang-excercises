package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var port int
var timezone string
var location *time.Location

func init() {
	flag.IntVar(&port, "port", 8080, "Port on which program will listening for connections")
	flag.StringVar(&timezone, "tz", "Europe/Warsaw", "Sets the clocks timezone")
}

func main() {
	flag.Parse()
	loadLocatioon()

	addr := fmt.Sprintf("localhost:%d", port)
	listener, err := net.Listen("tcp", addr)

	fmt.Printf("Server listening on %s", addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConnection(conn)
	}
}

func loadLocatioon() {
	var err error
	location, err = time.LoadLocation(timezone)
	if err != nil {
		log.Fatal(err)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()

	for {
		_, err := io.WriteString(c, time.Now().In(location).Format("15:04:05\n"))
		if err != nil {
			return
		}

		time.Sleep(1 * time.Second)
	}
}
