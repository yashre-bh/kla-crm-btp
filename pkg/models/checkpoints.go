package models

import (
	"github.com/yashre-bh/kla-crm-btp/pkg/types"
)

func AddCheckpoint(checkpoint types.Checkpoint) error {
	db, err := Connection()
	if err != nil {
		return err
	}

	insertSql := "INSERT INTO checkpoints (checkpoint_id, checkpoint_name) VALUES (?, ?)"
	_, err = db.Exec(insertSql, checkpoint.CheckpointID, checkpoint.CheckpointName)
	db.Close()

	return err
}

func FetchAllCheckpoints() ([]types.Checkpoint, error) {
	db, err := Connection()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT checkpoint_id, checkpoint_name FROM checkpoints")
	db.Close()

	if err != nil {
		return nil, err
	}

	var checkpoints []types.Checkpoint
	for rows.Next() {
		var checkpoint types.Checkpoint
		err := rows.Scan(&checkpoint.CheckpointID, &checkpoint.CheckpointName)
		if err != nil {
			return nil, err
		}
		checkpoints = append(checkpoints, checkpoint)
	}

	return checkpoints, err
}
