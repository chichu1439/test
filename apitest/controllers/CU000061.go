package controllers

import (
	"github.com/astaxie/beego"
)

// new customer
type CU000061Controller struct {
	beego.Controller
}

// @Title
// @Description  new customer
// @Param	body		body 	models.CU000061Request	"Input parameter"
// @Success 200  {object} models.CU000061Response  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *CU000061Controller) Post() ()  {
}
