package main

import (
	"log"
	"golang.org/x/crypto/ssh"
	"net"
)


func main() {
	ip := "X.X.X.X"  //サーバのアドレス
	port := "22"  //ポート番号は文字列で指定
	user := "root"  //ユーザ名
	pass := "hoge"  //パスワード

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	client, err := ssh.Dial("tcp", ip+":"+port, config)
	//client, err := ssh.Dial("tcp", "X.X.X.X:P", config)
	if err != nil {
		log.Println(err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Println(err)
	}
	defer session.Close()

	//空のファイルを作成する
	session.Run("echo -n > empty.txt")
}
