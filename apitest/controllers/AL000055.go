package controllers

import (
	"github.com/astaxie/beego"
)

// Query Write Off Status for Validation
type AL000055Controller struct {
	beego.Controller
}

// @Title
// @Description Query Write Off Status for Validation
// @Param	body		body 	models.AL000055I	"Input parameter"
// @Success 200  {object} models.AL000055O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *AL000055Controller) Post() ()  {
}
