package domain

type Inspection struct {
	ID            string `json:"id"`
	DoorName      string `json:"door_name"`
	Site          string `json:"site"`
	Status        string `json:"status"`
	LastInspected string `json:"last_inspected"`
	DefectCount   int    `json:"defect_count"`
}

type StatusChange struct {
	Status string `json:"status"`
}
