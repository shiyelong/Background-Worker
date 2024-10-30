cmake_minimum_required(VERSION 3.17)
project(MyProject)

set(CMAKE_CXX_STANDARD 14)

# 添加 Boost 头文件路径
include_directories("/opt/homebrew/include")

# 创建可执行文件
add_executable(main ./include/includePublic.h)