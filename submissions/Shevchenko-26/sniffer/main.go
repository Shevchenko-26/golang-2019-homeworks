package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

func main() {
	host := flag.String("host", "8.8.8.8", "required argument. Can be either domain name, like google.com or IP address like 172.217.3.110")
	portsRange := flag.String("ports", "1-1000", "the range of ports to scan. Should be formatted as the following of <start_port>-<end_port>")
	flag.Parse()

	var start, end = 0, 1 << 16
	var err error

	if *portsRange != "" {
		portsNums := strings.Split(*portsRange, "-")
		if len(portsNums) != 2 {
			log.Fatal("Ports range is invalid. Should be in form of <start_port>-<end_port>")
		}
		start, err = strconv.Atoi(portsNums[0])
		if err != nil {
			log.Fatal("Starting port was specified incorrectly. Aborting...")
		}
		endPort, err := strconv.Atoi(portsNums[1])
		if err != nil {
			log.Fatal("Second port was specified incorrectly. Aborting...")
		}
		if endPort > end {
			log.Fatal("Second port overflows maximum value of 65536. Aborting...")
		}
		end = endPort
	}

	var ports []int
	for i := start; i < end; i++ {

		_, err := net.DialTimeout("tcp",
			fmt.Sprintf("%s:%v", *host, i), time.Millisecond*300)
		if err != nil {
			if err, ok := err.(*net.OpError); ok {
				if !err.Err.(*net.DNSError).IsTimeout {
					log.Fatal("Problems connecting to the host. Aborting...")
				}
			}
			continue
		}
		fmt.Print(".")
		ports = append(ports, i)
	}
	if len(ports) != 0 {
		fmt.Printf("\n%v\n", ports)
	}
}
