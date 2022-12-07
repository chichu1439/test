package controllers

import (
	"github.com/astaxie/beego"
)
// modify interest strategy
type PI000008Controller struct {
	beego.Controller
}

// @Title
// @Description  modify interest strategy
// @Param	body		body 	models.PI000008I	"Input parameter"
// @Success 200  {object} models.PI000008O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PI000008Controller) Post() {
}
