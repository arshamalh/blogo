package tools

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ExtractUserID(ctx *gin.Context) (uint, error) {
	user_id_str, exists := ctx.Get("user_id")
	if !exists {
		return 0, fmt.Errorf("user ID not found")
	}
	user_id, err := strconv.ParseUint(user_id_str.(string), 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(user_id), nil
}
