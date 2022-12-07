package controllers

import (
	"github.com/astaxie/beego"
)
// create new interest strategy
type PR000002Controller struct {
	beego.Controller
}

// @Title
// @Description  create new interest strategy
// @Param	body		body 	models.SPR0000002I	"Input parameter"
// @Success 200  {object} models.SPR0000002O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PR000002Controller) Post() {
}
