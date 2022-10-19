package response

type ConnectorStatus struct {
	Name      string    `json:"name"`
	Connector Connector `json:"connector"`
	Tasks     []Task    `json:"tasks"`
	Type      string    `json:"type"`
}

type Connector struct {
	State    string `json:"state"`
	WorkerID string `json:"worker_id"`
}

type Task struct {
	ID       int    `json:"id"`
	State    string `json:"state"`
	WorkerID string `json:"worker_id"`
}
