package controller

import "github.com/gin-gonic/gin"

// FP020010 godoc
// @Summary     execl download
// @Description execl download
// @Tags        paymentSwitch
// @Param       FP020010Request body     models.DownRequstParam true "execl download"
// @Success     200                        {object} controller.Response{response=models.ResetSvcData}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/import/execl/download [post]
func (c *Controller) FP020010(ctx *gin.Context) {

}
