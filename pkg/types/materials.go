package types

import "time"

type IncomingRawMaterial struct {
	Name                string    `json:"name"`
	DateOfArrival       time.Time `json:"date_of_arrival"`
	VehicleNumber       string    `json:"vehicle_number"`
	BatchCode           string    `json:"batch_code"`
	Variety             string    `json:"variety"`
	ReceivedFrom        string    `json:"received_from"`
	Supplier            string    `json:"supplier"`
	WeightSupplier      float64   `json:"weight_supplier"`
	WeightWM            float64   `json:"weight_WM"`
	Rate                float64   `json:"rate"`
	Color               string    `json:"color,omitempty"`
	Texture             string    `json:"texture,omitempty"`
	Size                string    `json:"size,omitempty"`
	Maturity            string    `json:"maturity,omitempty"`
	Aroma               string    `json:"aroma,omitempty"`
	Appearance          string    `json:"appearance,omitempty"`
	WeightAccepted      float64   `json:"weight_accepted"`
	WeighmentSlipNumber string    `json:"weightment_slip_number"`
	QuantityRejected    float64   `json:"quantity_rejected"`
	Remarks             string    `json:"remarks,omitempty"`
}

type PurchaseRegister struct {
	OrderNumber        int     `json:"order_number"`
	OrderDate          string  `json:"order_date"`
	BrokerName         string  `json:"broker_name"`
	ProductName        string  `json:"product_name"`
	ConditionOfProduct string  `json:"condition_of_product"`
	Amount             float64 `json:"amount"`
	QtyBags            int     `json:"qty_bags"`
	QtyKgs             float64 `json:"qty_kgs"`
	VehicleNumber      string  `json:"vehicle_number"`
	Recovery           string  `json:"recovery"`
	BatchCode          string  `json:"batch_code"`
	DateReceived       string  `json:"date_received"`
	RejectReason       string  `json:"reject_reason,omitempty"`
	PurchasedBy        string  `json:"purchased_by"`
	Remark             string  `json:"remark,omitempty"`
}

type RawMaterialCode struct {
	EntityCode string `json:"entity_code"`
	Entity     string `json:"entity"`
}

type Batch struct {
	BatchCode  string `json:"batch_code"`
	Date       string `json:"date"`
	Dispatched bool   `json:"dispatched"`
	Entity     string `json:"entity"`
}

type MasterTracking struct {
	ActiveStatus          *bool      `json:"active_status,omitempty"`
	BatchCode             string     `json:"batch_code" gorm:"primaryKey"`
	DateAdded             *time.Time `json:"date_added,omitempty"`
	Checkpoint1Passed     *bool      `json:"checkpoint_1_passed,omitempty"`
	Checkpoint1CheckedBy  *int       `json:"checkpoint_1_checked_by,omitempty"`
	Checkpoint1VerifiedBy *int       `json:"checkpoint_1_verified_by,omitempty"`
	Checkpoint1ClearDate  *time.Time `json:"checkpoint_1_clear_date,omitempty"`
	Checkpoint2Passed     *bool      `json:"checkpoint_2_passed,omitempty"`
	Checkpoint2CheckedBy  *int       `json:"checkpoint_2_checked_by,omitempty"`
	Checkpoint2VerifiedBy *int       `json:"checkpoint_2_verified_by,omitempty"`
	Checkpoint2ClearDate  *time.Time `json:"checkpoint_2_clear_date,omitempty"`
	Checkpoint3Passed     *bool      `json:"checkpoint_3_passed,omitempty"`
	Checkpoint3CheckedBy  *int       `json:"checkpoint_3_checked_by,omitempty"`
	Checkpoint3VerifiedBy *int       `json:"checkpoint_3_verified_by,omitempty"`
	Checkpoint3ClearDate  *time.Time `json:"checkpoint_3_clear_date,omitempty"`
	Checkpoint4Passed     *bool      `json:"checkpoint_4_passed,omitempty"`
	Checkpoint4CheckedBy  *int       `json:"checkpoint_4_checked_by,omitempty"`
	Checkpoint4VerifiedBy *int       `json:"checkpoint_4_verified_by,omitempty"`
	UseByDate             *time.Time `json:"use_by_date,omitempty"`
	Checkpoint1Checked    *bool      `json:"checkpoint_1_checked,omitempty"`
	Checkpoint2Checked    *bool      `json:"checkpoint_2_checked,omitempty"`
	Checkpoint3Checked    *bool      `json:"checkpoint_3_checked,omitempty"`
	Checkpoint4Checked    *bool      `json:"checkpoint_4_checked,omitempty"`
	Checkpoint1Verified   *bool      `json:"checkpoint_1_verified,omitempty"`
	Checkpoint2Verified   *bool      `json:"checkpoint_2_verified,omitempty"`
	Checkpoint3Verified   *bool      `json:"checkpoint_3_verified,omitempty"`
	Checkpoint4Verified   *bool      `json:"checkpoint_4_verified,omitempty"`
}

type AddToTracking struct {
	BatchCode string     `json:"batch_code" gorm:"primaryKey"`
	DateAdded *time.Time `json:"date_added,omitempty"`
}
