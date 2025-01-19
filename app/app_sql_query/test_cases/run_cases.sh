#!/bin/bash

# 获取当前目录的绝对路径
CURRENT_DIR=$(pwd)

# 设置 CHASSIS_HOME 环境变量
export CHASSIS_HOME="${CURRENT_DIR}"

# 运行测试
go test -v ./testcases