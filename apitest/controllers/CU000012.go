package controllers

import (
	"github.com/astaxie/beego"
)

// Search Customer List
type CU000012Controller struct {
	beego.Controller
}

// @Title
// @Description Search Customer List
// @Param	body		body 	models.SCU0000012I	"Input parameter"
// @Success 200  {object} models.SCU0000012O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *CU000012Controller) Post() ()  {
}
