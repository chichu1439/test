package controllers

import (
	"github.com/astaxie/beego"
)
// Create new category
type PD000001Controller struct {
	beego.Controller
}

// @Title
// @Description Create new category
// @Param	body		body 	models.PD000001I	"Input parameter"
// @Success 200  {object} models.PD000001O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000001Controller) Post() {
}
