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
