package service

import (
	"fmt"
	"time"

	"github.co/golang-programming/restaurant/api/internal/entity"
	"github.co/golang-programming/restaurant/api/internal/modules/customer/dto"
	"github.co/golang-programming/restaurant/api/internal/modules/customer/repository"
	utils "github.co/golang-programming/restaurant/api/pkg/utils/encryption"
)

func CreateCustomer(input *dto.CreateCustomerInput) (string, error) {
	customer := &entity.Customer{
		TotalGuests: input.TotalGuests,
		Name:        input.Name,
		TableID:     input.TableID,
		StartTime:   time.Now(),
	}
	if err := repository.CreateCustomer(customer); err != nil {
		return "", err
	}

	return utils.Encryptor(fmt.Sprint(customer.ID))
}

func GetCustomerByID(id uint) (*entity.Customer, error) {
	return repository.GetCustomerByID(id)
}

func UpdateCustomer(id uint, input *dto.UpdateCustomerInput) (*entity.Customer, error) {
	customer, err := repository.GetCustomerByID(id)
	if err != nil {
		return nil, err
	}

	customer.TotalGuests = input.TotalGuests
	customer.Name = input.Name
	if err := repository.UpdateCustomer(customer); err != nil {
		return nil, err
	}

	return customer, nil
}

func DeleteCustomer(id uint) error {
	customer, err := repository.GetCustomerByID(id)
	if err != nil {
		return err
	}
	return repository.DeleteCustomer(customer)
}

func ListCustomers() ([]*entity.Customer, error) {
	return repository.ListCustomers()
}

func DeactivateCustomer(customerID uint) error {
	customer, err := repository.GetCustomerByID(customerID)
	if err != nil {
		return err
	}

	customer.Status = entity.CustomerActive
	if err := repository.UpdateCustomer(customer); err != nil {
		return err
	}

	return nil
}
