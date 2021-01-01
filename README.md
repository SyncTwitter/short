# short
短链接生成

### 安装

```bash
echo "export PATH=\${PATH}:\`go env GOPATH\`/bin" >> .bashrc 
source .bashrc
go get github.com/SyncTwitter/short
```

### 配置

```bash
curl -o config.yaml \
    https://raw.githubusercontent.com/SyncTwitter/short/master/config.yaml 
```

> 配置你自己的 mysql 和 redis 地址

### 运行

```
short
```