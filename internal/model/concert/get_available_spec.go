package concert

const (
	defaultLimit uint64 = 30
)

// GetAvailableSpec hold request params to get available concerts list
type GetAvailableSpec struct {
	Search string
	Offset uint64
	Limit  uint64
}

// GetLimit get limit set on spec but if empty return defaultLimit
func (s GetAvailableSpec) GetLimit() uint64 {
	if s.Limit == 0 {
		return defaultLimit
	}
	return s.Limit
}
