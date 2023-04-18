package controller

import "github.com/gin-gonic/gin"

// FP000048 godoc
// @Summary     participant delete
// @Description participant delete
// @Tags        paymentSwitch
// @Param       FP000048Request body     models.FP000048I true "participant delete"
// @Success     200                        {object} controller.Response{response=models.FP000048O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/alert/participant/delete [post]
func (c *Controller) FP000048(ctx *gin.Context) {

}
