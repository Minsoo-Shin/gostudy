package msgqueue

type Emitter interface {
	Emit(messages []Event)
	Close()
}
