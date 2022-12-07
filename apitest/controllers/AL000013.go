package controllers

import (
	"github.com/astaxie/beego"
)
// Approve Write Off
type AL000013Controller struct {
	beego.Controller
}

// @Title
// @Description Approve Write Off
// @Param	body		body 	models.AL000013I	"Input parameter"
// @Success 200  {object} models.AL000013O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *AL000013Controller) Post() ()  {
}
