package controllers

import (
	"github.com/astaxie/beego"
)
// Query micro loan product
type PD000012Controller struct {
	beego.Controller
}

// @Title
// @Description Query micro loan product
// @Param	body		body 	models.SPD0000012I	"Input parameter"
// @Success 200  {object} models.SPD0000012O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000012Controller) Post() {
}
