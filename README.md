# gostduy

## clone 
### Go 분산 서비스 개발이라는 책에 있는 코드를 공부하면서 정리합니다. 
1. 로그 라이브러리를 작성
2. 네트워크 (gRPC 서비스, 보안, 시스템 관측) - 진행중
3. 분산 (서버 간 서비스 디스커버리, 합의를 통한 서비스 간 조율, 로드밸런싱)
4. 배포 (쿠버네티스 배포)

## architecture
clean architecture 기반으로 boilerplate 입니다. 
```
- config
- entity
- internal
  - user
    - controller
    - service
    - repository
- pkg
  - echo
  - util...
```

## util
-  nilasempty : json marshaling initialize nil slice as empty slice
   - (`sliceField: null` => `sliceField: []`)
## test
