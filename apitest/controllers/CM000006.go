package controllers

import (
	"github.com/astaxie/beego"
)

//  file upload
type CM000006Controller struct {
	beego.Controller
}

// @Title
// @Description  file upload
// @Param	body		body 	models.CM000006I	"Input parameter"
// @Success 200  {object} models.CM000006O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *CM000006Controller) Post() ()  {
}
