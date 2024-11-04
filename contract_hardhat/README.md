# Hardhat Project



## 환경 세팅

### 설치

```bash
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.1/install.sh | bash
nvm install 22
nvm use 22
nvm alias default 22
npm install npm --global # Upgrade npm to the latest version
```

- 이때  `~/.bashrc` 를 생성해서 환경 변수를 등록했다면, `nvm`부터의 코드는 bash를 열어서 사용 



### hardhat project

프로젝트 폴더 생성 후 해당 폴더 위치에서 명령어 실행

```bash
npm init
npm install --save-dev hardhat
npx hardhat init
npm install --save-dev @nomicfoundation/hardhat-toolbox
```



:bulb: contract-sizer 설정

```bash
npm install --save-dev hardhat-contract-sizer
```

```ts
// hardhat.config.ts

import "hardhat-contract-sizer";

const config: HardhatUserConfig = {
  // ...
  contractSizer: {
    runOnCompile: true,
    only: [], // You can specify contracts here if needed
  },
};
```

```bash
npx hardhat size-contracts --no-compile
```





### compile

```bash
npx hardhat compile
```



### test

```shell
npx hardhat test
```

















