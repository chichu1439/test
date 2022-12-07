package controllers

import (
	"github.com/astaxie/beego"
)
// Modify Micro Loan Product Simple
type PD000070Controller struct {
	beego.Controller
}

// @Title
// @Description  Modify Micro Loan Product Simple
// @Param	body		body 	models.PD000070I	"Input parameter"
// @Success 200  {object} models.PD000070O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000070Controller) Post() {
}
