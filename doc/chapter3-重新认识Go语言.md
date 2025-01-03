# Go语言的特点
## 3-1为什么要用Go语言
### 与C/C++语言对比
1. C语言不是面向对象
2. 直接编译为机器吗，不需要执行环境
3. 一次编码只能适用一种平台
4. 自己处理GC问题  

### 与Java语言对比
1. 编译为字节码
2. 需要JVM执行环境
3. build onece，run everywhere
4. 有虚拟化损失 

### 与JS语言对比
1. 不需要编译，直接解释执行
2. 需要执行环境
3. 有虚拟化损失 

### Go语言特点
1. 直接编译为二进制，没有虚拟化损失
2. 自带运行环境，无需处理GC问题
3. 一次编码可以适用多种平台 
4. 超强的并发支持能力与并发易用性

### 优秀Go语言项目
docker、kubernetes、etcd、codis、BeeGo、falcon、bilibili、TiDB 

### 总结
1. Go综合了多种语言优势
2. Go是一种天生支持高性能并发的语言
3. Go在工业界有广泛的应用 

## 3-2何为Runtime？

### 很多语言都有Runtime
1. Runtime就是程序的运行环境
2. Java：Java虚拟机
3. JavaScript：浏览器内核  

### Go的Runtime特点
1. 没有虚拟机的概念
2. Runtime作为程序的一部分打包进二进制产物（Go+Runtime）
3. Runtime随用户程序一起运行
4. Runtime与用户程序没有明显界限，直接通过函数调用 

### Go的Runtime能力
1. 内存管理能力
2. 垃圾回收能力（GC）
3. 超强的并发能力（协程调度）

### Go的Runtime其他特点
1. Runtime有一定的屏蔽系统调用能力
2. 一些go的关键字其实是Runtime下的函数 

### 总结
1. Go的Runtime负责内存管理、垃圾回收、协程调度
2. Go的Runtime被编译为用户程序的一部分，一起运行

## 3-3Go程序是如何编译的？
go build -n:输出编译过程
### 词法分析
1. 将源代码翻译成Token
2. Token是代码中的最小语义结构

### 句法分析
1. Token序列经过处理，变成语法树

### 语义分析
1. 类型检查
2. 类型推断
3. 查看类型是否匹配
4. 函数调用内联
5. 逃逸分析（Go的变量是放到堆、还是栈，去选择分配）

### 中间码生成（SSA）
1. 为了处理不同平台的差异，先生成中间代码（SSA）
2. 查看从代码到SSA中间码的整个过程
* $env:GOSSAFUNC="main"
* export GOSSAFUNC=main

### 机器码生成
1. 先生成Plan9汇编代码
* go build -gcflags -S main.go
2. 最后编译为机器码
3. 输出的机器码为.a文件

### 链接
1. 将各个包进行链接，包括runtime（将.a文件生成exe文件）

### 总结
1. 编译前端
* 词法分析
* 句法分析
* 语义分析
2. 编译后端
* 中间码生成
* 代码优化
* 机器码生成
3. 链接

## 3-4Go程序是如何运行的？
### Go程序的入口？
* main方法？
* runtime/rt0_XXX.s(汇编文件)
* 读取命令行参数
1. 复制参数数量argc和参数值argv到栈上
* 初始化g0执行栈
1. g0是为了调度协程而产生的协程
2. g0是每个Go程序的第一个协程
* 运行时检测
1. 检查各种类型长度
2. 检查指针操作
3. 检查结构体字段的偏移量
4. 检查atomic原子操作
5. 检查CAS操作
6. 检查栈大小是否是2的幂次
* 参数初始化 runtime.args
1. 对命令行中的参数进行处理
2. 参数数量赋值给argc int32
3. 参数值赋值给argv **byte
* 调度器初始化runtime.schedinit
1. 全局栈空间内存分配
2. 加载命令行参数到os.Args
3. 堆内存空间初始化
4. 加载操作系统环境变量
5. 初始化当前系统线程
6. 垃圾回收器的参数初始化
7. 算法初始化（map、hash）
8. 设置process数量
* 创建主协程
1. 创建一个新协程，执行runtime.main
2. 放入调度器等待调度
* 初始化M
1. 初始化一个M，用来调度主协程
* 主协程执行主函数
1. 执行runtime包中的init方法
2. 启动GC垃圾收集器
3. 执行用户包依赖的init方法
4. 执行用户主函数main.main()

### 总结
* Go启动时经历了检查，各种初始化、初始化协程调度的过程
* main.main()方法也是在协程中运行的

### 问题
* 调度器是什么？
* 为什么要初始化M？
* 为什么不是直接执行main.main，而是将它放入调度器？

### 体会
* Go程序的启动过程像不像一个虚拟机，或者框架？

### 源码怎么看？
1. rt0_linux_amd64.s文件 
* 调用方法_rt0_amd64
2. asm_amd64.s文件

## 3-5Go是面向对象的吗？
### Yes and no
1. Go是允许OO的编程风格
2. Go的Struct可以看作其他语言的Class
3. Go缺乏其它语言的继承结构

### Go的“类”
1. 其他语言中，往往使用class表示一类数据
2. class的每个实例称做“对象”
3. Go是用struct表示一类数据
4. struct的每个实例并不是“对象”，而是此类型的“值”
5. struct也可以定义方法

### Go的继承
1. Go并没有继承关系
2. 所谓Go的继承只是组合
2. 组合中的匿名字段，通过语法糖达成了类似继承的效果

### Go的接口
1. 接口可以定义Go中的一组行为相似的struct
2. struct并不显示实现接口，而是隐式实现

### 总结
1. Go没有对象、没有类、没有继承
2. Go通过组合匿名字段来达到类似继承的效果
3. 通过以上手段去掉了面向对象中复杂而冗余的部分
4. 保留了基本的面向对象特性

## 3-6企业级Go项目包管理方法
### Go包管理困境
1. 没有统一的包管理方式
2. 包之间的依赖关系难以维护
3. 如果同时需要一个包的不同版本，非常麻烦

### 尝试解决
1. 尝试使用godep、govendor、glide等解决
2. 为彻底解决GOPATH存在的问题
3. 使用起来麻烦

### Go Modules
1. 本质上，一个Go包就是一个项目的源码
2. gomod的作用：将Go包和Git项目关联起来
3. GO包的版本就是git项目的Tag
4. gomod就是解决 “需要哪个git项目的什么版本”

### Github无法访问怎么办
go env -w GOPROXY=https://goproxy.cn,direct

### 想用本地文件替代怎么办？
1. go.mod文件追加：replace github.com/Jeffial/tunny => xxx/xxx
2. go vender缓存到本地
* go mod vendor//不是之前的go vendor
* go build -mod vendor

### 如何创建Go Module
1. 删除本地go.mod
* go mod init github.com/liulei18/goliukong
2. 推送至代码仓库
3. 增加新版本时，在仓库打新Tag

### 总结
1. Go Modules将每个包视为一个git项目
2. gomod就是解决“需要哪个git项目的什么版本“
3. 无法连接远程仓库时，使用重定向或者mod vender方案

 