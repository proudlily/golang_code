package controllers

import (
	"github.com/revel/revel"
)

func onStart(c *revel.Controller) revel.Result {
	if c.Request.URL.Path == "/B_System/Index" {
		if c.Session["name"] == "" {
			return c.Redirect("User/Login?to=%s", c.Request.URL)
		}
	}
	return nil
}

func init() {
	revel.InterceptFunc(onStart, revel.BEFORE, revel.ALL_CONTROLLERS)
}
