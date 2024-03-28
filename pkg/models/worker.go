package models

import (
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
