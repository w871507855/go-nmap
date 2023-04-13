package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

func scanPortSync(ip string, port int, wg *sync.WaitGroup) {
	defer wg.Done()
	address := fmt.Sprintf("%s:%d", ip, port)
	d := net.Dialer{Timeout: 5 * time.Second}
	conn, err := d.Dial("tcp", address)
	if err == nil {
		fmt.Printf("%s的%d端口是开放的\n", ip, port)
		conn.Close()
	}
}

func sanPort(ip string, port int) {
	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.Dial("tcp", address)
	if err == nil {
		fmt.Printf("%s的%d端口是开放的\n", ip, port)
		conn.Close()
	} else {
		fmt.Printf("%s的%d端口是关闭的\n", ip, port)
	}
}

func main() {
	ip := flag.String("ip", "127.0.0.1", "IP地址")
	port := flag.Int("port", 0, "端口")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s: \n", os.Args[0])
		flag.PrintDefaults()
		fmt.Println("go-nmap提示")
	}
	flag.Parse()
	if *port == 0 {
		wg := new(sync.WaitGroup)
		for p := 1; p <= 65535; p++ {
			wg.Add(1)
			go scanPortSync(*ip, p, wg)
		}
		wg.Wait()
	} else {
		sanPort(*ip, *port)
	}
}
