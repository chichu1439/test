package controllers

import (
	"github.com/astaxie/beego"
)

// query interest strategy by product
type PR000005Controller struct {
	beego.Controller
}

// @Title
// @Description  query interest strategy by product
// @Param	body		body 	models.SPR0000005I	"Input parameter"
// @Success 200  {object} models.SPR0000005O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PR000005Controller) Post() {
}
