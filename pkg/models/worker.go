package models

import (
	"strings"

	"github.com/yashre-bh/kla-crm-btp/pkg/types"
)

func AddIncomingRawMaterial(incomingRawMaterial *types.IncomingRawMaterial) error {
	database, err := Connect()
	if err != nil {
		return err
	}
	err = database.Table("incoming_raw_material").Create(&incomingRawMaterial).Error
	return err
}

func GetEntityCode(entity string) (string, error) {
	var entityCode types.RawMaterialCode
	database, err := Connect()
	if err != nil {
		return "", err
	}
	err = database.Table("raw_material_code").Where("LOWER(entity) = ?", strings.ToLower(entity)).First(&entityCode).Error
	return entityCode.EntityCode, err
}

func AddToActiveBatches(batchCode string, date string, entity string) error {
	database, err := Connect()
	if err != nil {
		return err
	}

	entityCode, err := GetEntityCode(entity)
	err = database.Table("batches").Create(&types.Batch{
		BatchCode:  batchCode,
		Date:       date,
		Dispatched: false,
		Entity:     entityCode,
	}).Error
	return err
}
