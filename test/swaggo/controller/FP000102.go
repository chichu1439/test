package controller

import "github.com/gin-gonic/gin"

// FP000102 godoc
// @Summary     sample query
// @Description sample query
// @Tags        paymentSwitch
// @Param       FP000102Request body     models.FP000102I true "sample query"
// @Success     200                        {object} controller.Response{response=models.FP000102O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/report/sample/query [post]
func (c *Controller) FP000102(ctx *gin.Context) {

}
