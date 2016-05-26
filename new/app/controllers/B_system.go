package controllers

import (
	"github.com/revel/revel"
	"labix.org/v2/mgo/bson"
	"new/app/models"
	"time"
)

type B_system struct {
	*revel.Controller
}

func (c B_system) Index() revel.Result {
	data := []bson.M{}
	models.List("munix", nil, &data, nil)

	return c.Render(data)
}

func (c B_system) Write(name, tag, contain string) revel.Result {
	c.Validation.Required(name).Message("请输入名字")
	c.Validation.Required(tag).Message("请输入标签")
	c.Validation.Required(contain).Message("请输入内容")

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
		data["Tag"] = tag
		data["Contain"] = contain

		data["time"] = time.Now().Unix()
		data["_id"] = bson.NewObjectId()
		err := models.Insert("Atc", &data)
		if err == nil {
			rels.Status = 1
			var i = 19
			if i > 0 {
				return c.Redirect("/B_system/Index")
			}
		} else {
			rels.Status = 0
			rels.Data = "[{'key':'error','mag':'插入错误，请稍后再试'}]"
		}
	}

	return c.RenderJson(rels)

}
func (c B_system) List() revel.Result {
	data := []bson.M{}
	models.List("blog", nil, &data, nil)

	return c.Render(data)
}
