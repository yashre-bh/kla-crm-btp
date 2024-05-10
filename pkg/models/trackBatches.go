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

func BatchProgressToCheckpoint2(batchCode string) error {
	database, err := Connect()
	if err != nil {
		return err
	}

	err = database.Table("master_tracking").Where("batch_code = ?", batchCode).Update("checkpoint_2_entered", true).Error
	return err
}

func BatchProgressToCheckpoint3(batchCode string) error {
	database, err := Connect()
	if err != nil {
		return err
	}

	err = database.Table("master_tracking").Where("batch_code = ?", batchCode).Update("checkpoint_3_entered", true).Error
	return err
}

func BatchProgressToCheckpoint4(batchCode string) error {
	database, err := Connect()
	if err != nil {
		return err
	}

	err = database.Table("master_tracking").Where("batch_code = ?", batchCode).Update("checkpoint_4_entered", true).Error
	return err
}

//make similar functions for chkpt 2,3,4

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

func FetchFormDataFromCheckpoint1(checkpointID int32, batchCode string) (*types.IncomingRawMaterialDBQuery, error) {
	var FormDataIncomingRawMaterial *types.IncomingRawMaterialDBQuery

	database, err := Connect()
	if err != nil {
		return nil, err
	}

	err = database.Table("incoming_raw_material").Where("batch_code = ?", batchCode).First(&FormDataIncomingRawMaterial).Error
	return FormDataIncomingRawMaterial, err
}

func FetchActiveBatches() (*[]types.MasterTracking, error) {
	database, err := Connect()
	if err != nil {
		return nil, err
	}

	var masterTracking *[]types.MasterTracking

	err = database.Table("master_tracking").Where("active_status = ?", true).Select("batch_code").Find(&masterTracking).Error

	return masterTracking, err
}

func AddSubBatchRecords(batchCode string, subBatchCode []string) error {
	database, err := Connect()
	if err != nil {
		return err
	}

	for _, subBatch := range subBatchCode {
		err = database.Table("sub_batch_records").Create(&types.SubBatches{BatchCode: batchCode, SubBatchCode: subBatch}).Error
		if err != nil {
			return err
		}
	}

	return nil
}
