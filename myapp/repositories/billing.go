package repositories

import (
	"errors"
	"myapp/database"
	"myapp/models"

	"gorm.io/gorm"
)

// BillsRepository is a struct to represent the bills repository.
type BillsRepository struct{}

// Createbills inserts a new bills into the database.
func (r *BillsRepository) CreateBills(bills *models.Billing) (*models.Billing, error) {
	if err := database.DB.Create(bills).Error; err != nil {
		return nil, err
	}
	return bills, nil
}

// GetbillsByID retrieves a bills by their ID from the database.
func (r *BillsRepository) GetBillsByID(id int) (*models.Billing, error) {
	var bills models.Billing
	result := database.DB.First(&bills, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("bills not found")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &bills, nil
}

// Updatebills updates an existing bills in the database.
func (r *BillsRepository) UpdateBills(bills *models.Billing) error {
	if err := database.DB.Save(bills).Error; err != nil {
		return err
	}
	return nil
}

// Deletebills removes a bills from the database.
func (r *BillsRepository) DeleteBills(id int) error {
	result := database.DB.Delete(&models.Billing{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no bills found to delete")
	}
	return nil
}

// Listbillss retrieves all billss from the database.
func (r *BillsRepository) ListBills() ([]models.Billing, error) {
	var bills []models.Billing
	if err := database.DB.Find(&bills).Error; err != nil {
		return nil, err
	}
	return bills, nil
}
