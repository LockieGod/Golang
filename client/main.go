package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
)

func Encode(message string) ([]byte, error) {
	// 读取消息的长度，转换成int32类型（占4个字节）
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)
	// 写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	// 写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil

}
func Decode(reader *bufio.Reader) (string, error) {
	// 读取消息的长度
	lengthByte, _ := reader.Peek(4) // 读取前4个字节的数据
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	// Buffered返回缓冲中现有的可读取的字节数。
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}

	// 读取真正的消息数据
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}
///《服务端
func server() {
	///<1 本地端口启动服务
	server, err := net.Listen("tcp", "127.0.0.1:10000")
	if err != nil {
		fmt.Println("Listen failed, err:", err)
		return
	}
	///<2 等待别人建立连接
	for {
		conn, err := server.Accept()
		go func(conn net.Conn) {

			if err != nil {
				fmt.Printf("Accept failed,err:%v", err)
				return
			}
			reader := bufio.NewReader(conn)
			///<3 与客户端通信
			for {
				msg, err := Decode(reader)
				if err == io.EOF {
					break
				}
				if msg == "exit" {
					break
				}
				fmt.Println(msg)
			}
			conn.Close()
		}(conn)
	}
}
//<客户端
func client() {
	conn, err := net.Dial("tcp", "127.0.0.1:10000")
	if err != nil {
		fmt.Println("Dial failed ,err:", err)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("please enter msg")
		msg, _ := reader.ReadString('\n')
		fmt.Println(msg)
		message, _ := Encode(msg)
		fmt.Println(string(message))
		conn.Write(message)
		if msg == "exit" {
			break
		}
	}
	conn.Close()
}

func main() {
	//server()
	client()
}
