package permissions

import (
	"strconv"
	"strings"
)

// Any user has permission to delete or edit their own posts or categories, or approve comments on their own posts.
// Post permissions are not related to category permissions in any way.
type Permission uint8

func (perm Permission) String() string {
	return strconv.FormatUint(uint64(perm), 10)
}

const (
	// Full access means this user can do anything with other users, their posts, and categories.
	FullAccess Permission = 0

	CreateRole Permission = 1
	UpdateRole Permission = 2
	DeleteRole Permission = 3

	// Permission 4 to 20 are reserved for future use.

	// Full content access means this user can make or delete any post or category or comment, but can't do anything with other users.
	FullContents Permission = 20

	CreatePost Permission = 21
	// Users with this permission can edit others posts.
	EditPost Permission = 22
	// Users with this permission can delete others posts.
	DeletePost     Permission = 23
	CreateCategory Permission = 24
	EditCategory   Permission = 25
	DeleteCategory Permission = 26
	ApproveComment Permission = 27
)

func Compress(permissions []Permission) string {
	var compressed string
	for _, permission := range permissions {
		compressed += permission.String() + ":"
	}
	return compressed[:len(compressed)-1]
}

func Decompress(compressed string) []Permission {
	var perms []Permission
	for _, perm := range strings.Split(compressed, ":") {
		if perm != "" {
			uintPerm, _ := strconv.ParseUint(perm, 10, 8)
			perms = append(perms, Permission(uintPerm))
		}
	}
	return perms
}
