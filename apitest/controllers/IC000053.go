package controllers

import (
	"github.com/astaxie/beego"
)

// Report File List Query
type IC000053Controller struct {
	beego.Controller
}

// @Title
// @Description Report File List Query
// @Param	body		body 	models.IC000053I	"Input parameter"
// @Success 200  {object} models.IC000053O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000053Controller) Post() ()  {
}
