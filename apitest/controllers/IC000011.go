package controllers

import (
	"github.com/astaxie/beego"
)

// approve modify customer
type IC000011Controller struct {
	beego.Controller
}

// @Title
// @Description approve modify customer
// @Param	body		body 	models.IC000011I	"Input parameter"
// @Success 200  {object} models.IC000011O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000011Controller) Post() ()  {
}
