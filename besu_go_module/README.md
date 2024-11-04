# Besu Go Module



## Besu 설정



### WSL 환경

프로젝트의 config/besu 폴더를 WSL의 ~/besu 폴더와 연결

```sh
sudo mount -t drvfs C:/Users/조상혁/repo/Practice_Code/besu_go_module/config/besu ~/besu
```
:bulb: mounts it for the current session

```sh
besu --data-path=data --genesis-file=../cliqueGenesis.json  --network-id 123 --p2p-port=30304 --rpc-http-enabled --rpc-http-api=ETH,NET,CLIQUE --host-allowlist="*" --rpc-http-cors-origins="all" --rpc-http-port=8546
```



### docker 환경

#### docker container 실행 - flag

- 노드 실행

```sh
docker run -v ./:/opt/besu/node --name testBesu --net besu_bridge hyperledger/besu --config-file=/opt/besu/node/config.toml
```



#### key 생성 명령어 수행

```sh
sudo docker run -v ./data:/opt/besu/data hyperledger/besu --data-path=data public-key export-address --to=data/node1Address
```
- 키 기반으로 주소 다시 생성

```sh
sudo docker run -v ./data:/opt/besu/data hyperledger/besu --data-path=data public-key export-address --node-private-key-file=/opt/besu/data/key --to=data/node1Address
```
- qbft에서 사용하는 형식으로 extra_data 인코딩

```sh
docker run -v ./:/opt/besu/data hyperledger/besu rlp encode --from=/opt/besu/data/toEncode.json --type=QBFT_EXTRA_DATA
```


:bulb: **docker network**

- docker container로 besu 노드를 띄우는 경우에는 docker Network 설정이 필수적임
  - 네트워크 설정 없이 container를 띄우더라도 컨테이너는 자체적인 주소를 가지고 있음
  - 내부 로그 상에는 127.0.0.1(localhost)로 p2p-host와 enode url이 표시되지만, 이는 컨테이너 내부에서 localhost로 기동된 것일 뿐
  - 또한 **컨테이너끼리는 같은 네트워크를 사용하고 있어야 통신 가능**
- 설정 방법
  1. network_mode를 host로 설정하여, 호스트의 네트워크 환경을 그대로 사용하게 함
     (단, 이때는 port-binding 사용 불가)
  2. bridge network mode 사용하여, 컨테이너 네트워크 구성
     - 네트워크 대역을 설정하고, 컨테이너의 ip를 사전에 지정 가능
     - 포트 바인딩으로 외부 ip:port와 컨테이너 내부 localhost:port를 바인딩 가능



#### docker container 실행 - docker-compose

```yaml
# version '3.9'

networks:
  custom_bridge: # docker compose는 name을 유니크하게 유지하게 하기 위해, 현재 폴더명을 네트워크 이름 앞에 붙임
    driver: bridge
    ipam:
      driver: default
      config:
        - subnet: 172.18.0.0/16

services:
  Node-1:
    image: hyperledger/besu:24.9.1
    container_name: Node-1
    networks: 
      custom_bridge:
        ipv4_address: 172.18.0.2
    ports:
      - "30303:30303"
      - "8545:8545"
      - "8547:8547"
    volumes:
      - ./Node-1:/opt/besu/node
      - ./genesis:/opt/besu/genesis
    command: --config-file=/opt/besu/node/config.toml
  Node-2:
    image: hyperledger/besu:24.9.1
    container_name: Node-2
    networks:
      custom_bridge:
        ipv4_address: 172.18.0.3
    ports:
      - "30304:30303"
      - "18545:8545"
      - "18547:8547"
    volumes:
      - ./Node-2:/opt/besu/node
      - ./genesis:/opt/besu/genesis
    command: --config-file=/opt/besu/node/config.toml
  Node-3:
    image: hyperledger/besu:24.9.1
    container_name: Node-3
    networks:
      custom_bridge:
        ipv4_address: 172.18.0.4
    ports:
      - "30305:30303"
      - "28545:8545"
      - "28547:8547"
    volumes:
      - ./Node-3:/opt/besu/node
      - ./genesis:/opt/besu/genesis
    command: --config-file=/opt/besu/node/config.toml
```

- yaml에서의 `-` : key 아래에 list 나열



**:bulb: network inspect**

`docker network inspect besu_bridge`

`docker inspect testBesu | grep IPAddress`

- 이때 나오는 주소는 네트워크 내부에서 사용되는 컨테이너의 IP 주소로 외부에서 접근하는 용도가 아님
- 해당 주소는 같은 네트워크를 사용하는 컨테이너끼리 통신하기 위한 주소

`docker network inspect {네트워크} `

- 이때 나오는 gateway 값은 다른 네트워크의 컨테이너로 가는 디폴트 경로
- ensures that there's a dedicated point for routing traffic in and out of the Docker network.





## 컨트랙트 컴파일



### abi, bin 파일 생성

- 기본

````sh
docker run -v ./:/compile/ ethereum/solc:0.8.6 -o /compile/output --abi --bin ./compile/AgeInfoStorage.sol
````

- 외부 라이브러리 지정

```sh
docker run -v ./:/compile/ ethereum/solc:0.8.20 -o /compile/upgradeable/proxy/output --abi --bin @openzeppelin=/compile/node_modules/@openzeppelin/ ./compile/upgradeable/proxy/MyProxy.sol
```

- library를 import해서 사용하는 경우, 별도로 패키지의 경로 지정이 필요함



### Go bind 파일 생성

#### docker 사용하는 경우

```sh
docker pull ethereum/client-go
```

- 공식 ethereum/client-go 이미지에 abigen 설치되어 있지 않음. 별도 설정 필요

  ```sh
FROM golang:1.20-alpine

# Install dependencies
RUN apk add --no-cache git make gcc musl-dev

# Install go-ethereum which includes abigen
RUN go install github.com/ethereum/go-ethereum/cmd/abigen@latest

# Set up the entrypoint to use abigen by default
ENTRYPOINT ["abigen"]
  ```
  => 그냥 로컬에 go install로 설치하는 것이 더 간편



#### local 설치

```sh
go install github.com/ethereum/go-ethereum/cmd/abigen@latest

~/go/bin/abigen --bin=./output/AgeInfoStorage.bin --abi=./output/AgeInfoStorage.abi --pkg=sample_contract --out=./AgeInfoStorage.go
```

- 이때, pkg로 지정한 것으로 클래스 이름 지정됨. 



## package managing

### yarn
- 설치
```
brew install node
brew install yarn
```
- 사용
```
yarn init
yarn add @openzeppelin/contracts
```









TODO.  -- 성공하는 거보고 다 지워

- solidity 0.8.20으로 컴파일한 컨트랙트 berlin milestone에서 배포 실패
  - => mirae에서 사용한 milestone 확인. + solidity version
  
  - 흐음.... 합의 알고리즘 차이인가
  
    ```json
    "berlinBlock" : 0,
    "qbft" : {
        "blockperiodseconds" : 3,
        "epochlength" : 30000,
        "requesttimeoutseconds" : 6
    }
    ```
  
- genesis block의 Pos 적용된 버전에서의 validator 동작 및 초기 alloc 설정 (전반적인 genesis file 작성법)

- genesis file에서 milestone 지정하면, 초기 besu 노드 올라갈 때 표시됨  (근데 parisBlock 할때는 frontier로 회귀해버림. 별도 작성법 있는지 확인 필요)

- hardhat ts test 코드 템플릿
```
2024-10-21 05:36:39.199+00:00 | main | INFO  | ProtocolScheduleBuilder | Protocol schedule created with milestones: [Berlin:0]
```



Frontier : 0.1.x - 0.2.x
Homestead : 0.2.x - 0.4.0
Byzantium : 0.4.21 - 0.5.x
Constantinople and Petersburg : 0.5.x
Istanbul : 0.5.x - 0.6.x
Berlin : 0.6.x - 0.8.4
London : 0.8.5 - 0.8.9
Paris : 0.8.15 - 0.8.20
Shanghai : 0.8.19+
Cancun : 0.8.21+