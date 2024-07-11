

프로젝트의 config/besu 폴더를 WSL의 ~/besu 폴더와 연결

```sh
sudo mount -t drvfs C:/Users/조상혁/repo/Practice_Code/besu_go_module/config/besu ~/besu
```
:bulb: mounts it for the current session



docker container 실행
```sh
docker run -v /home/chosh901/besu/Node-1:/opt/besu/node hyperledger/besu --config-file=/opt/besu/node/config.toml
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