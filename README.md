# go-weather

使用Cobra搭建的地区天气信息查询工具。

## 背景
`go-weather`是在学习Cobra基础库的过程中开发的小工具，本项目使用高德提供的天气查询api进行开发。

## 使用说明
获取实时天气信息：
```sh
go run main.go weather realtime -c="广州"
```
获取天气预报信息：
```sh
go run main.go weather forecast -c="广州"
```
