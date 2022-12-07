package controllers

import (
	"github.com/astaxie/beego"
)

// Get trial calculation of repayment plan
type BP000002Controller struct {
	beego.Controller
}

// @Title
// @Description  Get trial calculation of repayment plan
// @Param	body		body 	models.BP000002Request	"Input parameter"
// @Success 200  {object} models.BP000002Response  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *BP000002Controller) Post() ()  {
}
