package runtimes

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/mrzack99s/wsl2-forwarding-port-engine-go/pkgs/cmds"
	"github.com/mrzack99s/wsl2-forwarding-port-engine-go/pkgs/types"
)

func StopForwarder(hash8Id string) (status bool) {
	status = false
	if FindSession(hash8Id) {
		defer Session[hash8Id].CtxCancel()
		delete(Session, hash8Id)
		delete(RuleTables, hash8Id)
		WriteToFile()
		status = true
	}
	return
}

func StartChecking() {
	for key, element := range RuleTables {
		switch element.Protocol {
		case "UDP":
			if element.WSLIPAddr != cmds.GetWSLIP() {
				CreateUDPForwarder(element.WINPort, element.WSLPort, true)
				delete(RuleTables, key)
			} else {
				CreateUDPForwarder(element.WINPort, element.WSLPort, true)
			}

		case "TCP":
			if element.WSLIPAddr != cmds.GetWSLIP() {
				CreateTCPForwarder(element.WINPort, element.WSLPort, true)
				delete(RuleTables, key)
			} else {
				CreateTCPForwarder(element.WINPort, element.WSLPort, true)
			}
		}
	}
}

func MGMTServe() {
	for {
		listenAddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:40123")
		conn, err := net.ListenUDP("udp", listenAddr)
		if err != nil {
			fmt.Println(err)
		}
		for {
			buf := make([]byte, 65535)
			n, addr, err := conn.ReadFromUDP(buf)
			if err != nil {
				fmt.Println(err)
				return
			}
			recvDataSplit := strings.Split(string(buf[0:n]), "@")

			switch recvDataSplit[0] {
			case "create":
				sport, _ := strconv.Atoi(recvDataSplit[2])
				dport, _ := strconv.Atoi(recvDataSplit[2])
				newTask := types.Task{
					IPAddr: addr.IP.String(),
					Proto:  recvDataSplit[1],
					SPort:  sport,
					DPort:  dport,
				}
				switch newTask.Proto {
				case "UDP":
					if CreateUDPForwarder(sport, dport, false) {
						conn.WriteToUDP([]byte("SUCCESS"), addr)
					} else {
						conn.WriteToUDP([]byte("ALREADY"), addr)
					}
				case "TCP":
					if CreateTCPForwarder(sport, dport, false) {
						conn.WriteToUDP([]byte("SUCCESS"), addr)
					} else {
						conn.WriteToUDP([]byte("ALREADY"), addr)
					}
				}
			case "delete":
				id := recvDataSplit[1]
				if StopForwarder(id) {
					conn.WriteToUDP([]byte("SUCCESS"), addr)
				} else {
					conn.WriteToUDP([]byte("ALREADY"), addr)
				}

			case "purge":
				if recvDataSplit[1] == "Y" {
					for key, _ := range Session {
						StopForwarder(key)
						delete(Session, key)
					}
					WriteToFile()
					conn.WriteToUDP([]byte("SUCCESS"), addr)
				}
			case "get":
				switch recvDataSplit[1] {
				case "ls":
					strTasks := ""
					i := 0
					for key, element := range RuleTables {
						if i == len(RuleTables)-1 {
							strTasks += key + "@" + element.WSLIPAddr + "@" + element.Protocol + "@" +
								fmt.Sprintf("%d", element.WINPort) + "@" + fmt.Sprintf("%d", element.WSLPort)
						} else {
							strTasks += key + "@" + element.WSLIPAddr + "@" + element.Protocol + "@" +
								fmt.Sprintf("%d", element.WINPort) + "@" + fmt.Sprintf("%d", element.WSLPort) + "@@"
						}
						i++
					}

					if len(RuleTables) > 0 {
						conn.WriteToUDP([]byte(strTasks), addr)
					} else {
						conn.WriteToUDP([]byte("FAILLED"), addr)
					}
				case "engine_version":
					conn.WriteToUDP([]byte(types.EngineVersion), addr)
				}

			}
		}
	}

}
