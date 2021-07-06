package bd

import "github.com/brayanzv/FamiliarBudget2/models"

func InsertDetail(d models.Details)(bool,error){

	err := ConectionBD().Create(&d).Error
	ConectionBD().Close()
	if err != nil {
		return false, err
	}
	return true, nil
}
