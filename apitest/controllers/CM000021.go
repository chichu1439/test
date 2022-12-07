package controllers

import (
	"github.com/astaxie/beego"
)

// Query Option Info
type CM000021Controller struct {
	beego.Controller
}

// @Title
// @Description Query Option Info
// @Param	body		body 	models.SCM0000021I	"Input parameter"
// @Success 200  {object} models.SCM0000021O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *CM000021Controller) Post() ()  {
}
