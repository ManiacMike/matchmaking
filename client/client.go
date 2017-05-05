package main

import (
	"fmt"
	"log"
	"net"
	"time"
	"strconv"
	"github.com/gansidui/gotcp/examples/echo"
	"encoding/json"
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

	startPacket := map[string]interface{}{
		"game":	"csgo",
		"mode": "competitive",
		"pid":       uid,
		"gid":        0,
		"maps":  []string{},
	}
	packetBody, err := json.Marshal(startPacket)
	checkError(err)
	// return string(body)

	// for i := 0; i < 3; i++ {
		// write
	conn.Write(echo.NewEchoPacket(packetBody, false).Serialize())

	// read
	echoProtocol := &echo.EchoProtocol{}
	for{
		p, err := echoProtocol.ReadPacket(conn)
		if err == nil {
			echoPacket := p.(*echo.EchoPacket)
			fmt.Printf("Server reply:[%v] [%v]\n", echoPacket.GetLength(), string(echoPacket.GetBody()))
		}
	}
	// }

	conn.Close()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
