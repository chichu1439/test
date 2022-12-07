package controllers

import (
	"github.com/astaxie/beego"
)
// delete fee strategy
type PF000009Controller struct {
	beego.Controller
}

// @Title
// @Description  delete fee strategy
// @Param	body		body 	models.PF000009I	"Input parameter"
// @Success 200  {object} models.PF000009O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PF000009Controller) Post() {
}
