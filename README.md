# go-uno

> 🚀 The Premier Go Web Scaffold - Your #1 Starting Point for Go Web Projects

go-uno is a meticulously crafted, production-ready web application scaffold built with Go. It embodies best practices and clean architecture to kickstart your next project with speed and confidence.

## 简介

go-uno 是一个基于 Go 语言的 Web 应用脚手架，旨在提供一个快速、可靠的起点来构建生产级别的 Web 服务。它集成了 Gin 框架、Cobra CLI 工具和 Viper 配置管理，并遵循清晰的架构设计，帮助开发者专注于业务逻辑的实现。

## 特性

-   **Gin 框架**: 高性能的 HTTP Web 框架。
-   **Cobra CLI**: 强大的命令行接口工具，方便构建命令行应用。
-   **Viper 配置管理**: 灵活的配置解决方案，支持多种格式和热加载。
-   **清晰的架构**: 采用分层设计，易于理解和扩展。
-   **优雅的关机**: 支持平滑关机，确保请求得到妥善处理。

## 安装与运行

### 前提条件

-   Go 1.18+ 环境

### 克隆项目

```bash
git clone https://github.com/rushairer/go-uno.git
cd go-uno
```

### 构建项目

使用 `make` 命令：

```bash
make
```

### 运行 Web 服务

```bash
./bin/gouno web --config ./config/config.yaml --address 0.0.0.0 --port 8080 --debug
```

默认情况下，Web 服务将在 `http://0.0.0.0:8080` 启动。

## 使用说明

### 命令行参数

-   `--config` 或 `-c`: 指定配置文件路径，默认为 `./config/config.yaml`。
-   `--address` 或 `-a`: 指定监听地址，默认为 `0.0.0.0`。
-   `--port` 或 `-p`: 指定监听端口，默认为 `8080`。
-   `--debug` 或 `-d`: 开启调试模式，默认为 `false`。

### 配置文件

配置文件位于 `./config/config.yaml`，您可以根据需要修改其中的配置项。

```yaml
web_server:
    address: 0.0.0.0
    port: 8080
    debug: false
```

### 路由示例

项目提供了一个简单的 `/test/alive` 路由示例，用于检查服务是否存活。

```bash
curl http://localhost:8080/test/alive
# Expected output: pong
```

## 贡献

欢迎通过以下方式为 go-uno 贡献代码：

1.  Fork 本仓库。
2.  创建您的特性分支 (`git checkout -b feature/AmazingFeature`)。
3.  提交您的更改 (`git commit -m 'Add some AmazingFeature'`)。
4.  推送到分支 (`git push origin feature/AmazingFeature`)。
5.  提交 Pull Request。

## 许可证

本项目采用 MIT 许可证。详情请参阅 [LICENSE](LICENSE) 文件。
