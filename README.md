# 项目
[![编译状态](https://github.ruijc.com:20443/api/badges/dronestock/drone/status.svg)](https://github.ruijc.com:20443/dronestock/drone)
[![Golang质量](https://goreportcard.com/badge/github.com/dronestock/drone)](https://goreportcard.com/report/github.com/dronestock/drone)
![版本](https://img.shields.io/github/go-mod/go-version/dronestock/drone)
![仓库大小](https://img.shields.io/github/repo-size/dronestock/drone)
![最后提交](https://img.shields.io/github/last-commit/dronestock/drone)
![授权协议](https://img.shields.io/github/license/dronestock/drone)
![语言个数](https://img.shields.io/github/languages/count/dronestock/drone)
![最佳语言](https://img.shields.io/github/languages/top/dronestock/drone)
![星星个数](https://img.shields.io/github/stars/dronestock/drone?style=social)

Drone插件模板，要编写Drone插件，可以从此模板开始创建项目

## 使用

非常简单，只需要在`.drone.yml`里增加配置

```yaml
- name: 上传到腾讯云
  image: ccr.ccs.tencentyun.com/dronestock/cos
  settings:
    secret_id: xxx
    secret_key: xxx
```

更多使用教程，请参考[文档](https://www.dronestock.tech/plugin/cos)

## 交流

![微信群](https://www.dronestock.tech/communication/wxwork.jpg)

## 捐助

![支持宝](https://github.com/storezhang/donate/raw/master/alipay-small.jpg)
![微信](https://github.com/storezhang/donate/raw/master/weipay-small.jpg)

## 插件列表
- [Git](https://www.dronestock.tech/plugin/git)gg
- [Maven](https://www.dronestock.tech/plugin/maven)
- [Protobuf](https://www.dronestock.tech/plugin/protobuf)
- [Docker](https://www.dronestock.tech/plugin/docker)
- [Node](https://www.dronestock.tech/plugin/node)
- [MCU](https://www.dronestock.tech/plugin/mcu)

## 感谢Jetbrains

本项目通过`Jetbrains开源许可IDE`编写源代码，特此感谢

[![Jetbrains图标](https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.svg)](https://www.jetbrains.com/?from=dronestock/drone)
