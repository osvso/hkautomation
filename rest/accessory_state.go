package rest

const (
	HighState = 1 // on / open
	LowState  = 0 // off / closed
)

type AccessoryStateAction struct {
	Name string
	State int
}
