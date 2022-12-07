package controllers

import (
	"github.com/astaxie/beego"
)
// Query Product Interest History
type PD000087Controller struct {
	beego.Controller
}

// @Title
// @Description  Query Product Interest History
// @Param	body		body 	models.PD000087I	"Input parameter"
// @Success 200  {object} models.PD000087O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000087Controller) Post() {
}
