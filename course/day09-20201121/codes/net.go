package main

import (
	"fmt"
	"net"
)

func main() {
	// ipv4, ipv6, 主机名/域名
	fmt.Println(net.JoinHostPort("127.0.0.1", "8080"))
	fmt.Println(net.SplitHostPort("127.0.0.1:888"))

	// IP -> 主机名
	fmt.Println(net.LookupAddr("127.0.0.1"))
	fmt.Println(net.LookupAddr("212.64.62.183"))
	// 主机名 -> IP
	fmt.Println(net.LookupHost("www.baidu.com"))
	fmt.Println(net.LookupHost("www.oschina.com"))
	fmt.Println(net.LookupHost("localhost"))

	// IP, IPNet
	for _, ipStr := range []string{"127.0.0.1", "::1", "xxxx"} {
		ip := net.ParseIP(ipStr)
		fmt.Printf("%#v\n", ip)
	}

	ip, ipnet, err := net.ParseCIDR("192.168.0.0/24")
	// 192.168.0.0-192.168.0.255
	fmt.Printf("%#v %#v %#v\n", ip, ipnet, err)

	fmt.Println(ipnet.Contains(net.ParseIP("192.168.1.1")))
	fmt.Println(ipnet.Contains(net.ParseIP("192.168.0.10")))

	addrs, err := net.InterfaceAddrs()
	fmt.Println(addrs, err)

	intfs, err := net.Interfaces()
	fmt.Println(intfs, err)
	for _, intf := range intfs {
		fmt.Println(intf.Index, intf.Name, intf.MTU, intf.HardwareAddr, intf.Flags)
		fmt.Println(intf.Addrs())
	}
}
