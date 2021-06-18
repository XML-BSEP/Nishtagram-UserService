package enum

type VerificationState int

const (
	Waiting VerificationState = iota
	Accepted
	Rejected

)

func (g VerificationState) String() string {
	return [...]string{"Waiting", "Accepted", "Rejected"}[g]
}

