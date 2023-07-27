### Verify Code Service（VCS）

#### 说明

这是一个基于Go编写的简单的验证码验证服务，它的主要功能如下

```shell
1：基于此服务的API可以向指定的邮箱发送一条带有UUID和验证码的链接，需要用户点击邮箱内的链接可以验证这次回调，然后返回验证成功
```

#### 打包

```shell
# Windows
PS D:\Codes\callback_platform> go build -o vcs.exe .\cmd\main.go

# Linux
[root@localhost ~]# go build -o vcs cmd/main.go
```

#### 启动

```shell
PS D:\Codes\callback_platform> .\vcs.exe -h
Usage of D:\Codes\callback_platform\vcs.exe:
  -c string
        config file (default "config.yaml")

# 支持-c参数指定配置文件，这样我们的配置文件就不会因为被目录所限制
# Windows
PS D:\Codes\callback_platform> .\vcs.exe -c .\config.yaml

# Linux
[root@localhost ~]# vcs -c $HOME/config.yaml
```

#### 验证

```apl
GET /send/<example@admin.com>
{
	"message": "验证码已发送"
}
```

#### 配置

```yaml
# 默认配置
default:
  # 配置服务监听的端口
  listen: ":80"

# 邮件服务配置
smtp:
  # SMTP服务器地址
  host: "smtp.163.com"
  # SMTP服务端口
  port: "25"
  # SMTP服务账号
  user: "example@admin.com"
  # SMTP服务密码（并非邮箱密码）
  pass: "xxxxxxxxxxx"

domain:
  # 发送验证码时的地址，此地址的端口需要与default下的listen端口保持一致
  url: "http://localhost"
```

#### 配图

![image](https://img2023.cnblogs.com/blog/2222036/202307/2222036-20230727151629469-1215864237.png)

![image](https://img2023.cnblogs.com/blog/2222036/202307/2222036-20230727151717515-227026219.png)
