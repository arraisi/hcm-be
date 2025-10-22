package user

// GetUsersRequest represents the request parameters for getting users
type GetUsersRequest struct {
	Limit  int
	Offset int
	Search string
	SortBy string
	Order  string
}

// CreateUserRequest represents the request for creating a user
type CreateUserRequest struct {
	Email string
	Name  string
}

// UpdateUserRequest represents the request for updating a user
type UpdateUserRequest struct {
	Email *string
	Name  *string
}
