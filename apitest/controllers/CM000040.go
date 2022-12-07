package controllers

import (
	"github.com/astaxie/beego"
)

// Blocking random code generation
type CM000040Controller struct {
	beego.Controller
}

// @Title
// @Description Blocking random code generation
// @Param	body		body 	models.CM000040I	"Input parameter"
// @Success 200  {object} models.CM000040O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *CM000040Controller) Post() ()  {
}
