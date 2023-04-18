package controller

import "github.com/gin-gonic/gin"

// FP000072 godoc
// @Summary     tcpReceive
// @Description tcpReceive
// @Tags        paymentSwitch
// @Param       FP000072Request body     models.FP000072I true "tcpReceive"
// @Success     200                        {object} controller.Response{response=models.FP000072O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/simulator/tcp/receive [post]
func (c *Controller) FP000072(ctx *gin.Context) {

}
