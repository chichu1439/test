package controllers

import (
	"github.com/astaxie/beego"
)
// Query Fee Item
type PF000002Controller struct {
	beego.Controller
}

// @Title
// @Description  Query Fee Item
// @Param	body		body 	models.PF000002I	"Input parameter"
// @Success 200  {object} models.PF000002O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PF000002Controller) Post() {
}
