package controllers

import (
	"github.com/astaxie/beego"
)
// create new base interest rate
type PR000001Controller struct {
	beego.Controller
}

// @Title
// @Description  create new base interest rate
// @Param	body		body 	models.SPR0000001I	"Input parameter"
// @Success 200  {object} models.SPR0000001O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PR000001Controller) Post() {
}
