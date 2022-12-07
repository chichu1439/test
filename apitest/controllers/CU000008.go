package controllers

import (
	"github.com/astaxie/beego"
)

// Query customer’s contract info
type CU000008Controller struct {
	beego.Controller
}

// @Title
// @Description Query customer’s contract info
// @Param	body		body 	models.SCU0000008I	"Input parameter"
// @Success 200  {object} models.SCU0000008O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *CU000008Controller) Post() ()  {
}
