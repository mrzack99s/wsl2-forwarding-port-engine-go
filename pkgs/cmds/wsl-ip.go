package cmds

import (
	"os/exec"
	"strings"
)

func GetWSLIP() string {
	cmds := exec.Command("bash.exe", "-c", "ip addr show eth0 | grep -oP '(?<=inet\\s)\\d+(\\.\\d+){3}'")
	output, _ := cmds.Output()
	outputStr := strings.TrimSpace(string(output))
	return outputStr
}
