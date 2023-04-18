package controller

import "github.com/gin-gonic/gin"

// FP000008 godoc
// @Summary     participant agency relation detail inquiry
// @Description participant agency relation detail inquiry
// @Tags        paymentSwitch
// @Param       FP000008Request body     models.FP000008I true "participant agency relation detail inquiry"
// @Success     200                        {object} controller.Response{response=models.FP000008O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/participant-agency-relation-detail/inquiry [post]
func (c *Controller) FP000008(ctx *gin.Context) {

}
