package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func checkWebsiteHealth(url string, timeout time.Duration) (bool, error) {
	// Ensure URL includes a port
	if !strings.Contains(url, ":") {
        url = url + ":80" // Default to HTTP port
    }

	conn, err := net.DialTimeout("tcp", url, timeout)
	if err != nil {
		fmt.Printf("Error while creating connection to website")
		return false, err
	}
	conn.Close()
	return true, nil
}
