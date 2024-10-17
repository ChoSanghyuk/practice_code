

프로젝트의 config/besu 폴더를 WSL의 ~/besu 폴더와 연결

```sh
sudo mount -t drvfs C:/Users/조상혁/repo/Practice_Code/besu_go_module/config/besu ~/besu
```
:bulb: mounts it for the current session



docker container 실행
```sh
docker run -v /home/chosh901/besu/Node-1:/opt/besu/node hyperledger/besu --config-file=/opt/besu/node/config.toml
```
```sh
besu --data-path=data --genesis-file=../cliqueGenesis.json  --network-id 123 --p2p-port=30304 --rpc-http-enabled --rpc-http-api=ETH,NET,CLIQUE --host-allowlist="*" --rpc-http-cors-origins="all" --rpc-http-port=8546
```
```sh
docker run -v ./:/opt/besu/node --name testBesu --net besu_bridge hyperledger/besu --config-file=/opt/besu/node/config.toml
```


Node-1 위치에서 key 생성 명령어 수행
```sh
sudo docker run -v ./data:/opt/besu/data hyperledger/besu --data-path=data public-key export-address --to=data/node1Address
```


docker container로 besu 노드를 띄우는 경우에는 docker Network 설정이 필수적임
네트워크 설정 없이, container를 띄우게 되면, 로그 상에는 127.0.0.1(localhost)로 p2p-host와 enode url이 표시되지만, 이는 컨테이너 내부에서 localhost로 기동된 것일 뿐, 실제로는 container 자체 주소를 가지어, 해당 주소로 접근해야 함.
따라서, network_mode를 host로 설정하여, 호스트의 네트워크 환경을 그대로 사용하게 함
(단, 이때는 port-binding 사용 불가)



wsl -d Ubuntu

curl -X POST --data '{"jsonrpc":"2.0","method":"clique_propose","params":["0xd6a98d4f51e8072ef3c994fcd619e28c6072c7cf", true], "id":1}' localhost:8545   
docker run -v ./:/opt/besu/node --name testBesu2 --net besu_bridge -p 8547:8547 hyperledger/besu --config-file=/opt/besu/node/config.toml                


ping -c 1 172.18.0.1  
docker network inspect besu_bridge                                                                                                                

docker inspect testBesu | grep IPAddress                                                                                                          
- 이때 나오는 주소는 내부 영역(?)에서 사용되는 컨테이너의 IP 주소로 외부에서 접근하는 용도가 아님
- 해당 주소는 같은 네트워크를 사용하는 컨테이너끼리 통신하기 위한 주소
- 컨테이너끼리는 같은 네트워크를 사용하고 있어야 통신 가능

포트 바인딩
- 해당 컨테이너를 외부에서도 접근할 수 있도록 호스트의 포트와 컨테이너의 포트를 바인딩하는 것
- 따라서, host network mode가 아니더라도 localhost:포트 로 요청해야 컨테이너에 요청 가능함

네트워크 gateway
- docker network inspect {네트워크} 했을 때 나오는 gateway 값은 the gateway serves as the default route for containers to communicate with other networks, including the host's network and the internet. This reserved IP address ensures that there's a dedicated point for routing traffic in and out of the Docker network.

yaml file
Use of -: Used to denote sequence items or elements in lists under a key in YAML.
Keys in YAML: Defined without -, followed by a colon (:) and then the value.



docker pull ethereum/solc:0.8.6

```
version: '3.8'

networks:
  besu_bridge:
    driver: bridge
    ipam:
      driver: default
      config:
        - subnet: 172.18.0.0/16   # Example subnet for the custom bridge network

services:
  testBesu:
    image: your-besu-image:tag
    networks:
      - besu_bridge
    # other service configuration options
    # ...
```



docker run -v ./:/compile/ ethereum/solc:0.8.6 -o /compile/output --abi --bin ./compile/AgeInfoStorage.sol

docker run -v ./:/compile/ ethereum/solc:0.8.20 -o /compile/upgradeable/proxy/output --abi --bin @openzeppelin=/compile/node_modules/@openzeppelin/ ./compile/upgradeable/proxy/MyProxy.sol
- library를 import해서 사용하는 경우, 별도로 패키지의 경로 지정이 필요함

docker pull ethereum/client-go
공식 ethereum/client-go 이미지에 abigen 설치되어 있지 않음. 별도 설정 필요
```
FROM golang:1.20-alpine

# Install dependencies
RUN apk add --no-cache git make gcc musl-dev

# Install go-ethereum which includes abigen
RUN go install github.com/ethereum/go-ethereum/cmd/abigen@latest

# Set up the entrypoint to use abigen by default
ENTRYPOINT ["abigen"]
```
=> 그냥 로컬에 go install로 설치하는 것이 더 간편
go install github.com/ethereum/go-ethereum/cmd/abigen@latest

~/go/bin/abigen --bin=./output/AgeInfoStorage.bin --abi=./output/AgeInfoStorage.abi --pkg=sample_contract --out=./AgeInfoStorage.go


:bulb: docker run -it  로 실행하지 않았다면 ctrl + p + q 로 빠져나오는거 안됨




curl -X POST -H "Content-Type: application/json" --data '{ "query": "{__schema { queryType { fields { name } }}}"}' http://172.18.0.2:8547/graphql

port binding 없이는 로컬 호스트에서 http 요청으로 접근할 순 없음!
For a port to be accessible to containers or non-docker hosts on different networks, that port must be published using the -p or --publish flag
docs.docker.com/network/drivers/bridge


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

