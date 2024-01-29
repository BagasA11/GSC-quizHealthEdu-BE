package repository

import (
	"BagasA11/go-POS/api/models"
	"BagasA11/go-POS/configs"
	"time"

	"gorm.io/gorm"
)

type CashierRepository struct {
	Db *gorm.DB
}

/*create Cashier Repository instance*/
func NewCashierRepository() *CashierRepository {
	return &CashierRepository{
		Db: configs.GetDB(),
	}
}

/*create new record*/
func (CR *CashierRepository) Create(csr models.Cashier) error {
	tx := CR.Db.Begin()
	err := tx.Create(&csr).Error
	if err != nil {
		return err
	}
	tx.Commit()
	return err
}

/*retrieve cashier object if id given value match to Cashier ID */
func (CR *CashierRepository) FindByID(id uint) (models.Cashier, error) {
	var csr models.Cashier
	err := CR.Db.Where("id = ?", id).First(&csr).Error
	return csr, err
}

/*find Cashier Object by username attribute*/
func (CR *CashierRepository) FindByUsername(username string) (models.Cashier, error) {
	var csr models.Cashier
	err := CR.Db.Where("username = ?", username).First(&csr).Error
	return csr, err
}

/*retrieve entire record of Cashier object*/
func (CR *CashierRepository) All() ([]models.Cashier, error) {
	var csr []models.Cashier
	err := CR.Db.Find(&csr).Error
	return csr, err
}

/*change Cashier attribute*/
func (CR *CashierRepository) Update(csr *models.Cashier) error {
	tx := CR.Db.Begin()
	err := tx.Model(&models.Cashier{}).Where("id = ?", csr.ID).Updates(&csr).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

/*delete Retail object*/
func (CR *CashierRepository) Delete(ID uint) error {
	tx := CR.Db.Begin()
	err := tx.Model(&models.Cashier{}).Where("id = ?", ID).Update("deleted_at", time.Now().Unix()).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
