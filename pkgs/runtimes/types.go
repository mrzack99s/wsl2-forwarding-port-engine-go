package runtimes

import "github.com/mrzack99s/wsl2-forwarding-port-engine-go/pkgs/types"

var Session = make(map[string]types.ContextSession)
var RuleTables = make(map[string]types.PacketForwarder)

func FindSession(hash8Id string) bool {
	if _, found := Session[hash8Id]; found {
		return true
	}

	return false
}

func FindRuleTables(hash8Id string) bool {
	if _, found := RuleTables[hash8Id]; found {
		return true
	}

	return false
}
