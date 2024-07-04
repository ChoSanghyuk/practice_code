

프로젝트의 config/besu 폴더를 WSL의 ~/besu 폴더와 연결

```sh
sudo mount -t drvfs C:/Users/조상혁/repo/Practice_Code/besu_go_module/config/besu ~/besu
```

docker container 실행
```sh
docker run -v /home/chosh901/besu/:/opt/besu/config hyperledger/besu --config-file=/opt/besu/config/config.toml
```