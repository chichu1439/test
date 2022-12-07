package controllers

import (
	"github.com/astaxie/beego"
)

// Query Template List
type NT000002Controller struct {
	beego.Controller
}

// @Title
// @Description Query Template List
// @Param	body		body 	models.NT000002I	"Input parameter"
// @Success 200  {object} models.NT000002O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *NT000002Controller) Post() ()  {
}
