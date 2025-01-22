package flow

type Tasker interface {
	GetStatus() FlowStatus
	SetStatus(status FlowStatus) error
}
