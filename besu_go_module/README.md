

프로젝트의 config/besu 폴더를 WSL의 ~/besu 폴더와 연결

```sh
sudo mount -t drvfs C:/Users/조상혁/repo/Practice_Code/besu_go_module/config/besu ~/besu
```
:bulb: mounts it for the current session



docker container 실행
```sh
docker run -v /home/chosh901/besu/Node-2:/opt/besu/node hyperledger/besu --config-file=/opt/besu/node/config.toml
```



Node-1 위치에서 key 생성 명령어 수행
```sh
sudo docker run -v ./data:/opt/besu/data hyperledger/besu --data-path=data public-key export-address --to=data/node1Address
```

