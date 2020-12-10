package main

import (
	"net"
	"fmt"
	"strings"
)

// User user
type User struct {
	Username		string
	Chatwith		string
	Msg				string
	ServerMsg		string
}

var (
	userMap = make(map[string]net.Conn)
	user	= new(User)
)

func main() {

	addr, _ := net.ResolveTCPAddr("tcp4", "localhost:8899")
	lis, _ := net.ListenTCP("tcp4", addr)

	fmt.Println("server started at 8899!")

	for {
		conn, _ := lis.Accept()

		go func() {
			for {
				// read data to slice
				b := make([]byte, 1024)
				count, _ := conn.Read(b)
				fmt.Println("data received :", string(b[:count]))

				arrStr := strings.Split(string(b[:count]), "->")
				user.Username = arrStr[0]
				user.Chatwith = arrStr[1]
				user.Msg = arrStr[2]
				user.ServerMsg = arrStr[3]

				userMap[user.Username] = conn

				if revConn, ok := userMap[user.Chatwith]; ok && revConn != nil {
					
					n, err := revConn.Write([]byte(fmt.Sprintf("%s->%s->%s->%s", user.Username, user.Chatwith, user.Msg, user.ServerMsg)))
					if n == 0 || err != nil {
						delete(userMap, user.Chatwith)
						conn.Close()
						revConn.Close()
						break
					}
				} else {
					user.ServerMsg = user.Chatwith + " is not online!"
					conn.Write([]byte(fmt.Sprintf("%s->%s->%s->%s", user.Username, user.Chatwith, user.Msg, user.ServerMsg)))
				}
			}
		}()
	}
}