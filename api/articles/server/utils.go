package server

import (
	"net"
	"os"
)

func getHostName() (name string) {
	name, _ = os.Hostname()
	return name
}

func getHostIP() (ip string) {
	listIP, _ := net.LookupHost(getHostName())
	if len(listIP) > 0 {
		ip = listIP[0]
	}
	return ip
}
