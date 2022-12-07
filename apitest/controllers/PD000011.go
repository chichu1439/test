package controllers

import (
	"github.com/astaxie/beego"
)
// Create micro loan product
type PD000011Controller struct {
	beego.Controller
}

// @Title
// @Description Create micro loan product
// @Param	body		body 	models.SPD0000011I	"Input parameter"
// @Success 200  {object} models.SPD0000011O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000011Controller) Post() {
}
