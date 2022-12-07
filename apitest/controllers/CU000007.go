package controllers

import (
	"github.com/astaxie/beego"
)

// Query Customer Contract Count
type CU000007Controller struct {
	beego.Controller
}

// @Title
// @Description Query Customer Contract Count
// @Param	body		body 	models.SCU0000007I	"Input parameter"
// @Success 200  {object} models.SCU0000007O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *CU000007Controller) Post() ()  {
}
