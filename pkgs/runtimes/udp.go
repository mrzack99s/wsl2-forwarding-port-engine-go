package runtimes

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net"

	"github.com/mrzack99s/wsl2-forwarding-port-engine-go/pkgs/cmds"
	"github.com/mrzack99s/wsl2-forwarding-port-engine-go/pkgs/types"
)

func udpServe(packetFwd types.PacketForwarder, c types.ContextSession) {
	for {
		select {
		//if the context is done, then finish the operation
		case <-c.Ctx.Done():
			return
		default:
			listenAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("0.0.0.0:%d", packetFwd.WINPort))
			conn, err := net.ListenUDP("udp", listenAddr)
			if err != nil {
				fmt.Println(err)
			}
			buf := make([]byte, 65535)
			n, addr, err := conn.ReadFromUDP(buf)
			if err != nil {
				fmt.Println("Error Reading")
				return
			} else {
				wslAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", packetFwd.WSLIPAddr, packetFwd.WINPort))
				c, err := net.DialUDP("udp", nil, wslAddr)
				if err != nil {
					fmt.Println(err)
					return
				}

				_, err = c.Write(buf[0:n])
				if err != nil {
					fmt.Println(err)
					return
				}

				buffer := make([]byte, 65535)
				n, _, err := c.ReadFromUDP(buffer)
				if err != nil {
					fmt.Println(err)
					return
				}
				conn.WriteToUDP(buffer[0:n], addr)

			}

		}
	}
}

func CreateUDPForwarder(WINPort int, WSLPort int, force bool) (status bool) {
	ctx, ctxCancel := context.WithCancel(context.Background())
	status = false
	newUdp := types.PacketForwarder{
		Protocol:  "UDP",
		WINPort:   WINPort,
		WSLPort:   WSLPort,
		WSLIPAddr: cmds.GetWSLIP(),
	}
	newContext := types.ContextSession{
		Ctx:       ctx,
		CtxCancel: ctxCancel,
	}

	jsonByte, _ := json.Marshal(newUdp)
	sum := sha256.Sum256(jsonByte)
	hash8Id := hex.EncodeToString(sum[0:])[0:8]

	if !FindSession(hash8Id) || force {
		Session[hash8Id] = newContext
		RuleTables[hash8Id] = newUdp
		go udpServe(newUdp, newContext)
		WriteToFile()
		status = true
	}

	return

}
