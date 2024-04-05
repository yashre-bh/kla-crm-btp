package models

import (
	"strings"

	"github.com/yashre-bh/kla-crm-btp/pkg/types"
)

func AddIncomingRawMaterial(incomingRawMaterial *types.IncomingRawMaterialDBQuery) error {
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
