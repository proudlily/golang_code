package models

import (
	"github.com/revel/revel"
	"labix.org/v2/mgo"
	"fmt"

)

var (
	DB*mgo.Database
	Updir string
)

//初始化数据库
func Init() {
	fmt.Println("dburl")
	dburl, _ := revel.Config.String("dburl")//从app.conf取出dburl
	session, err := mgo.Dial(dburl)//启动mongo
	if err != nil {
		panic("无法连接数据库")
	}else {
		fmt.Println("初始化munix数据库成功")
	}
	session.SetMode(mgo.Monotonic, true)
	DB = session.DB("munix")

	//Updir, _ = revel.Config.String("updir")  //初始化上传的路径

}

/*ajax返回数据*/
type AjaxRel struct {
	Status int
	mag    string
	Data   string

}
