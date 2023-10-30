package permissions

import (
	"log"
	"strconv"
	"strings"
)

// Any user have one or some permissions.
//
// Any user has permission to delete or edit their own posts or categories, or approve comment on their own posts.
//
// Post permissions are not related to category permissions in anyway.

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

	// Permission to make new posts and be an author.
	CreatePost Permission = 21

	// Users with this permission can edit others posts.
	EditPost Permission = 22

	// Users with this permission can delete others posts.
	DeletePost Permission = 23

	// Permission to create new categories.
	CreateCategory Permission = 24

	// Users with this permission can edit others categories.
	EditCategory Permission = 25

	// Users with this permission can delete others categories.
	DeleteCategory Permission = 26

	// Permission to approve or reject comments.
	ApproveComment Permission = 27
)

func Compress(permissions []Permission) string {
	var compressed string
	for _, permission := range permissions {
		compressed += permission.String() + ":"
	}
	log.Printf("Compressed permissions: %s", compressed) // Log the compressed permissions
	return compressed[:len(compressed)-1]
}

func Decompress(compressed string) []Permission {
	var perms []Permission
	for _, perm := range strings.Split(compressed, ":") {
		if perm != "" {
			uint_perm, _ := strconv.ParseUint(perm, 10, 8)
			permission := Permission(uint_perm)
			perms = append(perms, permission)
			log.Printf("Decompressed permission: %s", permission.String()) // Log the decompressed permissions
		}
	}
	return perms
}
