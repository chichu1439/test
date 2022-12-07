package controllers

import (
	"github.com/astaxie/beego"
)
// Apply for Write Off
type AL000012Controller struct {
	beego.Controller
}

// @Title
// @Description Apply for Write Off
// @Param	body		body 	models.AL000012I	"Input parameter"
// @Success 200  {object} models.AL000012O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *AL000012Controller) Post() ()  {
}
