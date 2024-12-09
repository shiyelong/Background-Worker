# 设置 CMake 的最低版本要求
cmake_minimum_required(VERSION 3.10)

# 设置项目名称和语言
project(GameServerProject VERSION 1.0 LANGUAGES CXX)

# 设置 C++ 标准
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED True)

# 添加包含目录
include_directories(${CMAKE_SOURCE_DIR}/include)

# 添加库目录
link_directories(${CMAKE_SOURCE_DIR}/lib)

# 添加 Boost 库路径（假设 Boost 安装在 /usr/local/lib 和 /usr/local/include）
set(BOOST_ROOT "/usr/local")
set(BOOST_LIBRARYDIR "/usr/local/lib")
set(BOOST_INCLUDEDIR "/usr/local/include")

# 查找 Boost 库
find_package(Boost REQUIRED COMPONENTS system filesystem)
# 老崔记得看下面注释,文件跑通他!!! 给你的例子
# 添加 Unreal Engine 5 的路径（假设 UE5 安装在 /path/to/UnrealEngine）
# set(UE5_ROOT "/path/to/UnrealEngine")
# include_directories(${UE5_ROOT}/Engine/Source/Runtime/Core/Public)
# include_directories(${UE5_ROOT}/Engine/Source/Runtime/Core/Public/Containers)
# include_directories(${UE5_ROOT}/Engine/Source/Runtime/Core/Public/HAL)
# include_directories(${UE5_ROOT}/Engine/Source/Runtime/Core/Public/Misc)
# include_directories(${UE5_ROOT}/Engine/Source/Runtime/Core/Public/Templates)
# include_directories(${UE5_ROOT}/Engine/Source/Runtime/Core/Public/UObject)
# include_directories(${UE5_ROOT}/Engine/Source/Runtime/Core/Public/Serialization)
# include_directories(${UE5_ROOT}/Engine/Source/Runtime/Core/Public/Math)
# include_directories(${UE5_ROOT}/Engine/Source/Runtime/Core/Public/Modules)
# include_directories(${UE5_ROOT}/Engine/Source/Runtime/Core/Public/Logging)
# include_directories(${UE5_ROOT}/Engine/Source/Runtime/Core/Public/Internationalization)
# include_directories(${UE5_ROOT}/Engine/Source/Runtime/Core/Public/Misc)
# include_directories(${UE5_ROOT}/Engine/Source/Runtime/Core/Public/Serialization)
# include_directories(${UE5_ROOT}/Engine/Source/Runtime/Core/Public/Templates)
# include_directories(${UE5_ROOT}/Engine/Source/Runtime/Core/Public/UObject)
# include_directories(${UE5_ROOT}/Engine/Source/Runtime/Core/Public/Serialization)
# include_directories(${UE5_ROOT}/Engine/Source/Runtime/Core/Public/Math)
# include_directories(${UE5_ROOT}/Engine/Source/Runtime/Core/Public/Modules)
# include_directories(${UE5_ROOT}/Engine/Source/Runtime/Core/Public/Logging)
# 添加源文件
file(GLOB_RECURSE SOURCES "src/*.cpp" "src/*.h")

# 添加可执行文件
add_executable(GameServer ${SOURCES})

# 链接 Boost 库
target_link_libraries(GameServer PRIVATE Boost::system Boost::filesystem)

# 设置输出目录
set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/bin)
set(CMAKE_LIBRARY_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/lib)
set(CMAKE_ARCHIVE_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/lib)

# 添加子目录（如果项目有子目录）
# add_subdirectory(subdir)

# 安装目标（可选）
# install(TARGETS GameServer DESTINATION bin)