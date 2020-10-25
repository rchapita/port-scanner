package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func PortScan(server string) []int {

	var available []int

	for i := 1; i <= 65535; i++ {
		ip := server + ":" + strconv.Itoa(i)

		_, err := net.DialTimeout("tcp", ip, time.Duration(300)*time.Millisecond)
		if err == nil {
			available = append(available, i)
		}
	}

	return available
}

func main() {

	fmt.Println("Checking for available ports...")
	ports := PortScan(os.Args[1])

	fmt.Println("Ports available: ", ports)

}
