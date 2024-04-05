package models

import (
	"fmt"
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

func FetchDataForUncheckedFormsCheckpoint1() (*[]types.PendingCheckItems, error) {
	database, err := Connect()
	if err != nil {
		return nil, err
	}

	var pendingCheckItems *[]types.PendingCheckItems

	err = database.Select("incoming_raw_material.date_of_arrival, incoming_raw_material.added_by_employee, incoming_raw_material.batch_code", "incoming_raw_material.name").Table("incoming_raw_material").Joins("JOIN master_tracking ON incoming_raw_material.batch_code = master_tracking.batch_code").Where("master_tracking.checkpoint_1_checked = ?", false).Find(&pendingCheckItems).Error
	fmt.Println(pendingCheckItems)

	return pendingCheckItems, err
}
