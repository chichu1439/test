package controllers

import (
	"github.com/astaxie/beego"
)

// Repayment plan query
type IC000027Controller struct {
	beego.Controller
}

// @Title
// @Description Repayment plan query
// @Param	body		body 	models.IC000027I	"Input parameter"
// @Success 200  {object} models.IC000027O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000027Controller) Post() ()  {
}
