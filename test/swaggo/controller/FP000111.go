package controller

import "github.com/gin-gonic/gin"

// FP000111 godoc
// @Summary     result enquiry
// @Description result enquiry
// @Tags        paymentSwitch
// @Param       FP000111Request body     models.FP000111I true "result enquiry"
// @Success     200                        {object} controller.Response{response=models.FP000111O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/netting/result/enquiry [post]
func (c *Controller) FP000111(ctx *gin.Context) {

}
