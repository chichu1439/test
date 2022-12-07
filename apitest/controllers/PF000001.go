package controllers

import (
	"github.com/astaxie/beego"
)
// Create Fee Item
type PF000001Controller struct {
	beego.Controller
}

// @Title
// @Description  Create Fee Item
// @Param	body		body 	models.PF000001I	"Input parameter"
// @Success 200  {object} models.PF000001O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PF000001Controller) Post() {
}
