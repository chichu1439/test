package controllers

import (
	"github.com/astaxie/beego"
)
//  create contract interest relationship
type PR000004Controller struct {
	beego.Controller
}

// @Title
// @Description  create contract interest relationship
// @Param	body		body 	models.SPR0000004I	"Input parameter"
// @Success 200  {object} models.SPR0000004O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PR000004Controller) Post() {
}
