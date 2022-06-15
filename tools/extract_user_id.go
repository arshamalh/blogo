package tools

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ExtractUserID extracts user ID from the request.
// returns error if there is no user ID or if the user ID is not a number.
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

func ExtractPermissable(ctx *gin.Context) bool {
	permissable, exists := ctx.Get("permissable")
	if !exists {
		return false
	}
	return permissable.(bool)
}
