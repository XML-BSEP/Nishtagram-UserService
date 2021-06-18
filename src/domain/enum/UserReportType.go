package enum

type UserReportType int
const (
	Fake_Acc UserReportType = iota
	Under_age
	Inappropriate

)

func (g UserReportType) String() string {
	return [...]string{"Fake_Acc", "Under_age", "Inappropriate"}[g]
}

