package controllers

import (
	"github.com/astaxie/beego"
)
// modify fee strategy
type PF000008Controller struct {
	beego.Controller
}

// @Title
// @Description  modify fee strategy
// @Param	body		body 	models.PF000008I	"Input parameter"
// @Success 200  {object} models.PF000008O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PF000008Controller) Post() {
}
