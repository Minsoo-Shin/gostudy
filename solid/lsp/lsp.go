package lsp

type Reporter interface {
	Report() string
}

type FinanceReport struct {
	report string
}

func (r *FinanceReport) Report() string {
	return r.report
}

type MarketingReport struct {
	report string
}

func (r *MarketingReport) Report() string {
	return r.report
}

type ReportSender struct{}

func (s *ReportSender) SendReport(report Reporter) {
	// 이런 구현체의 타입 캐스팅으로 구현되어있으면 Reporter 로 되어있는 객체들이 원하는대로 작동할 수 없다. (lsp 위반)
	// 함수 계약 관계를 잘 준수해야한다.
	// 상속을 지원하는 언어에서 보통 이런 문제가 많이 발생한다고 한다.
	// report.(*MarketingReport).Report()
}
