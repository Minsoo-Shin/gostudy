# 객체 지향 설계 5가지 원칙 (SOLID)
## 1. SRP(Single Responsibility Principle) : 단일 책임 원칙
    #### 정의
    > 모든 객체는 책임을 하나만 져야한다.

    #### 이점
    > 코드의 재사용성을 높여줍니다.
    예를 들어, FinancialReport, MarketingReport 라는 객체에서 이메일을 통해 보고서를 전송하는 기능을 추가한다면,
    각 객체에 맞게 메서드를 만들지 않고, Report interface 자체를 만들고, Sender를 다른 인터페이스로 만들어 구분한다.
    Report 객체/ Sender 객체를 분리하게 되면 Sender 역시 여러 구현체를 만들 수 있다.
## 2. OCP(Open/Closed Principle) : 개방/폐쇄 원칙
    > 확장에는 열려 있고, 변경에는 닫혀 있다.
    > 이 역시 상호 결합도를 줄여 새 기능을 추가할 때, 기존 구현을 변경하지 않아도 된다.
    ``` go
    func SendReport(r *Report, method SendType, receiver string) {
        switch method {
        case Email:
            // 이메일 전송
        case Fax:
            // 팩스 전송
        case PDF:
            // PDF 전송
        case Printer:
            // 프린터 전송
        }
    }

    // =========
    type EmailSender struct{}
    // 인터페이스를 파라미터로 받으면 그에 맞는 행위를 한번에 정의할 수 있다.
    func (e *EmailSender) Send(r *Reporter) {
    // 이메일 전송
    }
    ```
## 3. LSP(Liskov Substitution Principle) : 리스코프 치환 원칙
    >  상위 타입인 Reporter을 함수 인자로 받는다면 당연히 호출자가 그의 하위타입을 인자로 넣더라도 작동해야한다.
    func (s *ReportSender) SendReport(report Reporter) {
    // 이런 구현체의 타입 캐스팅으로 구현되어있으면 Reporter 로 되어있는 객체들이 원하는대로 작동할 수 없다. (lsp 위반)
    // 함수 계약 관계를 잘 준수해야한다.
    // 상속을 지원하는 언어에서 보통 이런 문제가 많이 발생한다고 한다.
    // report.(*MarketingReport).Report()
    }
## 4. ISP(Interface Segregation Principle) : 인터페이스 분리 원칙
    > 클라이언트는 자기가 이용하지 않는 메서드에 의존하지 않아야 한다. 이로써 더 가볍게 인터페이스를 이용할 수 있다.
## 5. DIP(Dependency Inversion Principle) : 의존 역전 원칙
    > 상위 계층이 하위 계층에 의존하는 전통적인 의존 관계를 반전시킴으로써 상위 계층이 하위 계층의 구현으로부터 독립되게 할 수 있다.
    > 원칙 1: 상위 모듈은 하위 모듈에 의존해서는 안 된다. 둘 다 추상 모듈에 의존해야 한다.
    > 원칙 2: 추상 모듈은 구체화된 모듈에 의존해서는 안 된다. 구체화된 추상 모듈에 의존해야 한다.

