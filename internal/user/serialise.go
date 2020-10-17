package user

func (u *User) Serialize() map[string]interface{} {
	return map[string]interface{}{
		"id":       u.ID,
		"username": u.Username,
		"email":    u.Email,
	}
}

func ReadJSON(m map[string]interface{}) (*User, error) {
	// @TODO: check values
	u := NewUser()
	u.ID = uint(m["id"].(float64))
	u.Username = m["username"].(string)
	u.Email = m["email"].(string)

	// err in case of invalid json
	return u, nil
}
