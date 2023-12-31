package onboarding

import (
	e "loanApp/pkg/error"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (o *onbaordObj) SendOtp(c *gin.Context) {
	var response ApiResponse
	var mobile map[string]interface{}
	if err := c.ShouldBindJSON(&mobile); err != nil {
		response.Errors = append(response.Errors, e.ErrorInfo[e.BadRequest].GetErrorDetails(""))
		response.Success = false
		c.JSON(http.StatusBadRequest, response)
	}
	success, err := o.db.SendOtp(c, mobile["id"].(int))
	if err != nil {
		response.Errors = append(response.Errors, e.ErrorInfo[e.BadRequest].GetErrorDetails(""))
		response.Success = false
		c.JSON(http.StatusBadRequest, response)
	}
	response.Data = success
	response.Success = true
	c.JSON(http.StatusAccepted, response)
}
func (o *onbaordObj) VerifyOtp(c *gin.Context) {
	var response ApiResponse
	var mpin map[string]interface{}
	if err := c.ShouldBindJSON(&mpin); err != nil {
		response.Errors = append(response.Errors, e.ErrorInfo[e.BadRequest].GetErrorDetails(""))
		c.JSON(http.StatusBadRequest, response)
	}
	_, err := o.db.VerifyOtp(c, mpin["id"].(int))
	if err != nil {
		response.Errors = append(response.Errors, e.ErrorInfo[e.BadRequest].GetErrorDetails(""))
		response.Success = false
		c.JSON(http.StatusBadRequest, response)
	}
	response.Data = ""
	response.Success = true
	c.JSON(http.StatusAccepted, response)
}

func (o *onbaordObj) SetMpin(c *gin.Context) {
	var response ApiResponse
	var mpin map[string]interface{}
	if err := c.ShouldBindJSON(&mpin); err != nil {
		response.Errors = append(response.Errors, e.ErrorInfo[e.BadRequest].GetErrorDetails(""))
		c.JSON(http.StatusBadRequest, response)
	}
	success, err := o.db.MpinUpdate(c, mpin["id"].(int))
	if err != nil {
		response.Errors = append(response.Errors, e.ErrorInfo[e.BadRequest].GetErrorDetails(""))
		response.Success = false
		c.JSON(http.StatusBadRequest, response)
	}
	response.Data = success.Id
	response.Success = true
	c.JSON(http.StatusAccepted, response)
}

func (o *onbaordObj) VerifyMpin(c *gin.Context) {
	var response ApiResponse
	var mpin GetMpinParam
	if err := c.ShouldBindJSON(&mpin); err != nil {
		response.Errors = append(response.Errors, e.ErrorInfo[e.BadRequest].GetErrorDetails(""))
		c.JSON(http.StatusBadRequest, response)
	}
	success, err := o.db.MpinCheck(c, mpin.Mpin)
	if err != nil {
		response.Errors = append(response.Errors, e.ErrorInfo[e.BadRequest].GetErrorDetails(""))
		c.JSON(http.StatusBadRequest, response)
	}
	response.Success = success
	c.JSON(http.StatusAccepted, response)
}
