package models

import (
	"fmt"

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

	err = database.Table("requests_raised").Omit("accepted", "accepted_by", "admin_comment", "resolve_date").Where("accepted = ?", false).Find(&pendingRequests).Error

	if err != nil {
		return nil, err
	}

	return pendingRequests, err

}

func FetchPendingRequestsOfEmployee(employeeID int32) ([]types.PendingRequests, error) {
	var pendingRequests []types.PendingRequests
	fmt.Println("phonch to gaye")

	database, err := Connect()
	if err != nil {
		return nil, err
	}

	err = database.Table("requests_raised").Omit("accepted", "accepted_by", "admin_comment", "resolve_date").Where("accepted = ?", false).Where("request_from = ?", employeeID).Find(&pendingRequests).Error

	if err != nil {
		return nil, err
	}

	return pendingRequests, err

}
