package db

// Insert or update user params matcher.

func (i InsertOrUpdateUserParams) Matches(x interface{}) bool {
	m, ok := x.(InsertOrUpdateUserParams)
	if !ok {
		return false
	}

	return i.ID == m.ID && i.Username == m.Username &&
		i.Email == m.Email && i.FullName == m.FullName &&
		i.PasswordHash == m.PasswordHash && i.PictureID == m.PictureID &&
		i.Role == m.Role && i.CreatedAt == m.CreatedAt &&
		!i.UpdatedAt.Equal(m.UpdatedAt)
}

func (i InsertOrUpdateUserParams) String() string {
	return "user update params match"
}
