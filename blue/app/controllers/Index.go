package controllers

import (
	"WEB/WebCore/app/models"
	//"fmt"
	"github.com/revel/revel"
	"labix.org/v2/mgo/bson"
)

type Index struct {
	*revel.Controller
}

func (c Index) Index() revel.Result {
	f := make([]bson.M, 0)
	models.List("webflash", nil, &f, nil)
	c.RenderArgs["flash"] = f

	list := make([]bson.M, 0)
	models.List("doc", bson.M{"Tid": "work"}, &list, nil)
	c.RenderArgs["list"] = list

	life := make([]bson.M, 0)
	models.List("doc", bson.M{"Tid": "life"}, &life, nil)
	if len(life) > 0 {
		c.RenderArgs["firstlife"] = life[0]
		//fmt.Println(life[0])
		c.RenderArgs["life"] = life[1:]
	} else {
		c.RenderArgs["firststory"] = map[string]string{}
	}

	story := make([]bson.M, 0)
	models.List("doc", bson.M{"Tid": "story"}, &story, nil)
	if len(story) > 0 {
		c.RenderArgs["firststory"] = story[0]
		//fmt.Println(story[0])
		c.RenderArgs["story"] = story[1:]
	} else {
		//	fmt.Println("meiyou")
		c.RenderArgs["firststory"] = map[string]string{}
	}

	return c.RenderTemplate("App" + "/Index.html")
}
