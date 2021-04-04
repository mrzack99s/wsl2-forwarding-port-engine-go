package runtimes

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/mrzack99s/wsl2-forwarding-port-engine-go/pkgs/cmds"
	"github.com/mrzack99s/wsl2-forwarding-port-engine-go/pkgs/types"
)

func handleConnection(c net.Conn, packetFwd types.PacketForwarder) {
	client, err := net.Dial("tcp", fmt.Sprintf("%s:%d", packetFwd.WSLIPAddr, packetFwd.WSLPort))
	if err != nil {
		log.Printf("Dial failed: %v", err)
		defer c.Close()
		return
	}
	go func() {
		defer client.Close()
		defer c.Close()
		io.Copy(client, c)
	}()
	go func() {
		defer client.Close()
		defer c.Close()
		io.Copy(c, client)
	}()
}

func tcpServe(packetFwd types.PacketForwarder, c types.ContextSession) {
	for {
		select {
		//if the context is done, then finish the operation
		case <-c.Ctx.Done():
			return
		default:
			listenAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("0.0.0.0:%d", packetFwd.WINPort))
			ln, err := net.ListenTCP("tcp", listenAddr)
			if err != nil {
				fmt.Println(err)
			}

			conn, err := ln.Accept()
			if err != nil {
				continue
			}
			log.Println("Client ", conn.RemoteAddr(), " connected")
			go handleConnection(conn, packetFwd)

		}
	}
}

func CreateTCPForwarder(WINPort int, WSLPort int, force bool) (status bool) {
	ctx, ctxCancel := context.WithCancel(context.Background())
	newTcp := types.PacketForwarder{
		Protocol:  "TCP",
		WINPort:   WINPort,
		WSLPort:   WSLPort,
		WSLIPAddr: cmds.GetWSLIP(),
	}
	newContext := types.ContextSession{
		Ctx:       ctx,
		CtxCancel: ctxCancel,
	}

	status = false

	jsonByte, _ := json.Marshal(newTcp)
	sum := sha256.Sum256(jsonByte)
	hash8Id := hex.EncodeToString(sum[0:])[0:8]

	if !FindSession(hash8Id) || force {
		Session[hash8Id] = newContext
		RuleTables[hash8Id] = newTcp
		go tcpServe(newTcp, newContext)
		WriteToFile()
		status = true

	}

	return

}
