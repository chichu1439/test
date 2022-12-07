package controllers

import (
	"github.com/astaxie/beego"
)

//
type CM000052Controller struct {
	beego.Controller
}

// @Title
// @Description
// @Param	body		body 	models.CM000052I	"Input parameter"
// @Success 200  {object} models.CM000052O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *CM000052Controller) Post() ()  {
}
