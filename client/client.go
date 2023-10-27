/**
 * 客户端
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
	"os"
	"strings"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:30000")

	if err != nil {
		fmt.Println("client Dial localhost:30000 failed,err:", err)
		return
	}

	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin) // 构建读取器，从os.Stdin(控制台，标准输入)读取数据

	// 一直可以接收/发送
	fmt.Println("请说话....")
	for {
		inputStr, err := inputReader.ReadString('\n') // 读取数据知道换行为止
		if err != nil {
			fmt.Println("os.Stdin 输入数据异常，err:", err)
			return
		}
		// 对数据进行处理
		nowInputStr := strings.Trim(inputStr, "\r\n") // 去除字符串首位的空格

		if nowInputStr == "exit" { // 退出
			return
		}

		// 将获取到的数据发送到服务端
		_, err = conn.Write([]byte(nowInputStr))
		if err != nil {
			fmt.Println("client send data failed,err:", err)
			return
		}
		// 接收服务端返回数据
		var temp [1024]byte
		server_reader := bufio.NewReader(conn)
		num, err := server_reader.Read(temp[:])
		if err != nil {
			fmt.Println("got server return data failed, err:", err)
		}

		fmt.Println("服务端:", string(temp[:num]))

	}

}
