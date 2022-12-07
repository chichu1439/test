package controllers

import (
	"github.com/astaxie/beego"
)

// receive new customer flow
type IC000009Controller struct {
	beego.Controller
}

// @Title
// @Description receive new customer flow
// @Param	body		body 	models.IC000009I	"Input parameter"
// @Success 200  {object} models.IC000009O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000009Controller) Post() ()  {
}
