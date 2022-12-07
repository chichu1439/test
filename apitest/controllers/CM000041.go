package controllers

import (
	"github.com/astaxie/beego"
)

// Blocking random code deletion
type CM000041Controller struct {
	beego.Controller
}

// @Title
// @Description Blocking random code deletion
// @Param	body		body 	models.CM000041I	"Input parameter"
// @Success 200  {object} models.CM000041O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *CM000041Controller) Post() ()  {
}
