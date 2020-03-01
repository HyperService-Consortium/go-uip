package isc

const (
	StateInitializing uint8 = 0 + iota
	StateInitialized
	StateOpening
	StateSettling
	StateClosed
)
