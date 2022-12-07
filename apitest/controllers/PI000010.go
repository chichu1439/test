package controllers

import (
	"github.com/astaxie/beego"
)
// calculate interest by product
type PI000010Controller struct {
	beego.Controller
}

// @Title
// @Description  calculate interest by product
// @Param	body		body 	models.PI000010I	"Input parameter"
// @Success 200  {object} models.PI000010O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PI000010Controller) Post() {
}
