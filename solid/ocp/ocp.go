package ocp

/*
개방-폐쇄 원칙 (OCP: Open-Closed Principle)
정의 확장에는 열령 있고, 변경에는 닫혀 있다.
-> 상호 결합도를 줄여 새 기능을 추가할 떄 기존 구현을 변경하지 않아도 된다.
*/

//func SendReport(r *Report, method SendType, receiver string) {
//	switch method {
//	case Email:
//		// 이메일 전송
//	case Fax:
//		// 팩스 전송
//	case PDF:
//		// PDF 전송
//	case Printer:
//		// 프린터 전송
//	}
//}

/*
아래 코드는 ReportSender 라는 인터페이스를 만들어서,
다른 Email, Fax, 등과 같은 구현체들을 만들어서 각자 구현하기 좋도록 되어있다.
*/
type Reporter interface {
	Report() string
}

type ReportSender interface {
	Send(r *Reporter)
}

type EmailSender struct{}

func (e *EmailSender) Send(r *Reporter) {
	// 이메일 전송
}

//
//type FaxSender struct {}
//
//func (f *FaxSender) Send(r *Reporter) {
//	// 팩스 전송
//}
