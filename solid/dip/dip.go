package dip

/*
추상 모듈은 구체화된 모듈에 의존해서는 안 된다. 구체화된 모듈은 추상 모듈에 의존해야 한다.
*/

//type Mail struct {
//	alarm Alarm
//}
//
//type Alarm struct {
//}
//
//func (m *Mail) OnRecv() {
//	m.alarm.Alarm()
//}

type Event interface {
	Register(EventListener)
}

type EventListener interface {
	OnFire()
}

type Mail struct {
	listener EventListener
}

func (m *Mail) Register(listener EventListener) {
	m.listener = listener
}

func (m *Mail) OnRecv() {
	m.listener.OnFire()
}

type Alarm struct {
}

func (a *Alarm) OnFire() {
	// alarm
}

func AService() {
	var mail = &Mail{}
	var listener = &Alarm{}

	mail.Register(listener)
	mail.OnRecv()

}
