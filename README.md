# certctl
此工具每天检测一次一些域名如果距离到期小于指定天数则使用工具certbot提前生成证书并上传到阿里云ssl应用然后发送信息到本人提示作者部署到相关云资源

使用有疑惑联系：2791923714

### 支持功能

---

- 检测证书到期时间
- 生成证书
- 上传证书
- 以上自动全流程

### 使用前准备

---

1、服务器需要有go环境,

2、安装使用过一次certbot工具（支持阿里云）

3、创建 certbot-dns-aliyun 配置文件

```bash
git clone https://github.com/LiuXiangBiao/certctl.git
```
```bash
cd certctl
```
```bash
go mod tidy
```
##### ！！！进入到conf目录编辑conf.yaml文件 
```
domains_file_path: 域名文件的绝对路径（文件内一行一个域名）
distance_day_time: 证书到期距离今日的天数
certbot_config_file_path: certbot的aksk配置文件
accessKeyId: ak
accessKeySecret: sk
token: 机器人webhook
secret: 机器人标签
```

- 编译
```bash
go build -o certctl
```
```bash
mv certctl /gopath/bin/
```
### 执行命令示例

---

- 自动检测指定文件的一些域名 （全流程）
```bash
 certctl  
```

### 以下是单个流程执行命令

---

- 检测某些域名到期时间
```bash
  certctl  check  [域名]......
```

- 生成指定域名证书和私钥文件
```bash
  certctl  generate  [aksk 配置文件路径]  [域名]
```

- 上传本地证书文件到阿里云ssl应用服务
```bash
   certctl  upload  [AccessKey ID]  [AccessKey Secret]  [证书名字]  [证书文件路径]  [私钥文件路径]
```
