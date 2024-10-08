package repository

import (
	"github.co/golang-programming/restaurant/api/internal/entity"
	"github.co/golang-programming/restaurant/api/pkg/database"
)

func CreateBill(bill *entity.Bill) error {
	return database.ActiveDB.Create(bill).Error
}

func GetBillByID(id uint) (*entity.Bill, error) {
	var bill entity.Bill
	if err := database.ActiveDB.First(&bill, id).Error; err != nil {
		return nil, err
	}
	return &bill, nil
}

func GetBillByTableID(tableID uint) (*entity.Bill, error) {
	var bill entity.Bill
	if err := database.ActiveDB.Where("table_id = ?", tableID).First(&bill).Error; err != nil {
		return nil, err
	}
	return &bill, nil
}

func UpdateBill(bill *entity.Bill) error {
	return database.ActiveDB.Save(bill).Error
}

func DeleteBill(bill *entity.Bill) error {
	return database.ActiveDB.Delete(bill).Error
}

func ListBills() ([]*entity.Bill, error) {
	var bills []*entity.Bill
	if err := database.ActiveDB.Find(&bills).Error; err != nil {
		return nil, err
	}
	return bills, nil
}
