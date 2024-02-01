package repository

import (
	"BagasA11/GSC-quizHealthEdu-BE/api/models"
	"BagasA11/GSC-quizHealthEdu-BE/configs"
	"time"

	"gorm.io/gorm"
)

// "time"
// "BagasA11/GSC-quizHealthEdu-BE/api/models"
type UserRepository struct {
	Db *gorm.DB
}

/*create User Repository instance*/
func NewUserRepository() *UserRepository {
	return &UserRepository{
		Db: configs.GetDB(),
	}
}

func (ur *UserRepository) CreateUser(user models.User) error {
	tx := ur.Db.Begin()
	err := tx.Create(&user).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

func (ur *UserRepository) FindId(id uint) (models.User, error) {
	var user models.User
	//SELECT * FROM users WHERE id = {id} AND admin = false
	err := ur.Db.Where("id = ?", id).First(&user).Error
	return user, err
}

// user section
func (ur *UserRepository) FindUserId(id uint) (models.User, error) {
	var user models.User
	//SELECT * FROM users WHERE id = {id} AND admin = false
	err := ur.Db.Where("id = ? AND admin = ?", id, false).First(&user).Error
	return user, err
}

func (ur *UserRepository) AllUser() ([]models.User, error) {
	var users []models.User
	//SELECT * FROM users WHERE admin = false AND block = false
	err := ur.Db.Where("admin = ? AND block = ?", false, false).Find(&users).Error
	return users, err
}

func (ur *UserRepository) FindUsername(uName string) (models.User, error) {
	var user models.User
	//SELECT * FROM users WHERE username = {uName} AND admin = false
	err := ur.Db.Where("username = ? AND admin = ?", uName, false).First(&user).Error
	return user, err
}

func (ur *UserRepository) FindEmail(email string) (models.User, error) {
	var user models.User
	//SELECT * FROM users WHERE email = {email} AND admin = false
	err := ur.Db.Where("email = ? AND admin = ?", email, false).First(&user).Error
	return user, err
}

// admin section
func (ur *UserRepository) AllAdmin() ([]models.User, error) {
	var admins []models.User
	//SELECT * FROM users WHERE admin = false AND block = false
	err := ur.Db.Where("admin = ?", true).Find(&admins).Error
	return admins, err
}

func (ur *UserRepository) AdminId(id uint) (models.User, error) {
	var admin models.User
	//SELECT * FROM users where id = {id} AND admin = true
	err := ur.Db.Where("id = ? AND admin = ?", id, true).First(&admin).Error
	return admin, err
}

func (ur *UserRepository) AdminEmail(email string) (models.User, error) {
	var admin models.User
	//SELECT * FROM users WHERE email = {email} AND admin = true
	err := ur.Db.Where("email = ? AND admin = ?", email, true).First(&admin).Error
	return admin, err
}

func (ur *UserRepository) AdminUsername(uName string) (models.User, error) {
	var admin models.User
	//SELECT * FROM users WHERE username = {uName} AND admin = true
	err := ur.Db.Where("username = ? AND admin = ?", uName, true).First(&admin).Error
	return admin, err
}

func (ur *UserRepository) Updates(user models.User) error {
	tx := ur.Db.Begin()
	err := tx.Model(&models.User{}).Where("id = ?", user.ID).Updates(&user).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (ur *UserRepository) SetBio(id uint, bio string) error {
	tx := ur.Db.Begin()
	err := tx.Model(&models.User{}).Where("id = ? AND admin = ?", id, false).Update("bio", bio).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (ur *UserRepository) BlockUser(id uint) error {
	tx := ur.Db.Begin()
	//SELECT * FROM users WHERE id = {id} AND admin = false
	err := tx.Model(&models.User{}).Where("id = ? AND admin = ?", id, false).Update("block", true).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (ur *UserRepository) Delete(id uint) error {
	tx := ur.Db.Begin()
	err := tx.Model(&models.User{}).Where("id = ?", id).Update("deleted_at", time.Now().Unix()).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
