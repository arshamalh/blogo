package permissions

// Any user have one or some permissions.
//
// Any user has permission to delete or edit their own posts or categories, or approve comment on their own posts.
//
// Post permissions are not related to category permissions in anyway.
type Permission uint8

const (
	// Full access means this user can do anything with other users, their posts, and categories.
	FullAccess Permission = 0

	// Anyone who has this permission has the ability to invoke others permissions or give them higher permissions.
	// There are just a few limitations about deleting users and like that.
	AddAndChangeRoles Permission = 1

	// Full content access means this user can make or delete any post or category or comment, but can't do anything with other users.
	FullContents Permission = 2

	// Permission to make new posts and be an author.
	CreatePost Permission = 3

	// Users with this permission can edit others posts.
	EditPost Permission = 4

	// Users with this permission can delete others posts.
	DeletePost Permission = 5

	// Permission to create new categories.
	CreateCategory Permission = 6

	// Users with this permission can edit others categories.
	EditCategory Permission = 7

	// Users with this permission can delete others categories.
	DeleteCategory Permission = 8

	// Permission to approve or reject comments.
	ApproveComment Permission = 9
)
