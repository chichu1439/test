package controller

import "github.com/gin-gonic/gin"

// FP000027 godoc
// @Summary     prod fee insert
// @Description prod fee insert
// @Tags        paymentSwitch
// @Param       FP000027Request body     models.FP000027I true "prod fee insert"
// @Success     200                        {object} controller.Response{response=models.FP000027O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/prod/fee/insert [post]
func (c *Controller) FP000027(ctx *gin.Context) {

}
