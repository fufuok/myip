package myip

import (
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
	"unsafe"
)

var externalIPAPI = map[string][]string{
	"ipv4": {
		"http://api.ipify.org",
		"http://ifconfig.me/ip",
		"http://ident.me",
		"http://ip.cip.cc",
		"http://myexternalip.com/raw",
		"http://ip.42.pl/short",
	},
	"ipv6": {
		"http://api64.ipify.org",
	},
}

// 获取外网地址 (IPv4)
func ExternalIPv4() string {
	return ExternalIP("ipv4")
}

// 获取外网地址 (IPv6)
func ExternalIPv6() string {
	if ip := ExternalIP("ipv6"); ip != "" && strings.Count(ip, ":") > 1 {
		return ip
	}

	return ""
}

// 获取外网地址 (出口公网地址)
func ExternalIP(v string) string {
	if v != "ipv6" {
		v = "ipv4"
	}

	for _, u := range externalIPAPI[v] {
		if ip := getExternalIP(u); ip != "" {
			return ip
		}
	}

	return ""
}

// 请求 API 获取公网 IP
func getExternalIP(u string) string {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get(u)
	if err != nil {
		return ""
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	ip := net.ParseIP(strings.TrimSpace(*(*string)(unsafe.Pointer(&b))))
	if ip != nil {
		return ip.String()
	}

	return ""
}

// 获取内网地址 (IPv4)
func InternalIPv4() string {
	return InternalIP("", "udp4")
}

// 获取内网地址 (临时 IPv6 地址)
func InternalIPv6() string {
	return InternalIP("[2001:4860:4860::8888]:53", "udp6")
}

// 获取内网地址 (出口本地地址)
func InternalIP(dstAddr, network string) string {
	if dstAddr == "" {
		dstAddr = "8.8.8.8:53"
	}
	if network == "" {
		network = "udp"
	}

	conn, err := net.DialTimeout(network, dstAddr, time.Second)
	if err != nil {
		return ""
	}

	defer func() {
		_ = conn.Close()
	}()

	addr := conn.LocalAddr().String()
	ip := net.ParseIP(addr).String()
	if ip == "<nil>" {
		ip, _, _ = net.SplitHostPort(addr)
	}

	return ip
}

// 获取本地地址 (第一个)
func LocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err == nil {
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					return ipnet.IP.String()
				}
			}
		}
	}

	return ""
}
