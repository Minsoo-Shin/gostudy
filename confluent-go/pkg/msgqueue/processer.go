package msgqueue

type Processer interface {
	Process(event Event)
}
