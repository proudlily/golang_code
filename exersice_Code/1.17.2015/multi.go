
//channel 本身身也可以作为数据发送给其他的 channel。

package main

import "fmt"
type Data struct {
	arg  []int
	ch chan string
}
func server() chan *Data{
	reqs:=make(chan *Data)// 服务器 Request Channel

	go func(){ // 使用独立的 goroutine 循环处理所有的客户端请求
	    for d :=range reqs{// 循环从 Request Channel 获取客户端请求数据
			go serverProcess(d)// 将请求的客户端数据交给另外的 goroutine 处理,并立立即等待下一一个请求。
		}
	}()

	return reqs// 返回 Request Channel,用于 close server
}
func serverProcess(data *Data){
	x:=0
	for i :=range data.arg{// 统计用用户请求数据
		x+=i
	}
	s:=fmt.Sprintf("server:%d",x)
	data.ch<-s // 使用用客户端请求数据包中的 Channel,返回结果给客户端。
}
func main(){
	/* 启动服务器 */
	serverReqs:=server()
	/* 客户端向服务器发送请求数据,并等待返回结果 */
	data:=&Data{[]int{1,2,3},make(chan string)}
	serverReqs<-data// 发送请求到服务器
	println(<-data.ch) // 获取服务器返回结果

	/* 关闭服务器 */
	close(serverReqs)
}

