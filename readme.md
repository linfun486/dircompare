# DirCompare

一个用于比较两个目录中相同文件名的工具。

## 功能特点

- 支持比较两个目录中的文件名
- 支持递归搜索子目录
- 支持文件名大小写敏感/不敏感匹配
- 显示重复文件的具体路径

## 使用方法

```bash
go run main.go [-r] [-i] <目录1> <目录2>
```

### 参数说明

- `-r`: 递归搜索子目录（可选）
- `-i`: 忽略文件名大小写（可选）
- `目录1`: 第一个要比较的目录路径
- `目录2`: 第二个要比较的目录路径

### 使用示例

1. 基本比较（仅比较根目录）：
```bash
go run main.go /path/to/dir1 /path/to/dir2
```

2. 递归比较所有子目录：
```bash
go run main.go -r /path/to/dir1 /path/to/dir2
```

3. 忽略大小写比较：
```bash
go run main.go -i /path/to/dir1 /path/to/dir2
```

4. 递归且忽略大小写比较：
```bash
go run main.go -r -i /path/to/dir1 /path/to/dir2
```

## 输出说明

程序会输出在两个目录中发现的相同文件名，并显示它们的完整路径。如果使用了 `-i` 参数，则 "Test.txt" 和 "test.TXT" 会被视为相同文件。