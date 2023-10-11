# AOP

Spring AOP 및 @Transactional 어노테이션 관련 블로그 글을 작성 도중 구현한 실습 코드입니다.

- **[토비의 스프링] 6장. AOP** 블로그  URL : https://dev-ote.tistory.com/entry/%ED%86%A0%EB%B9%84%EC%9D%98-%EC%8A%A4%ED%94%84%EB%A7%81-6%EC%9E%A5-AOP
- **[Spring] @Transactional** 블로그 URL  : https://dev-ote.tistory.com/entry/Spring-Transactional



## 테스트 소스

- `AopTests`
  - Client가 AOP 적용 타깃을 호출할 때에는 AOP가 적용되지만, 타깃에서 내부 메소드를 호출할 경우에는 AOP가 적용되지 않음을 확인함
- `EmployeeDaoTest`
  - 테스트 코드에서 `@Transactional` 부착 시 자동 롤백기능을 확인함
- `EmployeeServiceTest`
  - `@Transactional`의 대체 정책 적용에 대해 확인함
  - `Propagation.REQUIRED`와 `Propagation.REQUIRES_NEW`의 기존 트랜잭션 참여 여부 차이를 확인

