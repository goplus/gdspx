//go:build darwin && (arm64 || amd64) && cgo && ios

/*
iOS适配器 - 用于解决Go在iOS上运行的各种问题

包含的主要功能：
1. 信号处理 - 防止Go运行时的信号处理导致应用崩溃
2. 线程管理 - 配置合适的线程策略避免线程耗尽
3. 内存管理 - 设置GC参数优化内存使用
4. 初始化顺序 - 确保在Go运行时启动前执行必要的C代码

使用方法：
1. 将此文件及其关联的ios_adapter_complete.h和ios_adapter_complete.c复制到你的Go项目中
2. 确保编译时包含iOS构建标签
3. 在你的应用程序初始化时导入此包

作者: Codeium
日期: 2025-02-27
*/

package main

/*
#include <stdlib.h>
#include <signal.h>
#include <pthread.h>
#include "ios_adapter_complete.h" // 包含完整的iOS适配函数
*/
import "C"

import (
	"os"
	"os/signal"
	"runtime"
	"runtime/debug"
	"syscall"
)

// init函数，在Go运行时启动时自动执行
func init() {
	// 在此处初始化iOS适配
}

// doinit函数由main调用，用于在主协程中配置iOS环境
func doinit() {
	configureGoRuntime()
	disableGoSignals()
	redirectStderrToNull()
}

// 配置Go运行时以优化在iOS上的运行
func configureGoRuntime() {
	// 设置最大处理器数为2，避免创建太多OS线程
	runtime.GOMAXPROCS(2)
	// 设置较小的GC目标百分比以减少内存占用
	debug.SetGCPercent(10)
	// 强制执行一次GC
	runtime.GC()
}

// 禁用所有Go的信号处理
func disableGoSignals() {
	// 屏蔽所有信号以防止应用崩溃
	signal.Ignore(
		syscall.SIGABRT,
		syscall.SIGBUS,
		syscall.SIGFPE,
		syscall.SIGILL,
		syscall.SIGPIPE,
		syscall.SIGSEGV,
		syscall.SIGSYS,
		syscall.SIGTRAP,
	)
}

// 重定向标准错误输出到/dev/null
func redirectStderrToNull() {
	// 打开/dev/null文件
	devNull, err := os.OpenFile("/dev/null", os.O_WRONLY, 0)
	if err != nil {
		return
	}
	// 将标准错误重定向到/dev/null
	err = syscall.Dup2(int(devNull.Fd()), int(os.Stderr.Fd()))
	if err != nil {
		devNull.Close()
		return
	}
}
