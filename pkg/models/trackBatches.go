package models

import (
	"time"

	"github.com/yashre-bh/kla-crm-btp/pkg/types"
)

func AddToMasterTracking(batchCode string, dateOfArrival *time.Time) error {
	database, err := Connect()
	if err != nil {
		return err
	}

	var masterTracking types.AddToTracking
	masterTracking.BatchCode = batchCode
	masterTracking.DateAdded = dateOfArrival

	err = database.Table("master_tracking").Create(&masterTracking).Error
	return err
}
