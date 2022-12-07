package controllers

import (
	"github.com/astaxie/beego"
)

// Query transaction history detail
type IC000034Controller struct {
	beego.Controller
}

// @Title
// @Description Query transaction history detail
// @Param	body		body 	models.IC000034I	"Input parameter"
// @Success 200  {object} models.IC000034O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000034Controller) Post() ()  {
}
