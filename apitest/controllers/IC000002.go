package controllers

import (
	"github.com/astaxie/beego"
)

// Repayment approve
type IC000002Controller struct {
	beego.Controller
}

// @Title
// @Description Repayment approve
// @Param	body		body 	models.IC000002I	"Input parameter"
// @Success 200  {object} models.IC000002O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000002Controller) Post() ()  {
}
