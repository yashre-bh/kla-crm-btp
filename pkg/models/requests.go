package models

import (
	"github.com/yashre-bh/kla-crm-btp/pkg/types"
)

func RaisePasswordChangeRequest(request *types.RaiseRequestDBQuery) error {
	database, err := Connect()
	if err != nil {
		return err
	}

	err = database.Table("requests_raised").Create(&request).Error
	return err
}

func FetchPendingRequests() ([]types.PendingRequests, error) {
	var pendingRequests []types.PendingRequests

	database, err := Connect()
	if err != nil {
		return nil, err
	}

	err = database.Table("requests_raised").Omit("accepted", "accepted_by", "admin_comment").Where("accepted = ?", false).Find(&pendingRequests).Error

	if err != nil {
		return nil, err
	}

	return pendingRequests, err

}
