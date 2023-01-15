package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}
var api_url string = "http://192.168.144.210:80/api"
var storage02_mac string = "74:da:38:c2:3d:b6"
var storage02_ip string = "192.168.144.142"

func checkServerUp(loops int, ip string) int {

	for i := 1; i <= loops; i++ {
		// check it is awake
		cmd := exec.Command("/bin/sh", "-c", "ping -c 1 "+ip+" >/dev/null && echo 1")
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
			return 0
		}
		if string(out.String()[0]) == "1" {
			fmt.Println("Server Up.")
			return 1
		}
		time.Sleep(time.Duration(1000) * time.Millisecond)
	}
	return 0
}

func main() {
	// wake storage02
	body := []byte(`{"Name":"wol","Args":["wake","` + storage02_mac + `"]}`)
	r, err := http.NewRequest("POST", api_url, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	r.Header.Add("Content-Type", "application/json")
	resp, err := myClient.Do(r)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	t := checkServerUp(60, storage02_ip)

	// check it is awake
	if t == 1 {
		// mount storage02
		cmd := exec.Command("/bin/sh", "-c", "sudo mount -a")
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err = cmd.Run()
		if err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
			return
		}

		// run backup
		cmd1 := exec.Command("/bin/sh", "-c", "rsync -avzt /mnt/ns01/ /mnt/ns03/Backup/ns01/ ; rsync -avzt /mnt/ns02/ /mnt/ns03/Backup/ns02/")
		var out1 bytes.Buffer
		var stderr1 bytes.Buffer
		cmd.Stdout = &out1
		cmd.Stderr = &stderr1
		err = cmd1.Run()
		if err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
			return
		}

		// shutdown
		cmd2 := exec.Command("/bin/sh", "-c", "sudo shutdown now")
		var out2 bytes.Buffer
		var stderr2 bytes.Buffer
		cmd2.Stdout = &out2
		cmd2.Stderr = &stderr2
		err = cmd2.Run()
		if err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + stderr2.String())
			return
		}
	}
}

// env GOOS=linux GOARCH=amd64 go build -o /mnt/ns01/servers/factotum/01/api/apps/linux_amd64
