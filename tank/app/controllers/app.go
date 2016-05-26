package controllers

import (
	"code.google.com/p/go.net/websocket"
	"github.com/revel/revel"
	"tank/module/room"
	"encoding/json"
)

type App struct {
	*revel.Controller
}

type Event struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

var Global_event Event

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Socket(ws *websocket.Conn) revel.Result {
	uid := c.Session["name"]

	user := room.Connect(uid)
	defer user.Cancel()

	/*获取存储的事件  发送到新加入的人
	for _, event := range link.Archive {
		if websocket.JSON.Send(ws, &event) != nil {
			return nil
		}
	}*/

	//从客户端接收
	newEvent := make(chan string)
	go func() {
		var msg string
		for {
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				close(newEvent)
				return
			}
			json.Unmarshal([]byte(msg),&Global_event)
			switch Global_event.Type {
				case "move":
				room.Move(user.Uid,Global_event.Data)
				break
			}
		}
	}()

	//发送消息给客户端
	for {
		select {
		case event := <-user.New:
			if websocket.JSON.Send(ws, &event) != nil {
				return nil
			}
			/*
		case mag, ok := <-newEvent:
			if !ok {
				return nil
			}
			room.Do("广播")
			fmt.Println("一个萌萌的小坦克正在靠近，它正在",mag)*/
		}
	}
	return nil
}

func (c App) Login() revel.Result {
	return  c.Render()
}

func (c App) Loging(name string) revel.Result {
	c.Session["name"] = name
	return c.Render()
}

