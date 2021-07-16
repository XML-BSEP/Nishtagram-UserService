package enum

type ReportStatus int

const (
	Created ReportStatus = iota
	Declined
	Approved

)

func (g ReportStatus) String() string {
	return [...]string{"Created", "Declined", "Approved"}[g]
}
