package valueobjects

type SaleStatusType string

const (
	SaleStatusTypeApproved SaleStatusType = "APPROVED"
	SaleStatusTypePending  SaleStatusType = "PENDING"
)

func (ref SaleStatusType) String() string {
	return string(ref)
}
