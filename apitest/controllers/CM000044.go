package controllers

import (
	"github.com/astaxie/beego"
)

// Create Collection Task
type CM000044Controller struct {
	beego.Controller
}

// @Title
// @Description Create Collection Task
// @Param	body		body 	models.CM000044I	"Input parameter"
// @Success 200  {object} models.CM000044O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *CM000044Controller) Post() ()  {
}
