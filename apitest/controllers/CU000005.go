package controllers

import (
	"github.com/astaxie/beego"
)

// Create Customer Contract Info
type CU000005Controller struct {
	beego.Controller
}

// @Title
// @Description Create Customer Contract Info
// @Param	body		body 	models.SCU0000005I	"Input parameter"
// @Success 200  {object} models.SCU0000005O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *CU000005Controller) Post() ()  {
}
