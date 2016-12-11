// 客户端

package main

import(
    "net"
    "fmt"
)

func main() {
    //ip := net.ParseIP("127.0.0.111")
    //mask := ip.DefaultMask()
    //ip, _ := net.ResolveIPAddr("ip6", "www.baidu.com")
    //ip, _ := net.LookupIP("www.eumst.com")
    // ip, _ := net.LookupAddr("www.baidu.com")
    ip, _ := net.ResolveTCPAddr("tcp", "www.baidu.com:80")
    fmt.Println(ip)
}