package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/revel/revel"
	"labix.org/v2/mgo/bson"
	"new/app/models"
)

type User struct {
	*revel.Controller
}

func (c User) Register() revel.Result {
	return c.Render()
}

//插入用户到数据库
func (c User) Regist(name, password string) revel.Result {
	c.Validation.Required(name).Message("请输入名字")
	c.Validation.Required(password).Message("请输入密码")
	c.Validation.MaxSize(name, 15).Message("名字不超过15字节")
	c.Validation.MinSize(name, 4).Message("名字不少于4字节")

	rels := &models.AjaxRel{}

	if c.Validation.HasErrors() {
		errs := c.Validation.ErrorMap()
		rels.Data = "["
		for key, value := range errs {
			rels.Data = rels.Data + "{'key':'" + key + "','mag':'" + value.Message + "'},"
		}
		rels.Data += "]"
		rels.Status = 0
	} else {
		/*
			u := &models.User{}
			err := models.One("User", bson.M{"Name", name}, u, nil)
			if err != nil {
				rels.Status = 0
				rels.Data = "[{'key':'name','mag':'账户名已存在'}]"
			}else {
		*/
		m := md5.New()
		m.Write([]byte(password))
		password = hex.EncodeToString(m.Sum(nil))

		err := models.Insert("User", &bson.M{"Name": name, "Passwd": password})

		if err == nil {
			rels.Status = 1
		} else {
			rels.Status = 0
			rels.Data = "[{'key':'error','mag':'插入错误，请稍后再诗'}]"
		}
	}

	rel, _ := json.Marshal(rels)
	return c.RenderText(string(rel))
	return c.Redirect("/B_System/Index")

}
func (c User) Login() revel.Result {

	return c.Render()

}

func (c User) Loging(name, password string) revel.Result {
	c.Validation.Required(name).Message("name is wrong")
	c.Validation.Required(password).Message("password is wrong")

	rels := &models.AjaxRel{}

	if c.Validation.HasErrors() {
		errs := c.Validation.ErrorMap()
		rels.Data = "["
		for key, value := range errs {
			rels.Data = rels.Data + "{'key':'" + key + "','mag':'" + value.Message + "'},"
		}
		rels.Data += "]"
		rels.Status = 0

	} else {
		user := bson.M{}
		mgoc := models.DB.C("User")
		err := mgoc.Find(bson.M{"Name": name}).One(&user)

		if err != nil {
			rels.Status = 0
			rels.Data = "[{'key':'error','mag':'账户名和密码错误'}]"

		} else {

			m := md5.New()
			m.Write([]byte(password))
			password = hex.EncodeToString(m.Sum(nil))
			fmt.Println("要死额")

			if password == user["Passwd"].(string) {
				fmt.Println("如果用户名正确，则")
				//c.Session["Name"] = user["Name"].(string)
				rels.Status = 1
				var i = 19
				if i > 0 {
					return c.Redirect("/B_system/Index")
				}
			} else {
				fmt.Println("如果用户名错误，则")
				rels.Status = 0
				rels.Data = "[{'key':'error','mag':'插入错误'}]"
			}
		}
	}
	rel, _ := json.Marshal(rels)
	return c.RenderText(string(rel))

}
