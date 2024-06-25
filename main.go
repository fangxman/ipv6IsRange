package main

import (
    "os"
	"net"
    "flag"
)

func check_ipv6_is_ipv6range(ip string, cidrRange string) int {

    // 将测试地址解析为IP对象
    ipAddress := net.ParseIP(ip)
    if ipAddress == nil {
        return 0
    }

    ipAddress2 := net.ParseIP(cidrRange)

    // 判断是相同的
    if ipAddress.Equal(ipAddress2) {
        return 1
    }

    // 解析CIDR范围并创建一个IPNet对象
    _, ipNet, err := net.ParseCIDR(cidrRange)
    if err != nil {
        return 0
    }

    // 判断测试地址是否在范围内
    if ipNet.Contains(ipAddress) {
        return 1
    }
    return 0
}

func main() {
    // 解析命令行参数
    var testAddress string
    flag.StringVar(&testAddress, "ipv6", "", "ipv6地址")
    var ipv6range string
    flag.StringVar(&ipv6range, "ipv6range", "", "ipv6范围")
    flag.Parse()
    var num = check_ipv6_is_ipv6range(testAddress, ipv6range)
    os.Exit(num)
}
