package controllers

import (
	"github.com/astaxie/beego"
)

// Repayment application
type IC000001Controller struct {
	beego.Controller
}

// @Title
// @Description Repayment application
// @Param	body		body 	models.IC000001I	"Input parameter"
// @Success 200  {object} models.IC000001O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000001Controller) Post() ()  {
}
