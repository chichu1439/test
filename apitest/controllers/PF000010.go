package controllers

import (
	"github.com/astaxie/beego"
)
// calculate fee by product
type PF000010Controller struct {
	beego.Controller
}

// @Title
// @Description  calculate fee by product
// @Param	body		body 	models.PF000010I	"Input parameter"
// @Success 200  {object} models.PF000010O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PF000010Controller) Post() {
}
