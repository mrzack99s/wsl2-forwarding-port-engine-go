package types

import (
	"context"
)

var EngineVersion string = "1.0.0-go-beta"

type PacketForwarder struct {
	Protocol  string `json:"protocol"`
	WSLIPAddr string `json:"wsl_ip_addr"`
	WINPort   int    `json:"win_port"`
	WSLPort   int    `json:"wsl_port"`
}

type ContextSession struct {
	Ctx       context.Context
	CtxCancel context.CancelFunc
}

type Task struct {
	IPAddr string `json:"ip_addr"`
	Proto  string `json:"proto"`
	SPort  int    `json:"sport"`
	DPort  int    `json:"dport"`
}
