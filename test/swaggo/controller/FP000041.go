package controller

import "github.com/gin-gonic/gin"

// FP000041 godoc
// @Summary     prod fee query
// @Description prod fee query
// @Tags        paymentSwitch
// @Param       FP000041Request body     models.FP000041I true "prod fee query"
// @Success     200                        {object} controller.Response{response=models.FP000041O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/prod/fee/query [post]
func (c *Controller) FP000041(ctx *gin.Context) {

}
