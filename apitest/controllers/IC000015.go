package controllers

import (
	"github.com/astaxie/beego"
)

// Loan risk classification inquiry
type IC000015Controller struct {
	beego.Controller
}

// @Title
// @Description Loan risk classification inquiry
// @Param	body		body 	models.IC000015I	"Input parameter"
// @Success 200  {object} models.IC000015O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000015Controller) Post() ()  {
}
