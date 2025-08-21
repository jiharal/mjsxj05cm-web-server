package network

import (
	"os/exec"
	"strings"
)

func GetIP() string {
	// Try multiple methods to get IP address

	// Method 1: Use ip command (modern Linux)
	output, err := exec.Command("sh", "-c", "ip route get 1.1.1.1 | grep -oP 'src \\K\\S+' | head -1").Output()
	if err == nil && len(output) > 0 {
		return string(output)
	}

	// Method 2: Use hostname command
	output, err = exec.Command("hostname", "-I").Output()
	if err == nil && len(output) > 0 {
		// hostname -I can return multiple IPs, take the first one
		ip := string(output)
		if len(ip) > 0 {
			// Remove trailing whitespace and take first IP
			fields := strings.Fields(ip)
			if len(fields) > 0 {
				return fields[0]
			}
		}
	}

	// Method 3: Fallback to ifconfig for older systems
	output, err = exec.Command("sh", "-c", "ifconfig | grep -oP '(?<=inet\\s)\\d+(\\.\\d+){3}' | grep -v '127.0.0.1' | head -1").Output()
	if err == nil && len(output) > 0 {
		return string(output)
	}

	// If all methods fail, return localhost
	return "127.0.0.1"
}
