package regismodel

import (
	"btpn-ta/config"
	"btpn-ta/entities"
)

func Create(category entities.User) bool {
	result, err := config.DB.Exec(`
		INSERT INTO categories (username, email, password, created_at, updated_at) 
		VALUE (?, ?, ?)`,
	)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}
