package controller

import "github.com/gin-gonic/gin"

// FP000042 godoc
// @Summary     prod fee delete
// @Description prod fee delete
// @Tags        paymentSwitch
// @Param       FP000042Request body     models.FP000042I true "prod fee delete"
// @Success     200                        {object} controller.Response{response=models.FP000042O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/prod/fee/delete [post]
func (c *Controller) FP000042(ctx *gin.Context) {

}
