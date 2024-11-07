cmake_minimum_required(VERSION 3.29)
project(Background_Worker)

set(CMAKE_CXX_STANDARD 14)

include_directories(src/cpp/include)

add_executable(Background_Worker
    src/cpp/include/GameLogic.h
    src/cpp/include/GameServer.h
    src/cpp/include/includePublic.h
    src/cpp/src/GameLogic.cpp
    src/cpp/src/GameServer.cpp
    tests/cpp/GameLogicTests.cpp
    tests/cpp/GameServerTests.cpp)
