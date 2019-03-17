package eureka

import (
	"log"
	"net"
	"os"
)

func GetLocalIP()  (string,  error){
	host, err := os.Hostname()
	if err != nil {
		log.Fatalf("Failed to get Hostname with error: %s", err)
		return "", err
	}
	addrs, err := net.LookupIP(host)
	if err != nil {
		log.Fatalf("Failed to get application IP with error: %s", err)
	}
	return addrs[len(addrs)-1].To4().String(),nil
}