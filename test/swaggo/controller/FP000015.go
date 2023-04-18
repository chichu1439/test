package controller

import "github.com/gin-gonic/gin"

// FP000015 godoc
// @Summary     settlement account inquiry
// @Description settlement account inquiry
// @Tags        paymentSwitch
// @Param       FP000015Request body     models.FP000015I true "settlement account inquiry"
// @Success     200                        {object} controller.Response{response=models.FP000015O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/settlement-account/inquiry [post]
func (c *Controller) FP000015(ctx *gin.Context) {

}
