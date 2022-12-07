package controllers

import (
	"github.com/astaxie/beego"
)
// Query category product
type PD000005Controller struct {
	beego.Controller
}

// @Title
// @Description Query category product
// @Param	body		body 	models.PD000005I	"Input parameter"
// @Success 200  {object} models.PD000005O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000005Controller) Post() {
}
