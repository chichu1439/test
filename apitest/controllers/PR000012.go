package controllers

import (
	"github.com/astaxie/beego"
)

//insert flow of deduction fee
type PR000012Controller struct {
	beego.Controller
}

// @Title
// @Description  insert flow of deduction fee
// @Param	body		body 	models.SPR0000012I	"Input parameter"
// @Success 200  {object} models.SPR0000012O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PR000012Controller) Post() {
}
