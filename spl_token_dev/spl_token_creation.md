# SPL Token 생성 (dev net 기준)



### account 생성

```bash
solana-keygen grind --starts-with bos:1
solana-keygen grind --starts-with mnt:1
```

- bos로 시작하는 account는 payer 및 권한을 들고 있는 계정으로 사용
- mnt로 시작하는 account는 토큰 계정으로 사용



### 네트워크 설정

```bash
solana config set --url devnet
solana config set --keypair bosy1VC2BH2gh5fdXA3oKn53EuATLwapLWC4VR2sGHJ.json
```

- bos account로 작업들 수행



### 토큰 생성

```bash
spl-token create-token --program-id TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb --enable-metadata {mnt address}.json --decimals 6
```

- `--decimals` : token을 쪼갤 수 있는 단위 (6이면 토큰의 최소 단위 = 0.000001 token)
  - default 값은 9



### 메타 데이터 생성

- 메타 데이터 업로드 환경
  - 운영 환경에서는 decentralized storage 사용 필요
    - ex) [Filebase](https://filebase.com/)
  - dev 환경에서는 github와 같은 centralized storage 사용도 가능
    - github에 이미지와 `metadata.json` 파일 올리고 해당 경로 뒤에 `?raw=true` 붙일 시, 직접적으로 접근할 수 있는 url로 이동됨

- 메타 데이터

  - 토큰 이미지

  - `metadata.json`  

    ```json
    {
      "name": "Example Token",
      "symbol": "EXMPL",
      "description": "Example token from Solana Making a Token guide.",
      "image": "{image url}"
    }
    ```



### 메타 데이터 등록

```bash
spl-token initialize-metadata {mnt address} '{token name}' '{token symbol}' {metadata.json url}
```





### 민팅

- 민팅 받을 token address 생성

```bash
spl-token create-account {mnt address} [{target address}]
```

- 민팅

```bash
spl-token mint {mnt address} {mint amount} [{token address}]
```

- 민팅 권한 삭제 (fixed supply)

```bash
spl-token authorize {mnt address} mint --disable
```



### 전송

```bash
spl-token transfer mntTymSqMU4e1NEDdxJ9XoPN4MitCgQ7xxGW6AuRAWQ 10 (recipient wallet address) --fund-recipient
```

- `--fund-recipient` : allows you to pay to create the token account (i.e. the account rent) for the recipien