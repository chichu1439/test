package controllers

import (
	"github.com/astaxie/beego"
)

// Query Customer Detail
type CU000003Controller struct {
	beego.Controller
}

// @Title
// @Description Query Customer Detail
// @Param	body		body 	models.SCU0000003I	"Input parameter"
// @Success 200  {object} models.SCU0000003O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *CU000003Controller) Post() ()  {
}
