package controllers

import (
	"WEB/WebCore/app/models"

	"github.com/revel/revel"
	"labix.org/v2/mgo/bson"
)

type Nor struct {
	*revel.Controller
}

func (c Nor) List(tid string) revel.Result {
	tree := bson.M{}
	models.One("webtree", bson.M{"_id": tid}, &tree, nil)
	c.RenderArgs["tree"] = tree

	list := make([]bson.M, 0)
	models.List("doc", bson.M{"Tid": tid}, &list, nil)
	c.RenderArgs["list"] = list

	f := make([]bson.M, 0)
	models.List("weblink", nil, &f, nil)
	c.RenderArgs["f"] = f

	return c.RenderTemplate("App" + "/" + tree["Tmp"].(string) + ".html")
}

func (c Nor) Cout(tid, id, name string) revel.Result {
	tree := bson.M{}
	models.One("webtree", bson.M{"_id": tid}, &tree, nil)
	cout := bson.M{}
	models.One("doc", bson.M{"_id": id}, &cout, nil)

	c.RenderArgs["tree"] = tree
	c.RenderArgs["cout"] = cout

	last := make([]bson.M, 0)
	models.List("doc", bson.M{"Name": name}, &last, nil)
	c.RenderArgs["last"] = last

	return c.RenderTemplate("App" + "/" + tree["Tmp"].(string) + "_cout.html")
}
