# Golang

## 26-01-19

### 环境配置

- 安装Golang
- 设置Golang/bin到Path环境变量
- 重启Vscode以读取全局环境变量

### 第一个go程序(helloworld)

- helloworld程序编写
- 运行命令`go run filename.go`
- { 不能放在单独一行

### 格式化输出(output_formatting)

- `%s: 字符串`
- `%d: 整数`
- `%f: 浮点数`
- `%t: 布尔值`
- `%c: 字符`
- `%v: 默认值`

### 斐波那契数列(Fibonacci)

- 并行赋值是先计算右侧再赋值，循环计算需要将a+b独立计算

### 显式转换(explicit_conversion)

- 两个类型名称不完全一致的场景，必须要**显式类型转换**。

## 26-1-21

### 学习slice和map概念(exclude&count)

- `if` `countinue` 的运用

### 学习结构体概念(struct)

- 简单的Struct（数据模型） + Slice（数据集） + Loop（遍历） + Logic（业务判断）代码

### 学习方法的概念(method)

- 使用方法给(struct) 加判断逻辑
- 普通函数： `func IsAdult(u User) bool` -> 这是一个公共工具，谁都能用。  
方法： `func (u User) IsAdult() bool` -> 这是 User 独有的技能。

### 学习使用指针修改数组中真实数据

- 使用指针才能修改数组中的真实数据  
`func (u *User) HappyBirthday() {
	u.Age = u.Age + 1
}`