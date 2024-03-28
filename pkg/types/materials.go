package types

type IncomingRawMaterial struct {
	Name             string  `json:"name"`
	DateOfArrival    string  `json:"date_of_arrival"`
	VehicleNumber    string  `json:"vehicle_number"`
	LotNumber        string  `json:"lot_number"`
	Variety          string  `json:"variety"`
	ReceivedFrom     string  `json:"received_from"`
	Supplier         string  `json:"supplier"`
	WeightSupplier   float64 `json:"weight_supplier"`
	WeightWM         float64 `json:"weight_WM"`
	Rate             float64 `json:"rate"`
	Color            string  `json:"color,omitempty"`
	Texture          string  `json:"texture,omitempty"`
	Size             string  `json:"size,omitempty"`
	Maturity         string  `json:"maturity,omitempty"`
	Aroma            string  `json:"aroma,omitempty"`
	Appearance       string  `json:"appearance,omitempty"`
	WeightAccepted   float64 `json:"weight_accepted"`
	QuantityRejected float64 `json:"quantity_rejected"`
	Remarks          string  `json:"remarks,omitempty"`
	CheckedBy        int     `json:"checked_by"`
	VerifiedBy       int     `json:"verified_by"`
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
	LotNumber          string  `json:"lot_number"`
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
