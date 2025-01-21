# Solana 성능 테스트



Solana 성능 테스트를 정리한 문서입니다.

테스트 기록에 앞서 테스트를 이해하는데 필요한 개념만 간략하게 정리하였습니다.



## Solana 개요



### Account

솔라나에서 모든 데이터는 'accounts'에 저장된다. accounts는 key-value store 형태로, address와 account 정보를 매핑하여 관리한다.

![Accounts](./assets/accounts.svg)

AccountInfo는 4개의 필드로 구성된다.

![AccountInfo](./assets/accountinfo.svg)

- `data` : 'account data'로도 불리우며, account의 상태(state) 혹은 실행 가능한 program[^1] 코드를 저장 (xor)
- `executable` : account가 program인지 나타내는 bool 값
- `lamports`: account의 lamport[^2] 잔고
- `owner` : account를 소유하고 있는 Program의 ID (public key)

[^1]: 솔라나의 smart contract
[^2]: SOL의 최소 단위로 1 SOL = 10억 lamports 



여기서 살펴보아야 할 것은, 첫번째로 account는 상태값을 저장하는 data account와 실행 코드를 저장하는 program으로 나누어진다. 이더리움은 하나의 contract 안에서 상태와 코드 모두를 관리했던 반면에, 솔라나는 상태와 코드를 별도의 account로 분리해서 관리한다.

두번째로, program account를 포함한 모든 account는 자신을 소유하는 program을 가진다. 소유자 program만이 account의 data를 변경하고 lamport 잔고를 줄일 수 있다(단, 잔고를 늘리는 것은 모두가 가능). 이는, account를 관리하는 주체가 private key를 가진 user account가 아닌 program임을 의미한다.

마지막으로, account의 data 필드는 byte array 타입으로 **고정된** 크기를 가진다. 이더리움은 메모리를 heap 영역에서 관리해서 contract가 저장한 데이터가 동적으로 계속 증가할 수 있었다. 하지만 솔라나는 고정되고 명확한 데이터 레이아웃을 요구하며, 모든 데이터를 stack에서 처리한다. 이로 인해 더 빠르며 더 적은 리소스 사용만으로 데이터를 처리할 수 있지만, 더 이상 하나의 smart contract에서 저장 데이터를 무한히 늘릴 수 없다.



### Program

Program은 솔라나에서의 smart contract이다. Program은 3가지 주요한 특징을 가진다.

먼저, Program은 로직을 담당하는 account로 한번만 배포된다. 이더리움에서는 실행 코드가 매 contract의 배포 때마다 같이 배포되었다. 하지만 솔라나에서는 실행 코드는 한번만 배포하고 이후 등록된 코드를 꺼내서 쓰는 구조를 가진다. 솔라나에서는 smart contract를 사용하기 위해서는 이미 온체인 상에 있는 program account의 ID 값만 가지고 와서 사용하면 된다.

또한, 솔라나에서 program은 기본적으로 '변경 가능'하다. Program 배포 시 upgrade 권한을 지정할 수 있으며, upgrade 권한이 null이 될 경우 변경 불가능(immutable)해진다.

끝으로, Program의 실행 코드는 'Instrction'이라는 하위 함수들로 조직되어 있다. Instruction은 transaction의 기본 단위가 되어 온체인에 특정 작업을 처리하도록 요청하는데 사용된다. 요청자는 여러 Instruction을 하나의 transaction으로 묶어 순차적이고 원자성을 보장하는 요청을 보낼 수 있다. 



Solana의 주요 프로그램 종류로는 다음과 같이 있다.

- Native Program
  - System Program : Account의 생성 및 Owner Program 지정 수행
  - BPF Loader : Native Program이 아닌 모든 Program의 owner로, program의 배포, 수정, 실행을 담당
- SPL Program : Solana 커뮤니티가 관리하는 표준 프로그램들의 집합
- Custom Program : 사용자 정의 프로그램



여기서 프로그램의 사용을 이더리움의 ERC-20과 같은 토큰 사용에 국한해서 살펴보겠다.

Solana 커뮤니티는 Solana에서의 여러 가지 표준 프로그램들을 모아 SPL(Solana Program Library)로 관리한다. SPL의 일환으로 SPL Token 프로그램이 제공되어, 신규 토큰 생성자는 SPL Token Program의 ID만 가지고 와서 쉽게 신규 토큰을 생성해서 사용할 수 있다. 하지만, 1. 'solana의 데이터와 로직의 분리'와 2.'고정된 account의 크기' 3.'SVM의 병렬처리[^3]'로 인해서 이더리움과 전혀 다른 토큰 구조를 가지게 된다.

[^3]: 솔라나의 SVM(Sealeve VM)은 트랜잭션들이 같은 account data를 덮어쓰지 않는다면 parallel로 처리한다.



#### SPL Token

토큰을 생성해서 사용하게 된다면, 다음과 같은 account 구조가 나오게 된다.

![image-20250118155217139](./assets/image-20250118155217139.png)

- Program Account[^4]
  - SPL Token Program account로, Token의 실행 로직이 들어있다.
- Mint Account
  - 생성된 토큰 그 자체를 나타내는 account
  - 토큰에 대한 메타 정보 (mintAuthority, freezeAuthority, totalBalance 등)을 저장한다.
  - :bulb: 단, 누가 얼마큼을 가지고 있는지는 저장하고 있지 않다.
- Token Account
  - 개개인이 들고 있는 토큰의 수량 정보 및 메타 정보를 저장한다.
  - 주로 랜덤한 주소가 아닌, user account의 주소와 mint account의 주소를 해시하여 생성한다.
- User Account
  - system account 혹은 native account라고 불리면 일반적인 사용자의 account를 의미한다.

[^4]: Program account에 대한 설명은 편의상 약식으로 표현됨



이러한 구조적인 차이로 인해서, smart contract를 이더리움과 동일하게 동작하도록 작성하는 것은 불가능하다(매우 권장되지 않는다). 





## SPL 토큰 테스트 시나리오



SPL 토큰에 대한 테스트로, 기본적인 토큰의 동작인 deploy, mint, transfer, query 테스트 수행한다.



### 테스트 설계

#### 계정 세팅

#### account 종류

- payer AC
  - transaction 비용 내는 계정
- funding AC
  - mint account가 유지되는 비용을 내는 계정
- freeze auth AC
  - 해당 토큰에 대한 Lock 권한을 가진 계정
- mint auth AC
  - 해당 토큰의 minting 권한을 가진 계정

- mint AC
  - 토큰에 대한 전체적인 정보가 담긴 계정
  - token itself

- user AC
  - 일반 solana 계정
  - 여기서는 최초 minting의 대상이 될 `init_holder`와 trasnfer의 대상이 될 `target`으로 나누어 사용

- token AC (= ata)
  - user AC가 들고 있는 토큰의 양이 저장되는 계정
  - 주로 ata(Associated Token Address) 계정으로 사용

#### account 세팅

- **init_holder** = payer = funding = freezeAuth = mintAuth
  - 최초 minting의 될 init_holder가 token에 대한 모든 지불 및 권한 수행
- config에서 account 개수 사전 지정
  - N : mint AC init_holder 개수
  - M : target 개수





### 테스트 시나리오

#### 개요

| 시나리오 |               url               |             parameter              |
| :------: | :-----------------------------: | :--------------------------------: |
|  deploy  | `[post] /spl/deploy-with-mint ` | `[json] { amount : ${int_value} }` |
| minting  |       `[post] /spl/mint `       | `[json] { amount : ${int_value} }` |
| transfer |     `[post] /spl/transfer `     | `[json] { amount : ${int_value} }` |
|  query   |       `[get] /spl/query `       |                                    |

- minting 또는 transfer 수행 전에는 `[post] /spl/set-mint-account` 수행 필요 (parameter 필요 없음)

- query 수행 전에는 충분히 많은 (N *M) 수 의 minting 혹은 transfer 수행 필요



#### 1. deploy(+minting) 시나리오

- 주어진 시간 내에 몇 개의 토큰 deploy + 최초 minting 트랜잭션 처리 확인
- 계정에 대한 payer 수(N)를 변수로 설정하고 parallel 수행 여부 확인

:bulb: 해당 케이스에선 mint account는 N개로 제한하지 않고, 무한히 생성

|  N   | 테스트 시간(min) | MTT(ms) |  TPS  | 비고 |
| :--: | :--------------: | :-----: | :---: | :--: |
|  1   |        2         |  2054   |  430  |      |
| 100  |        2         |  1529   | 580.9 |      |
| 1000 |        2         |  1522   | 580.2 |      |



#### 2. transfer 시나리오

- deploy된 token에 대해서 target account에 transfer 수행
- init_holder 수(N)를 변수로 설정하고 parallel 수행 여부에 따른 TPS 확인

|  N   |  M   | 테스트 시간(min) | MTT(ms) | TPS  |
| :--: | :--: | :--------------: | :-----: | :--: |
|  1   | 1000 |        2         |  5353   | 163  |
| 100  | 1000 |        2         |   548   | 1632 |
| 1000 | 1000 |        2         |   549   | 1785 |
| 1000 |  1   |                  |         |      |



#### 3. query 시나리오

- N개의 init_holder에 대해서 balance 조회 진행

|  N   | 테스트 시간(min) | MTT(ms) |  TPS  |
| :--: | :--------------: | :-----: | :---: |
|  1   |        2         |  11.39  | 86268 |
| 100  |        2         |  12.17  | 73257 |
| 1000 |        2         |         |       |







### 추후 테스트 일정

- 트랜잭션 signature를 모두 저장해 둔 후, confirmed 이후 finalized 이전 dropped되는 트랜잭션 존재 여부 확인
- 노드 파라미터 조정에 따른 성능 변화 확인
- 멀티 노드 환경에서의 성능 변화 확인











