package controllers

import (
	"github.com/astaxie/beego"
)

// Loan balance Statistics inquiry
type IC000013Controller struct {
	beego.Controller
}

// @Title
// @Description Loan balance Statistics inquiry
// @Param	body		body 	models.IC000013I	"Input parameter"
// @Success 200  {object} models.IC000013O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000013Controller) Post() ()  {
}
