package controllers

import (
	"github.com/astaxie/beego"
)

// [pdf][Email]Customer statement
type IC000012Controller struct {
	beego.Controller
}

// @Title
// @Description [pdf][Email]Customer statement
// @Param	body		body 	models.IC000012I	"Input parameter"
// @Success 200  {object} models.IC000012O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000012Controller) Post() ()  {
}
