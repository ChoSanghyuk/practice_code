# TestConfig

`@TestConfiguration`에 관한 블로그 글을 작성하며 실습한 코드입니다.



`@TestConfiguration` 클래스를 통하여 테스트 시에만 생성되는 bean을 설정합니다.

테스트 시 `@Configuration`과 `@TestConfiguration`의 빈들의 충돌을 방지하고자, 빈의 이름을 다르게 지정하는 방식과, `@profile`룰 통해서 빈이 생성되는 환경을 지정하는 방식을 테스트합니다.



"[Spring] @TestConfiguration" 블로그 URL : https://dev-ote.tistory.com/27