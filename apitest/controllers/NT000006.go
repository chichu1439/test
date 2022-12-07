package controllers

import (
	"github.com/astaxie/beego"
)

// Create Notification Strategy
type NT000006Controller struct {
	beego.Controller
}

// @Title
// @Description Create Notification Strategy
// @Param	body		body 	models.NT000006I	"Input parameter"
// @Success 200  {object} models.NT000006O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *NT000006Controller) Post() ()  {
}
