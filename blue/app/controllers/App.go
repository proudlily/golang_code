package controllers

import (
	"WEB/WebCore/app/models"

	"github.com/revel/revel"
	"labix.org/v2/mgo/bson"
)

type App struct {
	*revel.Controller
}

func (c App) Tag(tag, tid string) revel.Result {

	lkst := make([]bson.M, 0)
	models.List("doc", bson.M{"Tag": tag}, &lkst, nil)
	c.RenderArgs["Tag"] = tag

	last := make([]bson.M, 0)
	models.List("doc", bson.M{"Tid": tid}, &last, nil)
	c.RenderArgs["Tid"] = tid

	return c.Render(lkst, tag, last, tid)
}

func (c App) Search(name string) revel.Result {
	last := make([]bson.M, 0)
	models.List("doc", bson.M{"Name": bson.M{"$regex": name}}, &last, nil)
	c.RenderArgs["last"] = last

	return c.Render(last, name)
}

func (c App) Insert(name, email string) revel.Result {
	c.Validation.Required(name).Message("123")
	c.Validation.Email(email).Message("邮箱格式错误")
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
		data := bson.M{}
		data["Name"] = name
		data["Email"] = email

		data["_id"] = bson.NewObjectId()

		err := models.Insert("webfeedback", &data)

		if err == nil {
			rels.Status = 1

		} else {
			rels.Status = 0
			rels.Data = "[{'key':'error','mag':'Insert has some wrong :-)'}]"
		}
		return c.Redirect("/Index/Index")

	}
	return c.RenderJson(rels)

}
