package controllers

import (
	"github.com/astaxie/beego"
)

// Query Write Off Info
type IC000055Controller struct {
	beego.Controller
}

// @Title
// @Description Query Write Off Info
// @Param	body		body 	models.IC000055I	"Input parameter"
// @Success 200  {object} models.IC000055O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000055Controller) Post() ()  {
}
