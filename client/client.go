package main

import (
	"fmt"
	"log"
	"net"
	"time"
	"strconv"
	"github.com/gansidui/gotcp/examples/echo"
)

func main() {
	var connCount = 10

	for j := 0;j < connCount; j++{
		go tcpClient("uid_"+strconv.Itoa(j))
	}

	for{
		time.Sleep(time.Second)
	}
}

func tcpClient(uid string){
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8989")
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	echoProtocol := &echo.EchoProtocol{}
	for i := 0; i < 3; i++ {
		// write
		conn.Write(echo.NewEchoPacket([]byte("{\"uid\":\"" +uid + "\"}"), false).Serialize())

		// read
		p, err := echoProtocol.ReadPacket(conn)
		if err == nil {
			echoPacket := p.(*echo.EchoPacket)
			fmt.Printf("Server reply:[%v] [%v]\n", echoPacket.GetLength(), string(echoPacket.GetBody()))
		}

		time.Sleep(2 * time.Second)
	}

	conn.Close()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
