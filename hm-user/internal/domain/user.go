package domain

type User struct {
	ID    int64
	Name  string
	Email string
}

func FindUserByID(id int64) *User {
	if id == 1 {
		return &User{
			ID:    1,
			Name:  "Alice",
			Email: "alice@qq.com",
		}
	}

	return nil
}