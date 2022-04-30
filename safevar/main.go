/*
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex ///互斥锁
var rwmutex sync.RWMutex
var once sync.Once
var count int

///<互斥锁
func Add1() {
	mutex.Lock()
	defer mutex.Unlock()
	count++
}

///读写互斥锁
///只有锁竞争比较激烈，且绝大部分为读操作时，RWMutex 才有性能优势，否则不如 Mutex。
func Add2() {
	rwmutex.Lock()
	defer rwmutex.Unlock()
	count++
}
func Get2() int {
	rwmutex.RLock()
	defer rwmutex.RUnlock()
	return count
}

///<sync.Once 单例模式
///<保证只会调用一次
func printHello() {
	fmt.Printf("hello golang")
}
func Init() {
	once.Do(printHello)
}

///<channel
///<实现一个线程安全的map
type RWMap struct {
	Lock sync.RWMutex
	DMap map[string]string
}

func (r *RWMap) Get(key string) string {
	r.Lock.RLock()
	defer r.Lock.RUnlock()
	if value, ok := r.DMap[key]; ok {
		return value
	} else {
		return ""
	}
}
func (r *RWMap) Put(key, value string) {
	r.Lock.Lock()
	defer r.Lock.Unlock()
	if r.DMap == nil {
		r.DMap = make(map[string]string)
	}
	r.DMap[key] = value
}

///实现一个并发读写
func func1() {
	c := make(chan int, 10)
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background()) ////<当cancel被执行，将向ctx写入

	go func() {
		time.Sleep(1 * time.Minute)
		cancel()
		close(c)
		fmt.Println("cancel")
	}()
	///<写入数据
	for i := 0; i < 10; i++ {
		go func(ctx context.Context, id int) {

			select {
			case <-ctx.Done():
				fmt.Println("done............")
			case <-time.After(time.Second):

			case c <- id:
				fmt.Println("写入数据:", id)
			default:
				fmt.Println("........")
			}

		}(ctx, i)
	}
	///<读取数据
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range c {
				fmt.Println("读取数据:", v)
			}
		}()
	}
	fmt.Println("wait goroutine")
	wg.Wait() ///<等到所有的线程执行完成

	fmt.Println("all goroutine end")
}

func func2() {
	wg := sync.WaitGroup{}
	c := make(chan int, 10)
	done := make(chan bool)

	// (写端）
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			select {
			case <-done:
				fmt.Println("sender done...")
				return
			case c <- id:
				fmt.Println("写入数据: ", id)
				//default:
				//	fmt.Println("sender blocking...")
			}
		}(i)
	}

	// (读端）
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 。。。 处理channel 中的数据
			for {
				select {
				case <-done:
					fmt.Println("receiver done...")
					return
				case v := <-c:
					fmt.Println("receiver get :", v)
					time.Sleep(time.Millisecond)
				}
			}
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("close")
	close(done)
	// 等待所有的读端完成
	wg.Wait()
}

///<生产者
func Product(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

///<消费者
func Consumer(ch <-chan int, done chan bool) {
	for i := 0; i < 9; i++ {
		a := <-ch
		fmt.Println(a)
	}
	close(done)
}
func main() {
	ch := make(chan int, 10)

	done := make(chan bool)
	go Product(ch)
	go Consumer(ch, done)
	for {
		select {
		case <-done:

			fmt.Println("执行完毕")
			return
		}
	}
}
*/

/*
func func3() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg: ///<因为没有人接收msg，所以不会执行
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	huang := &Person{
		Name: "huangliangqiu",
		Age:  10,
	}
	liang := &Person{
		Name: "huangliangqiu",
		Age:  10,
	}
	if huang == liang {
		fmt.Println("same")
	}
}

package main

import (
	"fmt"
)

var slice []string = make([]string, 0, 1)

func main() {
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	slice = append(slice, "huang")
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	slice = append(slice, "liang")
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	slice = append(slice, "qiu")
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	mp := make(map[int]string, 1)
	slice := make([]int, 0, 1)
	mp[1] = "huang"
	mp[3] = "liang"
	mp[4] = "qiu"
	mp[2] = "zhang"
	for key, value := range mp {
		fmt.Println(key, value)
		slice = append(slice, key)
	}
	//	sort.Ints(slice)
	//	sort.Sort(sort.Reverse(sort.IntSlice(slice))) ///<逆序
	sort.Sort(sort.IntSlice(slice))
	for _, value := range slice {
		fmt.Println(value, mp[value])
	}
}


package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string // 姓名
	Age  int    // 年纪
}

// 按照 Person.Age 从大到小排序
type PersonSlice []Person

func (a PersonSlice) Len() int { // 重写 Len() 方法
	return len(a)
}
func (a PersonSlice) Swap(i, j int) { // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a PersonSlice) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return a[i].Age < a[j].Age
}
func main() {
	people := []Person{
		{"zhang san", 12},
		{"li si", 30},
		{"wang wu", 52},
		{"zhao liu", 26},
	}
	fmt.Println(people)
	sort.Sort(PersonSlice(people)) // 按照 Age 的顺序
	fmt.Println(people)

	sort.Sort(sort.Reverse(PersonSlice(people))) // 按照 Age 的降序排序
	fmt.Println(people)

}


package main

import (
	"fmt"
)

func main() {
	slice1 := make([]int, 0)            ///<非nil
	slice2 := []int{}                   ///<非nil
	var slice3 []int                    ///<nil

	if nil == slice1 {
		fmt.Println("make([]int, 0) is nil")
	} else {
		fmt.Println("make([]int, 0) not nil")
	}
	if nil == slice2 {
		fmt.Println("[]int{} is nil")
	} else {
		fmt.Println("[]int{} not nil")
	}
	if nil == slice3 {
		fmt.Println("var slice3 []int is nil")
	} else {
		fmt.Println("var slice3 []int not nil")
	}
}


package main

import (
	"fmt"
	"sync"
)

///<once.Do(f func())
var once sync.Once

func printAddr(ip, port string) {
	fmt.Printf("server:%s, port:%s\n", ip, port)
}

///<辅助函数
func auxFunc(f func(string, string), ip, port string) func() {
	return func() {
		f(ip, port)
	}
}
func Init() {
	once.Do(auxFunc(printAddr, "127.0.0.1", "8080"))
}

func main() {
	Init()
}


package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func ReadFileByIO() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer file.Close()
	var buffer [512]byte
	for {
		n, err := file.Read(buffer[:])
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read file failed,err:%v\n", err)
			return
		}
		fmt.Println(string(buffer[:n]))
	}
}

func WriteFileByIO(msg string) {
	file, err := os.OpenFile("./file.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Printf("OpenFile failed,err:%v\n", err)
	}
	defer file.Close()
	_, err = file.Write([]byte(msg))
	if err != nil {
		fmt.Printf("Write file failed, err:%v\n", err)
		return
	}
}

func ReadFileByBufio() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read string failed, err:%v\n", err)
			return
		}
		fmt.Println(line)
	}
}
func WriteFileByBufio(msg string) {
	file, err := os.OpenFile("./file.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Printf("OpenFile failed,err:%v\n", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString(msg)
	writer.Flush() ///<从缓冲区写入到文件
}
func ReadFileByIoutil() {
	ret, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("ReadFile failed,err :%v\n", err)
		return
	}
	fmt.Println(string(ret))
}
func WriterFileByIoutil(msg string) {
	err := ioutil.WriteFile("./file.txt", []byte(msg), 0777)
	if err != nil {
		fmt.Printf("WriteFile failed,err :%v\n", err)
	}
}
func main() {
	//	ReadFileByIO()
	//	ReadFileByBufio()
	//	ReadFileByIoutil()
	//	WriteFileByIO("WriteFileByIO")
	//	WriteFileByBufio("WriteFileByBufio")
	WriterFileByIoutil("WriterFileByIoutil")
}


package main

import (
	"bufio"
	"fmt"
	"os"
)

func GetInput() {
	fmt.Println("请输出内容:")
	var s string
	_, err := fmt.Scanln(&s)
	if err != nil {
		fmt.Printf("Scanln failed, err:%v\n", err)
	}
	fmt.Printf("输入的内容是:%s\n", s)
}

func GetInputByBufio() {
	fmt.Println("请输入内容:")
	var s string
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("ReadString failed,err:%v", err)
		return
	}
	fmt.Printf(s)

}
func main() {
	GetInputByBufio()
	//	GetInput()
}


package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			m.Store(n, strconv.Itoa(n))
			value, _ := m.Load(n)
			fmt.Printf("key:%d, value:%s\n", n, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
	m.Delete(1)
	fmt.Println("----------------------------------------")
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("key:%d, value:%s\n", key, value)
		return true
	})
}


package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var Num int32
var mutex sync.Mutex

func Add() {
	Num++
}
func AddMutex() {
	mutex.Lock()
	defer mutex.Unlock()
	Num++
}
func AddAtomic() {
	atomic.AddInt32(&Num, 1)
}

func main() {
	start := time.Now()
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			//Add() ///<Num 965  durtime 1099400
			//AddMutex() ///<Num 1000  durtime 1588100
			AddAtomic() ///<Num 1000 durtime 1153900
			wg.Done()
		}()
	}
	wg.Wait()
	end := time.Now()
	fmt.Printf("durtime:%d\n", end.Sub(start))
	fmt.Println(Num)
}

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

// Encode 将消息编码
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

// Decode 解码消息
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
		message, _ := Encode(msg)
		conn.Write(message)
		if msg == "exit" {
			break
		}
	}
	conn.Close()
}

func main() {
	server()
	//	client()
}

package main

import (
	"fmt"
	"net"
)

// UDP server端
func server() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:]) // 接收数据
		if err != nil {
			fmt.Println("read udp failed, err:", err)
			continue
		}
		fmt.Printf("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)
		_, err = listen.WriteToUDP(data[:n], addr) // 发送数据
		if err != nil {
			fmt.Println("write to udp failed, err:", err)
			continue
		}
	}
}
func client() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("连接服务端失败，err:", err)
		return
	}
	defer socket.Close()
	sendData := []byte("Hello server")
	_, err = socket.Write(sendData) // 发送数据
	if err != nil {
		fmt.Println("发送数据失败，err:", err)
		return
	}
	data := make([]byte, 4096)
	n, remoteAddr, err := socket.ReadFromUDP(data) // 接收数据
	if err != nil {
		fmt.Println("接收数据失败，err:", err)
		return
	}
	fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), remoteAddr, n)
}
func main() {
	//    client()
	server()
}


package main

import (
	"fmt"
	"strings"
)

func main() {
	var builder strings.Builder
	builder.WriteString("huang")
	builder.WriteString("liang")
	builder.WriteString("qiu")
	str := builder.String()
	fmt.Println(str)
}


package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	userInfo := Person{
		Name: "Bill",
		Age:  25,
	}
	// 结构体打印(json格式等...)
	fmt.Printf("%+v\n", userInfo) // {Name:Bill Age:25}
	fmt.Printf("%v\n", userInfo)  // {Bill 25}
}

package main

import (
	"fmt"
	"reflect"
)

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	mp1 := map[int]string{1: "huangliangqiu"}
	mp2 := map[int]string{1: "huangliangqiu"}
	///<无法比较

//		if slice1 == slice2 {
//			fmt.printf("same")
//		}
//		if mp1 == mp2 {
//			fmt.Println("Same")
//		}
//
	isSame := reflect.DeepEqual(slice1, slice2)
	if isSame {
		fmt.Println("Same")
	}
	isSame = reflect.DeepEqual(mp1, mp2)
	if isSame {
		fmt.Println("Same")
	}
}
*/

package main
