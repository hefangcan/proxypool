<h1 align="center">
  <br>proxypool<br>
</h1>

<h5 align="center">自动抓取tg频道、订阅地址、公开互联网上的ss、ssr、vmess、trojan、vless节点信息，去重测试可用性后提供节点列表</h5>

## 支持

- 支持ss、ssr、vmess、trojan、vless多种类型
- Telegram频道抓取
- 订阅地址抓取解析
- 公开互联网页面模糊抓取
- 定时抓取自动更新
- 通过配置文件设置抓取源
- 自动检测节点可用性（不支持vless，H2）
- 提供clash、surge配置文件
- 提供ss、ssr、vmess、sip002订阅

## 安装


需要安装Golang

```shell
$ git clone https://github.com/hefangcan/proxypool.git
```

运行

```shell
$ go run main.go -c ./config/config.yaml
```

编译

```shell
$ make
```

## 使用

运行该程序需要具有访问完整互联网的能力。

### 修改配置文件

首先修改 config.yaml 中的必要配置信息。带有默认值的字段均可不填写。完整的配置选项见[配置文件说明](https://github.com/bh-qt/proxypool/wiki/%E9%85%8D%E7%BD%AE%E6%96%87%E4%BB%B6%E8%AF%B4%E6%98%8E)

### 启动程序

使用 `-c` 参数指定配置文件路径，支持http链接

```shell
$ proxypool -c ./config/config.yaml
```

如果需要部署到VPS，更多细节请[查看wiki](https://github.com/bh-qt/proxypool/wiki/%E9%83%A8%E7%BD%B2%E5%88%B0VPS-Step-by-Step)。

## Clash配置文件

远程部署时Clash配置文件访问：<https://domain/clash/config>

本地运行时Clash配置文件访问：<http://127.0.0.1:[端口]/clash/localconfig>

查看所有节点信息查看：<https://domain/clash/proxies?type=all>
## 结语

clash年久失修，不支持VLESS测速。可配合其他如clash.meta,或xarY内核进行测速检测

发现一个很致命的问题，明明能联通得，程序检测不成功，暂未细看

后面的想法是固定住节点名称，按日期排序，目前节点名是不固定的

## 声明

本项目遵循 GNU General Public License v3.0 开源，在此基础上，所有使用本项目提供服务者都必须在网站首页保留指向本项目的链接

本项目仅限个人自己使用，禁止使用本项目进行营利和做其他违法事情，产生的一切后果本项目概不负责
