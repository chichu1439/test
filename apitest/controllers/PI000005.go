package controllers

import (
	"github.com/astaxie/beego"
)
// create interest strategy
type PI000005Controller struct {
	beego.Controller
}

// @Title
// @Description  create interest strategy
// @Param	body		body 	models.PI000005I	"Input parameter"
// @Success 200  {object} models.PI000005O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PI000005Controller) Post() {
}
