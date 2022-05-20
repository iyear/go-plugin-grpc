# 项目名称

一种基于 gRPC 的轻量级应用间双向通信框架

# 项目方案

## 背景

正如项目描述，`Go` 原生 `Plugin` 实现的问题多、限制大，并不适合作为插件机制来使用。在我的 `pure-live` 项目起步时，我对此也做过一番调研。跨平台限制、插件开发生命周期限制、兼容性限制均是不可接受的痛点。最终，我只能以代码耦合的形式暂时实现简单的插件设计。

借此次OSPP为契机，我想实现一个基于网络层的轻量插件框架，他将解决以下问题：

- **无法单独管理不同插件的维护周期，牵一发而动全身。**低耦合的插件设计下，项目扩展性、健壮性、灵活性将大大提高。
- **无法跨主机部署调用。**尽管在小项目中是不必要的，但对于大型项目应当是可选项。
- **无法跨语言维护。** `gRPC` 允许不同语言的开发者共同参与到插件的实现维护中。
- **插件的代码编写失误可能导致宿主出现崩溃。**

## 对已有开源实现的评估

在了解需求之后，本着“不重复造轮子”的原则，先尝试寻找已有的开源实现。

`go-plugin` 是 `hashicorp` 于2016年开始提交的仓库，如今经过了4年多的生产环境验证，已用于许多著名开源项目。

如此优秀的项目，是否可以直接用于这次的 `OSAPP` ？**我认为是否定的**，在阅读其源码以及历史 `commits` 后，以下是我认为 `go-plugin` 存在的一些问题：

- **代码实现历史包袱重。**其支持 `net/rpc` 与 `gRPC` 两种通信底层。`gRPC` 为后期添加，`net/rpc` 如今已基本弃用。由于项目从2016年1月开始维护，2017年六月开始加入 `gRPC` 支持。所以源码中使用了大量的接口为 `net/rpc` 做向下兼容，源码阅读也十分困难，不利于长期维护。

- **只支持本机插件。**`go-plugin` 以 `exec.Command` 的方式启动插件，以 `process id` 检查插件存活。这就完全无法将插件系统在云原生上组织。同时 `pid` 这种强依赖操作系统的检查也一定程度上增加了跨平台难度。

- **日志系统可扩展性差。**由于只支持本机系统，`go-plugin`的日志采用 `stdout/stderr` 方式直接 `copy` 到宿主进程。`stdout/stderr` 的灵活性比网络传输差，难以做云原生迁移。

- **需要开发者大量封装。**`hashicorp/vault` 中需要开发者提前封装接口、中间转换函数，过程冗杂，代码量高，不符合“轻量”目标。

## 实现方案

先对以上问题做出方案回应：

- **采用 `gRPC` 为唯一支持协议。**`net/rpc` 不支持跨语言，兼容性差。在维护力度、性能方面与 `gRPC` 差距也较大。我认为可以直接抛弃协议兼容性，采用 `gRPC` 为唯一支持协议，简化具体实现。另一方面站在巨人的肩膀上，`gRPC` 已自带 `TLS` `Metadata` 等通用组件。
- **插件与宿主完全采用网络通信、反转插件与宿主的 `C/S` 关系。**避免进程启动、`stdout/stderr` 日志汇集等操作系统层面的代码实现；默认实现健康检查、日志汇集；令宿主为 `Server` ，插件为 `Client` ，插件启动方式由上层管理。
- **多种传输类型支持、代码生成辅助开发**。默认支持 `[]byte` 和 `map[string]interface{}` 两种类型。开发者可根据需求任意选择；代码生成增强编译期约束，减少低级错误。

以下为详细方案实现：

### 核心概念

了解核心概念有利于把握整体。方案中的核心概念并不多，且非常容易理解，设计的一切都在保证“轻量”目标。

#### Core

`Core` 为宿主，它是一个项目的核心实现，主要包含项目逻辑，基本不包含业务逻辑。开发者、用户通过 `Core` 调用插件中的函数。`Core` 被设计为 `Server` 。

#### Plugin

Plugin 为插件，它实现了 `Core` 所需的具体业务逻辑。`Plugin` 被设计为 `Client` 。

#### Convention/约定

设计方案倾向于“约定”而非“约束”，如果想和 `go-plugin` 一样做到强约束，开发者必须付出大量时间对约束性做出封装。而对于一个小型、轻量的项目，约束应当只是可选项。`Core` 与 `Plugin` 之间的插件名、版本、函数名、参数类型、结果类型应尽可能先约定，后约束。

当然，对于一个大型项目，“约束”不应当被完全抛弃。下文将说明框架如何使用实现一定的约束性。

#### Interface

`Core` 可以定义多个 `interfce` 。相似但不同于 `Golang` 中的 `interface` ，为了贯彻单一职责，一个 `Plugin` 只允许实现 `Core` 的一个接口，履行一项职责。

#### Func/Handler

在 `Plugin` 执行 `Mount` 前，所有注册的业务逻辑函数被称为 `Func/Handler` 。函数签名为 `func(ctx plugin.Context) (interface{}, error)` 。其接收来自 `Core` 的参数，进行处理后将结果返回给 `Plugin` ，`Plugin` 帮助它向 `Core` 发送结果。

### 选型

- `gRPC Bidirectional Streaming` 为 `Main Stream` ，用于传输除 `Log` 以外的所有消息。
- `gRPC Client-Side Streaming` 为 `Log Stream`，只用于传输 `Plugin Log`。
- `ProtoBuf` 为所有消息的序列化协议。

### 架构图

<img src="arch.png" width="500px"  alt="arch"/>

`Core` 包含以下组件/设置：

- `Plugins Manager` 。记录 `Plugin` 信息与状态，一旦 `Plugin` `Mount` 到 `Core` ，`Plugins Manager` 将对 `Plugin` 拥有全管理权。
- `gRPC Server` 。承载起 `Core` 与所有 `Plugin` 的通信，支持一系列设置
- `Mount `
- `Log Hub`。负责收集 
- `Health Check Cron`。健康检查定时任务，`Core` 将定时检查 `Plugin` 健康状态。
- `Developer Hook`。框架使用者可以对框架内

### `ProtoBuf` 设计

<img src="communicate-proto.png" width="500px"  alt="communicate-proto"/>

在 `Communicate Stream` 中，所有消息均被 `Communicate Msg` 包装，`Type` 字段定义此次消息类型，`Core` `Plugin` 根据消息类型做出特定

<img src="log-proto.png" width="300px"  alt="log-proto"/>



### `Core`-`Plugin` 生命周期

<img src="life-circle.png" width="500px"  alt="life-circle"/>

### 方案详细描述

**注意：框架完美运行的前提为 `Core` `Plugin` 互相信任，即 `Core` 只对 `Plugin` 的 `Mount` `Unmount` 阶段做合法性校验，之后的内容传输将完全信任**

#### Mount

`Plugin` 



















