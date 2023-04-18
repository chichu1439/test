package controller

import "github.com/gin-gonic/gin"

// FP010009 godoc
// @Summary     message redeliver
// @Description message redeliver
// @Tags        paymentSwitch
// @Param       FP010009Request body     models.FP010009I true "message redeliver"
// @Success     200                        {object} controller.Response{response=models.FP010009O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/message/redeliver [post]
func (c *Controller) FP010009(ctx *gin.Context) {

}
