package main

import (
	"net"
	"fmt"
	"os"
	"strings"
	"sync"
	"bufio"
)

// User user
type User struct {
	Username		string
	Chatwith		string
	Msg				string
	ServerMsg		string
}

var (
	user = new(User)
	wg		sync.WaitGroup
)

func main() {

	wg.Add(1)

	fmt.Println("input your name:")
	fmt.Scanln(&user.Username)
	fmt.Println("please input who the msg you want to deliver:")
	fmt.Scanln(&user.Chatwith)

	addr, _ := net.ResolveTCPAddr("tcp4", "localhost:8899")
	conn, _ := net.DialTCP("tcp4", nil, addr)

	go func() {
		// send
		fmt.Println(user.Chatwith, ":")
		for {
			// fmt.Scanln(&user.Msg) // 不接受空格
			inputReader := bufio.NewReader(os.Stdin)
			input, err := inputReader.ReadString('\n')
			if err == nil {
				user.Msg = input
			} else {
				fmt.Println("error input")
			}
			
			if user.Msg == "byebye" {
				conn.Close()
				wg.Done()
				os.Exit(0)
			}
			conn.Write([]byte(fmt.Sprintf("%s->%s->%s->%s", user.Username, user.Chatwith, user.Msg, user.ServerMsg)))
		}
	}()

	go func() {
		// receive
		for {
			revb := make([]byte, 1024)
			c, _ := conn.Read(revb)
			fromUser := new(User)
			arrStr := strings.Split(string(revb[:c]), "->")
			fromUser.Username = arrStr[0]
			fromUser.Chatwith = arrStr[1]
			fromUser.Msg = arrStr[2]
			fromUser.ServerMsg = arrStr[3]
			if fromUser.ServerMsg != "" {
				fmt.Println("\t\t\t [system]:", fromUser.ServerMsg)
			} else {
				fmt.Println("\t\t\t", fromUser.Username, ":", fromUser.Msg)
			}
		}
	}()
	wg.Wait()
}