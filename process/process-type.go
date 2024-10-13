package process

type Process struct {
	Pid  string
	Name string
}

type ProcessDetails struct {
	PID     string  `json:"pid"`
	Command string  `json:"command"`
	Memory  float64 `json:"memory"` // In percentage
	CPU     float64 `json:"cpu"`    // In percentage
	UID     string  `json:"uid"`    // User ID
	GID     string  `json:"gid"`    // Group ID
	State   string  `json:"state"`  // Process state
}
