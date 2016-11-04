package ipc

import (
	"encoding/json"
	"fmt"
)

//请求信息
type Request struct {
	Method string `json:"method"`
	Params string `json:"params"`
}

//回复信息
type Response struct {
	Code string `json:"code"`
	Body string `json:"body"`
}

//服务接口
type Server interface {
	Name() string
	Handle(method, params string) *Response
}

type IpcServer struct {
	Server
}

//ipc通信服务端，需要传入一个实现了server接口的对象
func NewIpcServer(server Server) *IpcServer {
	return &IpcServer{server}
}

//链接通信，server接口的方法
func (server *IpcServer) Connect() chan string {
	session := make(chan string, 0)

	go func(c chan string) {
		for {
			//阻塞，直到通道中有消息为止
			request := <-c

			if request == "CLOSE" { // 关闭该连接
				break
			}

			var req Request
			//解析拿到的json数据，存入Request中
			err := json.Unmarshal([]byte(request), &req)
			if err != nil {
				fmt.Println("Invalid request format:", request)
			}

			//调用处理函数，拿到返回的结果
			resp := server.Handle(req.Method, req.Params)

			//解析json数据
			b, err := json.Marshal(resp)

			c <- string(b) // 返回结果
		}

		fmt.Println("Session closed.")

	}(session)

	fmt.Println("A new session has been created successfully.")

	return session
}
