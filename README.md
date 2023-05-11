# YUN

[![Build Status](https://travis-ci.org/pipikai/yun.svg?branch=master)](https://travis-ci.org//) [![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

项目介绍

这个项目是一个使用Golang语言实现的分布式文件存储系统，它具有高性能、可拓展、支持文件的秒传、断点续传、分片上传等特性。此外，该存储系统还支持大文件的上传，并通过embed集成了web界面，方便用户使用。

安装说明

这个项目是一个二进制可执行文件，不需要安装任何依赖项。只需要按照以下步骤进行即可：

1. 下载最新版本的可执行文件：

```bash
wget https://github.com/PIPIKAI/yun/releases/latest/download/yun
```

1. 控制台中执行以下命令来启动客户端或存储端：

启动客户端：

```bash
./yun start --f "tracker"
```

启动存储端：

```bash
./yun start --f "storage" 
```

使用示例

以下是一些使用示例，展示如何使用这个项目。

启动客户端：

```
yun start --f "tracker"
```

需要在配置文件中设置Name:Storage

启动存储端：

```
yun start --f "storage"
```

贡献指南

我们欢迎任何形式的贡献，包括但不限于：报告问题、提交bug修复、改进文档和代码等。如果您想做出贡献，请按照以下步骤进行：

1. Fork这个项目，并克隆到本地。
2. 在本地分支中进行修改。
3. 提交Pull Request，描述您的修改内容和用途。

许可证信息

本项目基于MIT许可证开放源代码。

联系方式

如果您对这个项目有任何问题或建议，请通过以下方式联系我们：

- 发送邮件到 z1652091948@outlook.com
- 在GitHub上提交问题

感谢您的使用和贡献！