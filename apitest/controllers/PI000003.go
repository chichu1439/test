package controllers

import (
	"github.com/astaxie/beego"
)
// modify base interest rate
type PI000003Controller struct {
	beego.Controller
}

// @Title
// @Description  modify base interest rate
// @Param	body		body 	models.PI000003I	"Input parameter"
// @Success 200  {object} models.PI000003O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PI000003Controller) Post() {
}
