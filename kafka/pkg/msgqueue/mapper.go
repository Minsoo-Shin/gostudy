package msgqueue

type EventMapper interface {
	MapEvent(string, interface{}) (Event, error)
}
