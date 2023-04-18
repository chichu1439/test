package controller

import "github.com/gin-gonic/gin"

// FP000110 godoc
// @Summary     checker enquiry
// @Description checker enquiry
// @Tags        paymentSwitch
// @Param       FP000110Request body     models.FP000110I true "checker enquiry"
// @Success     200                        {object} controller.Response{response=models.FP000110O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/balance/adjustment/checker/enquiry [post]
func (c *Controller) FP000110(ctx *gin.Context) {

}
