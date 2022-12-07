package controllers

import (
	"github.com/astaxie/beego"
)

// Query Template Detail
type NT000003Controller struct {
	beego.Controller
}

// @Title
// @Description Query Template Detail
// @Param	body		body 	models.NT000003I	"Input parameter"
// @Success 200  {object} models.NT000003O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *NT000003Controller) Post() ()  {
}
