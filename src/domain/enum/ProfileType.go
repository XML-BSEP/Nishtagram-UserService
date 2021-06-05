package enum

type ProfileType int
const (
	User ProfileType = iota
	Agent

)


func (g ProfileType) String() string {
	return [...]string{"User", "Agent"}[g]
}



