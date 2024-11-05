# Background Worker
- Background Worker是一个融合服务器的后端,基于.net,C++,Go,python,nodejs的后台任务处理框架。它可以帮助你快速实现后台任务的处理，并提供统一的接口，让你能够方便的将后台任务的处理结果通知到前端。
数据库采用融合数据库有MySQL,MongoDB,Redis等多种数据库，可以满足不同场景的需求。

## 特性
- 统一的接口：提供统一的接口，让你能够方便的将后台任务的处理结果通知到前端。
- 多种语言支持：支持.net,C++,Go,python,nodejs等多种语言，可以满足不同场景的需求。
- 多种数据库支持：数据库采用融合数据库有MySQL,MongoDB,Redis等多种数据库，可以满足不同场景的需求。


**注意：Background Worker项目，目前处于开发阶段，功能还在不断完善中，欢迎大家提出宝贵意见，共同打造一个更好的后端任务处理框架。**

## 快速开始
### 安装
#### 环境准备
- 安装.net core sdk
- 安装C++环境
- 安装Go环境
- 安装Python环境
- 安装Nodejs环境
#### 安装数据库
- 安装MySQL或MongoDB等
#### 安装Background Worker
- 下载Background Worker源码
- 编译源码
- 运行Background Worker
    - 启动命令：dotnet BackgroundWorker.dll
    - 后台任务默认端口：5000
### 配置
#### 接口配置
- 打开appsettings.json文件
- 添加以下配置：
```json
"BackgroundWorker": {
    "Host": "http://localhost:5000",
    "EnableCors": true,
    "EnableAuthorization": false,
    "AuthorizationEndpoint": "/api/token",
    "AuthorizationMethods": "Bearer",
    "AuthorizationHeader": "Authorization",
    "EnableLogging": false,
    "EnableStatusMonitor": false,
    "StatusMonitorEndpoint": "/api/status"
}
```
#### 数据库配置
- 配置MySQL数据库
    - 打开appsettings.json文件
    - 添加以下配置：
    ```json
    "ConnectionStrings": {
        "DefaultConnection": "Server=localhost;Database=backgroundworker;User Id=root;Password=123456;Charset=utf8mb4;"
    }
    ```
- 配置MongoDB数据库
    - 打开appsettings.json文件
    - 添加以下配置：
    ```json
    "MongoDb": {
        "ConnectionString": "mongodb://localhost:27017",
        "DatabaseName": "backgroundworker"
    }
    ```
#### 其他配置
- 打开appsettings.json文件
- 添加以下配置：
```json
"BackgroundWorker": {
    "Host": "http://localhost:5000",
    "EnableCors": true,
    "EnableAuthorization": false,
    "AuthorizationEndpoint": "/api/token",
    "AuthorizationMethods": "Bearer",
    "AuthorizationHeader": "Authorization",
    "EnableLogging": false,
    "EnableStatusMonitor": false,
    "StatusMonitorEndpoint": "/api/status"
}
```
- 启动Background Worker
- 访问后台任务接口：http://localhost:5000/api/backgroundtasks
### 开发

## 贡献者
- [shiyelong](https://github.com/shiyelong)