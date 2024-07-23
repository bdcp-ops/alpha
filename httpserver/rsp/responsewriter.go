package rsp

import (
	"github.com/bdcp-ops/alpha/aerror"
	"github.com/gin-gonic/gin"
)

func Error(c *gin.Context, err error) {
	var errResp *aerror.Error
	var ok bool

	if errResp, ok = err.(*aerror.Error); !ok {
		errResp = aerror.ErrInternalError().WithMessage(err.Error())
	}

	c.JSON(errResp.HTTPStatusCode, errResp)
}
