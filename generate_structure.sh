#!/bin/bash

# 读取 README.md 文件
FILENAME="mk.md"

# 检查 README.md 文件是否存在
if [ ! -f "$FILENAME" ]; then
    echo "$FILENAME 不存在!"
    exit 1
fi

# 创建目录和文件
while IFS= read -r line; do
    # 去掉前导和后缀空格
    trimmed_line=$(echo "$line" | sed 's/^[ \t]*//;s/[ \t]*$//')

    # 检查是否是目录
    if [[ "$trimmed_line" == *// ]]; then
        # 创建目录
        mkdir -p "$trimmed_line"
    elif [[ "$trimmed_line" == *.* ]]; then
        # 创建文件的目录
        dir=$(dirname "$trimmed_line")
        mkdir -p "$dir"  # 确保目录存在
        # 创建文件
        touch "$trimmed_line"
    fi
done < <(grep -v '^$' "$FILENAME") # 跳过空行

echo "文件夹和文件已生成!"