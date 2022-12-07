package controllers

import (
	"github.com/astaxie/beego"
)

type PF000007Controller struct {
	beego.Controller
}

// @Title
// @Description  Query Fee  Strategy Detail
// @Param	body		body 	models.PF000007I	"Input parameter"
// @Success 200  {object} models.PF000007O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PF000007Controller) Post() {
}
