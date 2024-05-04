package srp

import "testing"

func Test_SendReport(t *testing.T) {
	fr := FinanceReport{report: "coding is happy"}
	fr.SendReport("happycoding@gmail.com")

	// i want to make marketing report
	mr := MarketingReport{report: "this ad is better"}
	mr.SendReport("happycoding@gmail.com")

	// i want to make something report
	//sr := SomethingReport{report: "something report"}
	//sr.SendReport("somethingreport@gmail.com")
}
