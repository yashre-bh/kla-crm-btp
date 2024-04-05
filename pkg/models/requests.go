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

	err = database.Table("requests_raised").Joins("JOIN employees ON employees.employee_id = requests_raised.request_from").Select("requests_raised.request_id", "requests_raised.title", "employees.name as request_from", "requests_raised.request_description", "requests_raised.request_date").Where("resolved = ?", false).Find(&pendingRequests).Error

	fmt.Println(pendingRequests)

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

	err = database.Table("requests_raised").Joins("JOIN employees ON employees.employee_id = requests_raised.request_from").Select("requests_raised.request_id", "requests_raised.title", "employees.name as request_from", "requests_raised.request_description", "requests_raised.request_date", "requests_raised.accepted", "requests_raised.resolved", "requests_raised.admin_comment", "requests_raised.accepted_by", "requests_raised.resolve_date").Where("resolved = ?", true).Find(&resolvedRequests).Error

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

	err = database.Table("requests_raised").Joins("JOIN employees ON employees.employee_id = requests_raised.request_from").Select("requests_raised.request_id", "requests_raised.title", "employees.name as request_from", "requests_raised.request_description", "requests_raised.request_date").Where("requests_raised.resolved = ?", false).Where("requests_raised.request_from = ?", employeeID).Find(&pendingRequests).Error

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

	err = database.Table("requests_raised").Joins("JOIN employees ON employees.employee_id = requests_raised.request_from").Select("requests_raised.request_id", "requests_raised.title", "employees.name as request_from", "requests_raised.request_description", "requests_raised.request_date", "requests_raised.accepted", "requests_raised.resolved", "requests_raised.admin_comment", "requests_raised.accepted_by", "requests_raised.resolve_date").Where("requests_raised.resolved = ?", true).Where("requests_raised.request_from = ?", employeeID).Find(&resolvedRequests).Error

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
