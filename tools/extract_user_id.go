package tools

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

// ExtractUserID extracts user ID from the request.
// returns error if there is no user ID or if the user ID is not a number.
func ExtractUserID(ctx echo.Context) (uint, error) {
	user_id_str := ctx.Get("user_id")
	user_id, err := strconv.ParseUint(user_id_str.(string), 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(user_id), nil
}

func ExtractPermissable(ctx echo.Context) bool {
	permissable := ctx.Get("permissable")
	return permissable.(bool)
}
