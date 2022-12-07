package controllers

import (
	"github.com/astaxie/beego"
)

// Query Customer Product Credit Line Detail
type IC000061Controller struct {
	beego.Controller
}

// @Title
// @Description Query Customer Product Credit Line Detail
// @Param	body		body 	models.IC000061I	"Input parameter"
// @Success 200  {object} models.IC000061O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000061Controller) Post() ()  {
}
