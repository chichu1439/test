package controllers

import (
	"github.com/astaxie/beego"
)
// Query Fee Strategy List
type PF000006Controller struct {
	beego.Controller
}

// @Title
// @Description  Query Fee Strategy List
// @Param	body		body 	models.PF000006I	"Input parameter"
// @Success 200  {object} models.PF000006O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PF000006Controller) Post() {
}
