package controllers

import (
	"github.com/astaxie/beego"
)

//deducted of deposit account transfer
type PR000013Controller struct {
	beego.Controller
}

// @Title
// @Description  deducted of deposit account transfer
// @Param	body		body 	models.SPR0000013I	"Input parameter"
// @Success 200  {object} models.SPR0000013O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PR000013Controller) Post() ()  {
}
