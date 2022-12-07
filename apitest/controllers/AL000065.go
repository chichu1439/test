package controllers

import (
	"github.com/astaxie/beego"
)

// 2c2p Loan repayment
type AL000065Controller struct {
	beego.Controller
}

// @Title
// @Description 2c2p Loan repayment
// @Param	body		body 	models.AL000065I	"Input parameter"
// @Success 200  {object} models.AL000065O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *AL000065Controller) Post() ()  {
}
