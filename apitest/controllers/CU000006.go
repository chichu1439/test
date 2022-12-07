package controllers

import (
	"github.com/astaxie/beego"
)

// Query Customer Status
type CU000006Controller struct {
	beego.Controller
}

// @Title
// @Description Query Customer Status
// @Param	body		body 	models.SCU0000006I	"Input parameter"
// @Success 200  {object} models.SCU0000006O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *CU000006Controller) Post() ()  {
}
