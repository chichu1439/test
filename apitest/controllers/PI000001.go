package controllers

import (
	"github.com/astaxie/beego"
)
// Create Base Interest Rate
type PI000001Controller struct {
	beego.Controller
}

// @Title
// @Description  Create Base Interest Rate
// @Param	body		body 	models.PI000001I	"Input parameter"
// @Success 200  {object} models.PI000001O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PI000001Controller) Post() {
}
