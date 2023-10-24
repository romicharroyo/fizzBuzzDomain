package main

import (
	"fmt"
	"net"
	"os"
	"sort"
)

var ipv4List []net.IP
var sumList []int

func ipSum(ip net.IP) int {
	sum := 0
	for _, part := range ip.To4() {
		sum += int(part)
	}
	return sum
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println(os.Stderr, "ERROR: you need to provide one arg with domain (ex. google.com)")
		os.Exit(1)
	}

	domain := os.Args[1]
	ips, err := net.LookupIP(domain)
	if err != nil {
		fmt.Println("ERROR: could not to get A records")
		return
	}
	fmt.Println(ips)

	for _, ip := range ips {
		if v4 := ip.To4(); v4 != nil {
			ipv4List = append(ipv4List, v4)
		}
	}

	sumSet := make(map[int]struct{})
	for _, ip := range ipv4List {
		sum := ipSum(ip)
		if _, exists := sumSet[sum]; !exists {
			sumList = append(sumList, sum)
			sumSet[sum] = struct{}{}
		}
	}

	sort.Ints(sumList)
	for _, i := range sumList {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%5 == 0:
			fmt.Println("Buzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		default:
			fmt.Println(i)
		}
	}
}
