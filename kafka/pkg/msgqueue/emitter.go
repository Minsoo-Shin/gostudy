package msgqueue

type EventEmitter interface {
	Emit(req Event) error
}
