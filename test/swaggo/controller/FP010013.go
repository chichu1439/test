package controller

import "github.com/gin-gonic/gin"

// FP010013 godoc
// @Summary     sweeping enquiry
// @Description sweeping enquiry
// @Tags        paymentSwitch
// @Param       FP010013Request body     models.FP010013I true "sweeping enquiry"
// @Success     200                        {object} controller.Response{response=models.FP010013O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/balance/sweeping/enquiry [post]
func (c *Controller) FP010013(ctx *gin.Context) {

}
