package controllers

import (
	"github.com/astaxie/beego"
)

// Payment info encrypt for 2c2p
type AL000061Controller struct {
	beego.Controller
}

// @Title
// @Description Payment info encrypt for 2c2p
// @Param	body		body 	models.AL000061I	"Input parameter"
// @Success 200  {object} models.AL000061O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *AL000061Controller) Post() ()  {
}
