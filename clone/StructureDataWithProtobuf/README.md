# 네트워크를 통한 서비스 구축
네트워크를 통해서 여러 사람이 사용할 수 있는 서비스를 구축해볼 것이다. 네트워크로 서비스하면 세가지 이점을 제공한다.
- 가용성과 확장성을 위해 여러 컴퓨터에 걸쳐 실행할 수 있다.
- 여러 사람이 같은 데이터로 소통할 수 있다.
- 사람들이 쉽게 접근하고 사용할 수 있는 인터페이스를 제공한다.

현재 분산 서비스에 대한 요청을 처리하는 최고의 도구는 구글의 gRPC이다.

# gRPC란, 

## gRPC 서비스 정의
> gRPC란 관련 있는 RPC 엔드포인트들을 묶은 그룹이다. 
> 어떤 관련이 있는지 개발자 판단이다. 어떠한 문제를 해결하는 데 필요한 엔드포인트들이다. 

### service 키워드

--- 
컴파일러가 생성해야할 서비스라는 의미. rpc로 시작하는 각각의 줄은 서비스의 엔드포인트이다. 