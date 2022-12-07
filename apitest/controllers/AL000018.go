package controllers

import (
	"github.com/astaxie/beego"
)
// Cancel Write Off
type AL000018Controller struct {
	beego.Controller
}

// @Title
// @Description Cancel Write Off
// @Param	body		body 	models.AL000018I	"Input parameter"
// @Success 200  {object} models.AL000018O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *AL000018Controller) Post() ()  {
}
