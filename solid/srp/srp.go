package srp

// 그래서 아래와 같이 설계하기를 권고한다.
// Reporter <- FinanceReportV2, MarketingReportV2 구현
// ReportSender 로 책임을 분리한다.
type Reporter interface {
	Report() string
}

type FinanceReportV2 struct {
	report string
}

func (r *FinanceReportV2) Report() string {
	return r.report
}

//type MarketingReportV2 struct {
//	report string
//}
//
//func (r *MarketingReportV2) Report() string {
//	return r.report
//}

type ReportSender struct{}

func (s *ReportSender) SendReport(report Reporter) {
	// Reporter 라는 객체를 받아, ReportSender 라는 객체가 보내는 행위를 합니다.
	// Reporter 인터페이스는 모두 작동을 잘 할 수 있도록 작성해여합니다. (lsp)
}
