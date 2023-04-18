package controller

import "github.com/gin-gonic/gin"

// FP000108 godoc
// @Summary     adjustment enquiry
// @Description adjustment enquiry
// @Tags        paymentSwitch
// @Param       FP000108Request body     models.FP000108I true "adjustment enquiry"
// @Success     200                        {object} controller.Response{response=models.FP000108O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/balance/adjustment/enquiry [post]
func (c *Controller) FP000108(ctx *gin.Context) {

}
