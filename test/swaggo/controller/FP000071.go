package controller

import "github.com/gin-gonic/gin"

// FP000071 godoc
// @Summary     tcpSending
// @Description tcpSending
// @Tags        paymentSwitch
// @Param       FP000071Request body     models.FP000071I true "tcpSending"
// @Success     200                        {object} controller.Response{response=models.FP000071O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/simulator/tcp/sending [post]
func (c *Controller) FP000071(ctx *gin.Context) {

}
