package controller

import "github.com/gin-gonic/gin"

// FP000114 godoc
// @Summary     history enquiry
// @Description history enquiry
// @Tags        paymentSwitch
// @Param       FP000114Request body     models.FP000114I true "history enquiry"
// @Success     200                        {object} controller.Response{response=models.FP000114O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/balance/adjustment/history/enquiry [post]
func (c *Controller) FP000114(ctx *gin.Context) {

}
