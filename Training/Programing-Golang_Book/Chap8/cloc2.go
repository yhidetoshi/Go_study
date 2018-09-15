package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main(){
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil{
		log.Fatal(err)
	}
	for{
		conn, err := listener.Accept()
		if err != nil{
			log.Print(err)
			continue
		}
		go handleConn(conn) // 接続を平行して処理する
	}
}

func handleConn(c net.Conn){
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
	}
	time.Sleep(1*time.Second)
}

/*
$ go build clock2
$ ./clock2 &

$ go build netcat1
$ ./netcat1

$ ./netcat1

→ go handleConn(conn) で、接続を平行して処理することができる
*/