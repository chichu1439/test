package controller

import "github.com/gin-gonic/gin"

// FP000030 godoc
// @Summary     report file query
// @Description report file query
// @Tags        paymentSwitch
// @Param       FP000030Request body     models.FP000030I true "report file query"
// @Success     200                        {object} controller.Response{response=models.FP000030O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/report/file/query [post]
func (c *Controller) FP000030(ctx *gin.Context) {

}
