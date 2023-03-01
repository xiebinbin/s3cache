# 一个可用s3作为后端的缓存服务器 也可作为图片服务器
# 支持webp jpp png bmp等图片实时转码
## 配置文件
```yaml
minio:
  endpoint: ""
  accessKeyId: ""
  secretAccessKey: ""
server:
  debug: false
  # 可以默认关闭
  remoteEnable: false
  dataPath: "绝对路径 不设置则默认为当前目录的 data目录"
  address: "0.0.0.0:8009"
```

# 请求示例

### 请求参数

```json
{
  "bucket": "cover",//必填
  "f": "webp", // 支持 jpg png bmp webp
  "w": 100,
  "h": 1500,
  "m": "resize" // 支持 thumbnail resize
}
```

```shell
/file/{文件路径}
```

```shell
https://xxx.com/file/31/3d/313d8ce3f1174bcb53374d7d4d.jpg?bucket=cover
# 源文件对应的目录为
{dataPath}/cover/31/3d/313d8ce3f11742ce927bcb53374d7d4d.jpg
```

### 参开supervisor配置
>image-pro.conf
### 参考nginx配置
>nginx.conf



### 必看

```shell
# ubuntu 安装
apt-get install musl-dev
```