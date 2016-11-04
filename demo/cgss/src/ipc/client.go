package ipc

import (
	"encoding/json"
)

type IpcClient struct {
	conn chan string
}

//IPC通信客户端，传入一个ipc服务端接口实例，返回一个session通道
func NewIpcClient(server *IpcServer) *IpcClient {
	c := server.Connect()

	return &IpcClient{c}
}

//客户端连接操作，传入方法，参数，返回回复信息，错误信息
func (client *IpcClient) Call(method, params string) (resp *Response, err error) {

	//封装一个请求信息
	req := &Request{method, params}

	var b []byte
	//将对象转化为json数据
	b, err = json.Marshal(req)
	if err != nil {
		return
	}

	//将json数据传入通道中去也就是session中
	client.conn <- string(b)
	//阻塞
	str := <-client.conn // 等待返回值

	var resp1 Response
	//拿到响应的信息
	err = json.Unmarshal([]byte(str), &resp1)
	resp = &resp1

	return
}

func (client *IpcClient) Close() {
	client.conn <- "CLOSE"
}
