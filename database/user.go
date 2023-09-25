package database

import "github.com/itaypanda/fiber-auth-example/models"

func CreateUser(user models.User) error {
	tx := db.Create(&user)
	return tx.Error
}

func FindUser(id string) (*models.User, error) {
	user := new(models.User)
	tx := db.Where("id = ?", id).First(&user)
	return user, tx.Error
}

func FindUserByEmail(email string) (*models.User, error) {
	user := new(models.User)
	tx := db.Where("email = ?", email).First(&user)
	return user, tx.Error
}


