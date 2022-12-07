package controllers

import (
	"github.com/astaxie/beego"
)
// Create micro loan product Simple
type PD000068Controller struct {
	beego.Controller
}

// @Title
// @Description  Create micro loan product Simple
// @Param	body		body 	models.PD000068I	"Input parameter"
// @Success 200  {object} models.PD000068O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000068Controller) Post() {
}
