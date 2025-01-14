// Package main
// @Author: linfuchuan
// @Date: 2025-01-14 01:04:51
// @Description: 比较两个文件夹是否存在相同文件名称

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type item struct {
	Filename string
	Path1    string
	Path2    string
}

// 检查是否为系统隐藏文件
func isHiddenFile(name string) bool {
	// Mac系统隐藏文件
	if name == ".DS_Store" || strings.HasPrefix(name, "._") {
		return true
	}
	// Windows系统文件
	if strings.HasPrefix(name, "desktop.ini") || strings.HasPrefix(name, "thumbs.db") || strings.HasPrefix(name, "$") {
		return true
	}
	// 通用隐藏文件（以.开头的文件）
	// if strings.HasPrefix(name, ".") {
	// 	return true
	// }
	return false
}

func main() {
	// 定义命令行参数
	recursive := flag.Bool("r", false, "是否递归搜索子目录")
	ignoreCase := flag.Bool("i", false, "是否忽略文件名大小写")
	ignoreExt := flag.Bool("e", false, "是否忽略文件扩展名")
	flag.Parse()

	// 检查剩余的参数（两个目录路径）
	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("使用方法: go run main.go [-r] [-i] [-e] <目录1> <目录2>")
		os.Exit(1)
	}

	dir1, err := filepath.Abs(args[0])
	if err != nil {
		fmt.Printf("获取目录1绝对路径失败: %v\n", err)
		os.Exit(1)
	}

	dir2, err := filepath.Abs(args[1])
	if err != nil {
		fmt.Printf("获取目录2绝对路径失败: %v\n", err)
		os.Exit(1)
	}

	// 创建映射来存储文件名和路径的切片
	files1 := make(map[string][]string)
	files2 := make(map[string][]string)

	// 遍历第一个目录
	err = filepath.Walk(dir1, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 如果不是递归模式且不在根目录，则跳过
		if !*recursive {
			relPath, err := filepath.Rel(dir1, path)
			if err != nil {
				return err
			}
			// 如果相对路径包含路径分隔符，说明不在根目录
			if filepath.Dir(relPath) != "." {
				if info.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
		}

		if !info.IsDir() {
			fileName := info.Name()
			// 跳过系统隐藏文件
			if isHiddenFile(fileName) {
				return nil
			}
			if *ignoreCase {
				fileName = strings.ToLower(fileName)
			}
			if *ignoreExt {
				fileName = strings.TrimSuffix(fileName, filepath.Ext(fileName))
			}
			files1[fileName] = append(files1[fileName], path)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("遍历目录1出错: %v\n", err)
		os.Exit(1)
	}

	// 遍历第二个目录
	err = filepath.Walk(dir2, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 如果不是递归模式且不在根目录，则跳过
		if !*recursive {
			relPath, err := filepath.Rel(dir2, path)
			if err != nil {
				return err
			}
			// 如果相对路径包含路径分隔符，说明不在根目录
			if filepath.Dir(relPath) != "." {
				if info.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
		}

		if !info.IsDir() {
			fileName := info.Name()
			// 跳过系统隐藏文件
			if isHiddenFile(fileName) {
				return nil
			}
			if *ignoreCase {
				fileName = strings.ToLower(fileName)
			}
			if *ignoreExt {
				fileName = strings.TrimSuffix(fileName, filepath.Ext(fileName))
			}
			files2[fileName] = append(files2[fileName], path)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("遍历目录2出错: %v\n", err)
		os.Exit(1)
	}

	// 比较并输出相同的文件
	found := false
	var items []item

	for name, paths1 := range files1 {
		if paths2, exists := files2[name]; exists {
			for _, p1 := range paths1 {
				for _, p2 := range paths2 {
					items = append(items, item{
						Filename: name,
						Path1:    p1,
						Path2:    p2,
					})
				}
			}
			found = true
		}
	}

	if !found {
		fmt.Println("未找到相同文件名的文件")
		return
	}

	fmt.Println("找到以下相同文件名的文件：")
	for index, it := range items {
		fmt.Printf("序号%d: 文件名:[%s]  %s <--> %s\n", index, it.Filename, it.Path1, it.Path2)
	}
}
