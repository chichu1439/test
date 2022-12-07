package controllers

import (
	"github.com/astaxie/beego"
)
// Modify Fee Item
type PF000003Controller struct {
	beego.Controller
}

// @Title
// @Description  Modify Fee Item
// @Param	body		body 	models.PF000003I	"Input parameter"
// @Success 200  {object} models.PF000003O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PF000003Controller) Post() {
}
