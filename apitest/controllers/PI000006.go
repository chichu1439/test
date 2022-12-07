package controllers

import (
	"github.com/astaxie/beego"
)
// query interest strategy list
type PI000006Controller struct {
	beego.Controller
}

// @Title
// @Description  query interest strategy list
// @Param	body		body 	models.PI000006I	"Input parameter"
// @Success 200  {object} models.PI000006O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PI000006Controller) Post() {
}
