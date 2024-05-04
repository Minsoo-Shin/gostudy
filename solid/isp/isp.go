package isp

import (
	"fmt"
	"time"
)

type Report interface {
	Report() string
	Pages() int
	Author() string
	WrittenDate() time.Time
}

func SendReport(r Report) {
	fmt.Println(r.Report())
}

type report struct{}

func (r report) Report() string {
	//TODO implement me
	panic("implement me")
}

func (r report) Pages() int {
	//TODO implement me
	panic("implement me")
}

func (r report) Author() string {
	//TODO implement me
	panic("implement me")
}

func (r report) WrittenDate() time.Time {
	//TODO implement me
	panic("implement me")
}

func NewReport() Report {
	return &report{}
}

// isp
// 정의 : 클라이언트는 자신이 이용하지 않는 메서드에 의존하지 않아야 한다.
// 이점: 인터페이스를 분리하면 불필요한 메서드들과 의존 관계가 끊어져 더 가볍게 인터페이스를 이용할 수 있습니다.
// 아래와 같이 SendReport 라는 인터페이스를 이용하기 위해서 불필요한 메서드를 만들 필요가 없어집니다.

type ReportV2 interface {
	Report() string
}

type WrittenInfo interface {
	Pages() int
	Author() string
	WrittenDate() time.Time
}

func SendReportV2(r ReportV2) {
	fmt.Println(r.Report())
}

type reportV2 struct{}

func NewReportV2() ReportV2 {
	return &reportV2{}
}

func (r reportV2) Report() string {
	//TODO implement me
	panic("implement me")
}
