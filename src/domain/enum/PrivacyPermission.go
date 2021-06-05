package enum

type PrivacyPermission int
const (
	Private PrivacyPermission = iota
	Public
	Banned
)

func (g PrivacyPermission) String() string {
	return [...]string{"Private", "Public", "Banned"}[g]
}
