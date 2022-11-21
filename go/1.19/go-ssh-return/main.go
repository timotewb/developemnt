package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

func main() {
	// ssh config
	hostKeyCallback, err := knownhosts.New("/home/timotewb/.ssh/known_hosts")
	if err != nil {
		log.Fatal(err)
	}
	config := &ssh.ClientConfig{
		User: "sysadmin",
		Auth: []ssh.AuthMethod{
			ssh.Password("rX#=TZ4V"),
		},
		HostKeyCallback: hostKeyCallback,
	}
	// connect ot ssh server
	conn, err := ssh.Dial("tcp", "192.168.144.211:22", config)
	if err != nil {
		log.Fatal(err)
	}

	session, err := conn.NewSession()
	if err != nil {
		log.Fatal(err)
	}

	var buff bytes.Buffer
	session.Stdout = &buff
	if err := session.Run("systemctl status api.service --lines 5"); err != nil {
		log.Fatal(err)
	}
	var s string = buff.String()
	s1 := strings.Split(s, "\n")
	for _, s2 := range s1 {
		if strings.Contains(s2, "└─") || strings.Contains(s2, "├─") {
			fmt.Println("yes - ", s2)
		} else {
			fmt.Println(s2)
		}
	}

	defer session.Close()
	defer conn.Close()
}
