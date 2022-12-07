package controllers

import (
	"github.com/astaxie/beego"
)

// Calculation of repayment amount
type AL000003Controller struct {
	beego.Controller
}

// @Title
// @Description  Calculation of repayment amount
// @Param	body		body 	models.AL000003I	"Input parameter"
// @Success 200  {object} models.AL000003O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *AL000003Controller) Post() ()  {
}
