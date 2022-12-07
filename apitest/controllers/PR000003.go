package controllers

import (
	"github.com/astaxie/beego"
)
// calculate interest rate
type PR000003Controller struct {
	beego.Controller
}

// @Title
// @Description  calculate interest rate
// @Param	body		body 	models.SPR0000003I	"Input parameter"
// @Success 200  {object} models.SPR0000003O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PR000003Controller) Post() {
}
