package cg

import (
	"../ipc"
	"encoding/json"
	"errors"
	"sync"
)

//强转成IPC的server接口实例
var _ ipc.Server = &CenterServer{} // 确认实现了Server接口，最后面

//消息结构体，三个字段，发送人，接收人，内容
type Message struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Content string `json:"content"`
}

//中央服务器，slice，玩家，房间，读写控制锁
type CenterServer struct {
	servers map[string]ipc.Server
	players []*Player
	rooms   []*Room
	mutex   sync.RWMutex
}

//中央服务器的实例创建方法
func NewCenterServer() *CenterServer {
	servers := make(map[string]ipc.Server)
	players := make([]*Player, 0)

	//返回一个中央服务器实例，没有初始化房间和锁
	return &CenterServer{servers: servers, players: players}
}

//添加玩家，需要参数信息
func (server *CenterServer) addPlayer(params string) error {

	player := NewPlayer() //新建玩家

	//将json参数信息存入玩家实例中
	err := json.Unmarshal([]byte(params), &player)
	if err != nil {
		return err
	}

	//同步锁
	server.mutex.Lock()
	//延迟执行，函数返回的时候解开同步锁
	defer server.mutex.Unlock()

	// 偷懒了，没做重复登陆检查
	//追加一个玩家到通道中去
	server.players = append(server.players, player)

	return nil
}

//移出一名玩家，json参数，传入一个名字即可
func (server *CenterServer) removePlayer(params string) error {

	server.mutex.Lock()
	defer server.mutex.Unlock()

	//遍历当前的所有玩家
	for i, v := range server.players {

		//配置遍历的名字和传入的名字
		if v.Name == params {
			if len(server.players) == 1 { //只有一个玩家在线，清空slice
				server.players = make([]*Player, 0)
			} else if i == len(server.players)-1 { //最后一个玩家，取slice[0,i]
				server.players = server.players[:i]
			} else if i == 0 { //第一个玩家[1,size]
				server.players = server.players[1:]
			} else { //中间的情况,[0,i-1],[i,size]
				server.players = append(server.players[:i-1], server.players[:i+1]...)
			}
			return nil
		}
	}

	return errors.New("Player not found.")
}

//列出所有的玩家
func (server *CenterServer) listPlayer(params string) (players string, err error) {

	server.mutex.RLock()
	defer server.mutex.RUnlock()

	if len(server.players) > 0 {
		//生成json数据，返回
		b, _ := json.Marshal(server.players)
		players = string(b)
	} else {
		err = errors.New("No player online.")
	}
	return
}

//广播，向所有的玩家发送消息
func (server *CenterServer) broadcast(params string) error {

	var message Message
	//传入的json数据，存入消息实例中
	err := json.Unmarshal([]byte(params), &message)
	if err != nil {
		return err
	}

	server.mutex.Lock()
	defer server.mutex.Unlock()

	if len(server.players) > 0 {
		for _, player := range server.players {
			//循环的将消息发送到玩家消息通道中去
			player.mq <- &message
		}
	} else {
		err = errors.New("No player online.")
	}
	return err
}

//------------------------------实现server接口

//服务器处理数据方法实现，返回一个回应消息
func (server *CenterServer) Handle(method, params string) *ipc.Response {
	switch method {
	case "addplayer":
		err := server.addPlayer(params)
		if err != nil {
			return &ipc.Response{Code: err.Error()}
		}
	case "removeplayer":
		err := server.removePlayer(params)
		if err != nil {
			return &ipc.Response{Code: err.Error()}
		}
	case "listplayer":
		players, err := server.listPlayer(params)
		if err != nil {
			return &ipc.Response{Code: err.Error()}
		}
		return &ipc.Response{"200", players}
	case "broadcast":
		err := server.broadcast(params)
		if err != nil {
			return &ipc.Response{Code: err.Error()}
		}
		return &ipc.Response{Code: "200"}
	default:
		return &ipc.Response{Code: "404", Body: method + ":" + params}
	}
	return &ipc.Response{Code: "200"}
}

func (server *CenterServer) Name() string {
	return "CenterServer"
}
