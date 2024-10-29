cmake_minimum_required(VERSION 3.17)
project(6_boost_demo)

set(CMAKE_CXX_STANDARD 14)

set(BOOST_ROOT "/opt/homebrew/Cellar/boost/1.86.0_1/include/boost")

set(BOOST_INCLUDE "/opt/homebrew/Cellar/boost/1.86.0_1/include")

include_directories("/opt/homebrew/Cellar/boost/1.86.0_1/include") # 添加头文件搜索路径

include_directories("/opt/homebrew/Cellar/boost/1.86.0_1/include/boost") # 添加头文件搜索路径

link_directories("/opt/homebrew/Cellar/boost/1.86.0_1/lib") # 添加库文件搜索路径

aux_source_directory(. DIR_SRCS)

add_executable(6_boost_demo ${DIR_SRCS})