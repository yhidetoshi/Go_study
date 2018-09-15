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
		handleConn(conn)
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
$ nc localhost 8000

- Listen関数は、あるネットワークポート、この場合はTCPポートのlocalhost:8000へ入ってくる
  接続を待つオブジェクトnet.Listenerを作成する
- リスナーのAcceptメソッドは、接続要求がくるまでまち、それから接続を表す, net.Connオブジェクトを返す
- net.Conn
*/