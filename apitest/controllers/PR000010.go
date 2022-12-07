package controllers

import (
	"github.com/astaxie/beego"
)

//create fee business scenario
type PR000010Controller struct {
	beego.Controller
}

// @Title
// @Description  create fee business scenario
// @Param	body		body 	models.SPR0000010_CANCELI	"Input parameter"
// @Success 200  {object} models.SPR0000010_CANCELO  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PR000010Controller) Post() {
}
