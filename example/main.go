package main

import (
	"fmt"

	"github.com/fufuok/myip"
)

func main() {
	fmt.Println("获取外网地址 (IPv4):", myip.ExternalIPv4())
	fmt.Println("获取外网地址 (IPv6):", myip.ExternalIPv6())
	fmt.Println("获取外网地址 (出口公网地址):", myip.ExternalIP(""))
	fmt.Println("获取外网地址 (出口公网地址 IPv4):", myip.ExternalIP("ipv4"))
	fmt.Println("获取外网地址 (出口公网地址 IPv6):", myip.ExternalIP("ipv6"))

	fmt.Println("获取内网地址 (IPv4):", myip.InternalIPv4())
	fmt.Println("获取内网地址 (临时 IPv6 地址):", myip.InternalIPv6())
	fmt.Println("获取内网地址 (出口本地地址):", myip.InternalIP("", ""))
	fmt.Println("获取内网地址 (出口本地地址):", myip.InternalIP("1.1.1.1:53", "udp"))
	fmt.Println("获取内网地址 (出口本地地址):", myip.InternalIP("baidu.com:443", "tcp"))
	fmt.Println("获取内网地址 (出口本地地址):", myip.InternalIP("1.1.1.1", "ip4:icmp"))

	fmt.Println("获取本地地址 (第一个):", myip.LocalIP())
}
