package controllers

import (
	"github.com/astaxie/beego"
)

// Query Collection Contact Detail
type IC000044Controller struct {
	beego.Controller
}

// @Title
// @Description Query Collection Contact Detail
// @Param	body		body 	models.IC000044I	"Input parameter"
// @Success 200  {object} models.IC000044O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000044Controller) Post() ()  {
}
