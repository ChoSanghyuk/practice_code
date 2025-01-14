# Solana Go Client



## 토큰 테스트 시나리오



SPL 토큰에 대한 테스트로, 기본적인 토큰의 동작인 deploy, mint, transfer, query 테스트 수행



### 계정 세팅

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

|  N   | 테스트 시간(min) | 처리 트랜잭션 | TPS  |
| :--: | :--------------: | :-----------: | :--: |
|      |                  |               |      |
|      |                  |               |      |
|      |                  |               |      |



#### 2. minting 시나리오

- deploy된 token에 대해서 target account로 추가 minting 수행
- mint account 수(N)와 target account 수(M)를 변수로 설정하고 parallel 수행 여부 확인

|  N   |  M   | 테스트 시간(min) | 처리 트랜잭션 | TPS  |
| :--: | :--: | :--------------: | :-----------: | :--: |
|      |      |                  |               |      |
|      |      |                  |               |      |
|      |      |                  |               |      |



#### 3. transfer 시나리오

- deploy된 token에 대해서 target account에 transfer 수행
- init_holder 수(N)를 변수로 설정하고 parallel 수행 여부에 따른 TPS 확인

|  N   | 테스트 시간(min) | 처리 트랜잭션 | TPS  |
| :--: | :--------------: | :-----------: | :--: |
|      |                  |               |      |
|      |                  |               |      |
|      |                  |               |      |



#### 4. query 시나리오

- 2번 혹은 3번 시나리오 수행 직후 target account에 대해서 token account의 balance 조회

- 단, M개의 target account가 N개의 token에 대해 전부 한번이라도 mint 혹은 transfer를 수행한 이후 수행 필요

| 차수 | 테스트 시간(min) | 처리 트랜잭션 | TPS  |
| :--: | :--------------: | :-----------: | :--: |
|      |                  |               |      |
|      |                  |               |      |

(N:  / M: )





### 추후 테스트 일정

- 트랜잭션 signature를 모두 저장해 둔 후, confirmed 이후 finalized 이전 dropped되는 트랜잭션 존재 여부 확인
- 노드 파라미터 조정에 따른 성능 변화 확인
- 멀티 노드 환경에서의 성능 변화 확인











