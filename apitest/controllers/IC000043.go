package controllers

import (
	"github.com/astaxie/beego"
)

// Query Collection Contact List
type IC000043Controller struct {
	beego.Controller
}

// @Title
// @Description Query Collection Contact List
// @Param	body		body 	models.IC000043I	"Input parameter"
// @Success 200  {object} models.IC000043O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000043Controller) Post() ()  {
}
