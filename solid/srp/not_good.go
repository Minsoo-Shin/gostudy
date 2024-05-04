package srp

import "fmt"

/*
아래는 단일 책임 원칙을 위배한 코드라고 할 수 있다.
FinanceReport 라는 객체는 온전히 Report 라는 책임을 가지고 있다 하지만, 여기서 전송이라는 책임까지 있는 것이다.
새로운 MarketingReport 객체가 추가되면, SendReport 메서드를 구현해야 한다.
그리고 만약 전송 방법이 변경된다면 모든 객체에 대해 변경을 해야할 것이다.
*/
type FinanceReport struct {
	report string
}

func (r *FinanceReport) SendReport(email string) {
	// send email
	fmt.Printf("send it to email(%s)\n", email)
}

type MarketingReport struct {
	report string
}

func (r *MarketingReport) SendReport(email string) {
	// send email
	fmt.Printf("send it to email(%s)\n", email)
}
