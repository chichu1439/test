package controller

import "github.com/gin-gonic/gin"

// FP000018 godoc
// @Summary     settlement account detail inquiry
// @Description settlement account detail inquiry
// @Tags        paymentSwitch
// @Param       FP000018Request body     models.FP000018I true "settlement account detail inquiry"
// @Success     200                        {object} controller.Response{response=models.FP000018O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/settlement-account-detail/inquiry [post]
func (c *Controller) FP000018(ctx *gin.Context) {

}
