package controllers

import (
	"github.com/astaxie/beego"
)
// Delete Fee Item
type PF000004Controller struct {
	beego.Controller
}

// @Title
// @Description  Delete Fee Item
// @Param	body		body 	models.PF000004I	"Input parameter"
// @Success 200  {object} models.PF000004O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PF000004Controller) Post() {
}
