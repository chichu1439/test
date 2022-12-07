package controllers

import (
	"github.com/astaxie/beego"
)

// Query Customer Product Credit Line List
type IC000060Controller struct {
	beego.Controller
}

// @Title
// @Description Query Customer Product Credit Line List
// @Param	body		body 	models.IC000060I	"Input parameter"
// @Success 200  {object} models.IC000060O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000060Controller) Post() ()  {
}
