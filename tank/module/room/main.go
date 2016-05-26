package room

import (
	"fmt"
	"time"
)

//定义事件
type Event struct {
	Type      string `json:"Type"`
	Uid       string `json:"Uid"`
	Timestamp int    `json:"Timestamp"`
	Data      string `json:"Data"`
}

func Connect(uid string) *Player {
	user := &Player{}
	user.Uid = uid
	user.New = make(chan Event, 10)
	linking <- user
	return user
}

//玩家类
type Player struct {
	Uid      string
	Position string
	New      chan Event
}

//断开链接
func (s *Player) Cancel() {
	unlinking <- s
}

//新的事件
func newEvent(typ, user, msg string) Event {
	return Event{typ, user, int(time.Now().Unix()), msg}
}

//进入管道
func Join(user string) {
	Broadcast <- newEvent("join", user, "")
}

func Move(uid, data string) {
	user_list[uid].Position = data
	Broadcast <- newEvent("move", uid, data)
}

var (
	//正在连接列表
	linking = make(chan *Player, 10)
	//正在断开玩家列表
	unlinking = make(chan *Player, 10)
	//推送管道
	Broadcast = make(chan Event, 10)
	//在线玩家
	user_list = map[string]*Player{}
)

// This function loops forever, handling the chat room pubsub
func chatroom() {
	for {
		select {
		case user := <-linking:
			fmt.Println("一个叫" + user.Uid + "的萌坦克已经加入")
			user_list[user.Uid] = user
			user.Position = "{'x':'0','y':'0'}"
			event := Event{
				Type: "add",
				Uid:  user.Uid,
			}
			for _, v := range user_list {
				v.New <- event
				init_evnet := Event{
					Type: "init",
					Uid:  v.Uid,
					Data: v.Position,
				}
				user.New <- init_evnet
			}
		case event := <-Broadcast:
			for _, v := range user_list {
				v.New <- event
			}
		case user := <-unlinking:
			fmt.Println("一个叫" + user.Uid + "的萌坦克已经牺牲了")
			delete(user_list, user.Uid)
		}
	}
}

func init() {
	go chatroom()
}

// Drains a given channel of any messages.
func drain(ch <-chan Event) {
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				return
			}
		default:
			return
		}
	}
}
