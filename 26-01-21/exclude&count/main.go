package main

import "fmt"

func main() {

	ips := []string{
		"192.168.1.1",
		"10.0.0.1",
		"192.168.1.1",
		"192.168.1.2",
		"10.0.0.1",
		"192.168.1.1",
	}

	counts := make(map[string]int)

	for _, ip := range ips {
		if ip == "127.0.0.1" {
			continue
		}
		counts[ip]++
	}

	fmt.Println(counts)
}
