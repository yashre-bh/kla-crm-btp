package models

import (
	"github.com/yashre-bh/kla-crm-btp/pkg/types"
)

func AddCheckpoint(checkpoint *types.Checkpoint) error {
	database, err := Connect()
	if err != nil {
		return err
	}
	err = database.Create(&checkpoint).Error
	return err
}

func FetchAllCheckpoints() ([]types.Checkpoint, error) {
	database, err := Connect()
	if err != nil {
		return nil, err
	}

	var checkpoints []types.Checkpoint
	err = database.Find(&checkpoints).Error
	if err != nil {
		return nil, err
	}

	return checkpoints, err
}

func FetchCheckpointByID(checkpointID int) (types.Checkpoint, error) {
	var checkpoint types.Checkpoint
	database, err := Connect()
	if err != nil {
		return checkpoint, err
	}

	err = database.Where("checkpoint_id = ?", checkpointID).First(&checkpoint).Error

	return checkpoint, err
}

func DeleteCheckpoint(checkpointID int) error {
	database, err := Connect()
	if err != nil {
		return err
	}

	return database.Table("checkpoints").Where("checkpoint_id = ?", checkpointID).Delete(&types.Checkpoint{}).Error
}
