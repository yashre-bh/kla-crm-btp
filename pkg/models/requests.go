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

	err = database.Table("requests_raised").Omit("accepted", "accepted_by", "admin_comment", "resolve_date").Where("resolved = ?", false).Find(&pendingRequests).Error

	if err != nil {
		return nil, err
	}

	return pendingRequests, err

}

func FetchResolvedRequests() ([]types.ResolvedRequests, error) {
	var resolvedRequests []types.ResolvedRequests

	database, err := Connect()
	if err != nil {
		return nil, err
	}

	err = database.Table("requests_raised").Where("resolved = ?", true).Find(&resolvedRequests).Error

	if err != nil {
		return nil, err
	}

	return resolvedRequests, err
}

func FetchPendingRequestsOfEmployee(employeeID int32) ([]types.PendingRequests, error) {
	var pendingRequests []types.PendingRequests

	database, err := Connect()
	if err != nil {
		return nil, err
	}

	err = database.Table("requests_raised").Omit("accepted", "accepted_by", "admin_comment", "resolve_date", "resolved").Where("resloved = ?", false).Where("request_from = ?", employeeID).Find(&pendingRequests).Error

	if err != nil {
		return nil, err
	}

	return pendingRequests, err

}

func FetchResolvedRequestsOfEmployee(employeeID int32) ([]types.ResolvedRequests, error) {
	var resolvedRequests []types.ResolvedRequests

	database, err := Connect()
	if err != nil {
		return nil, err
	}

	err = database.Table("requests_raised").Where("resolved = ?", true).Where("request_from = ?", employeeID).Find(&resolvedRequests).Error

	if err != nil {
		return nil, err
	}

	return resolvedRequests, err

}

func ResolveByRequestID(resolveRequest types.ResolveRequestDBQuery) error {
	database, err := Connect()
	if err != nil {
		return err
	}

	err = database.Table("requests_raised").Where("request_id = ?", resolveRequest.RequestID).Updates(map[string]interface{}{
		"accepted":      resolveRequest.Accepted,
		"accepted_by":   resolveRequest.AcceptedBy,
		"admin_comment": resolveRequest.AdminComment,
		"resolve_date":  resolveRequest.ResolveDate,
		"resolved":      resolveRequest.Resolved,
	}).Error

	return err
}
