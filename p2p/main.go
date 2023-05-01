package p2p

import (
	"bufio"
	"log"
	"net"
	"sync"
)

type Peer struct {
	Id   string
	Ip   string
	Port int32
}

func Client(address string, wg *sync.WaitGroup) {
	defer wg.Done()
	conn, err := net.Dial("tcp", address)

	log.Println(address)

	if err != nil {
		log.Fatalln(err)
	}

	data, _ := bufio.NewReader(conn).ReadString(';')

	log.Println(data)
}

func Server() {

}

func main() {
	ln, err := net.Listen("tcp", ":8000")

	var wg sync.WaitGroup
	wg.Add(3)
	defer wg.Wait()
	if err != nil {
		log.Fatalln(err)
	}
	go Client("127.0.0.1:8000", &wg)

	for {
		conn, err := ln.Accept()

		if err != nil {
			log.Fatalln(err)
		}

		go func(conn net.Conn) {
			conn.Write([]byte("Hello from Go;"))
		}(conn)
	}

}
