/**
 * 服务端
 * @package       main
 * @author        YuanZhiGang <zackyuan@yeah.net>
 * @version       1.0.0
 * @copyright (c) 2013-2023, YuanZhiGang
 */

package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func process(conn net.Conn) {
	defer conn.Close()

	for {
		Reader := bufio.NewReader(conn) // 创建一个读取器
		var temp [1024]byte
		n, err := Reader.Read(temp[:]) // 读入temp，返回一个读取了多少个字节

		if err != nil {
			fmt.Println("read from conn failed,err:", err)
			break
		}

		// 从temp中取出指定字节的数据(temp[:n]),并将之转换成字符串
		restr := strings.ToUpper(string(temp[:n]))
		fmt.Println("收到client端发送的数据:", restr)
		_, err = conn.Write([]byte(restr))
		if err != nil {
			fmt.Println("server return data to client  failed,err:", err)
			break
		}
	}
}

func main() {
	listen, err := net.Listen("tcp", "localhost:30000")
	if err != nil {
		fmt.Println("tcp listen localhost:30000 failed, err:", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			return
		}

		go process(conn)
	}
}
