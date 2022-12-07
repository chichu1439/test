package controllers

import (
	"github.com/astaxie/beego"
)
// Create Fee Strategy
type PF000005Controller struct {
	beego.Controller
}

// @Title
// @Description  Create Fee Strategy
// @Param	body		body 	models.PF000005I	"Input parameter"
// @Success 200  {object} models.PF000005O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PF000005Controller) Post() {
}
