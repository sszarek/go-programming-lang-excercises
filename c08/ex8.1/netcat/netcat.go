package main

import (
	"strings"
	"io"
	"log"
	"net"
	"os"
	"fmt"
	"time"
)

func main() {
	clocks := getClocks()

	for _, addr := range clocks {
		go startListening(addr)
	}

	for {
		time.Sleep(time.Minute)
	}
}

func getClocks() map[string]string {
	result := make(map[string]string)
	for _, arg := range os.Args[1:] {
		split := strings.Split(arg, "=")
		if len(split) != 2 {
			continue
		}

		result[split[0]] = split[1]
	}

	return result
}

func startListening(addr string) {
	fmt.Println("aaddasd")
	fmt.Println(addr)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connected to %s", addr)
	defer conn.Close()

	mustCopy(os.Stdout, conn)
}

func mustCopy(target io.Writer, source io.Reader) {
	_, err := io.Copy(target, source)
	if err != nil {
		log.Fatal(err)
	}
}
