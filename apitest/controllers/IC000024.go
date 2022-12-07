package controllers

import (
	"github.com/astaxie/beego"
)

// Loan transaltion records inquiry
type IC000024Controller struct {
	beego.Controller
}

// @Title
// @Description Loan transaltion records inquiry
// @Param	body		body 	models.IC000024I	"Input parameter"
// @Success 200  {object} models.IC000024O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000024Controller) Post() ()  {
}
