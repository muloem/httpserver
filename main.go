package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	const wwwFolderPath = "./www"

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listening on", ln.Addr())

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		mParts := strings.Split(msg, " ")

		// check local www folder for file in path
		data, err := os.ReadFile(fmt.Sprintf("%s/%s", wwwFolderPath, mParts[1]))
		if err != nil {
			if os.IsNotExist(err) {
				resp := fmt.Sprintf("HTTP/1.1 404 Not Found\r\n\r\nRequested path: %s\r\n", mParts[1])
				conn.Write([]byte(resp))
				conn.Close()
			}
		}

		resp := fmt.Sprintf("HTTP/1.1 200 OK\r\n\r\n%s\r\n", string(data))
		conn.Write([]byte(resp))
		conn.Close()
	}
}
