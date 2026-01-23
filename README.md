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

### 学习使用指针修改数组中真实数据(pointer)

- 使用指针才能修改数组中的真实数据  
`func (u *User) HappyBirthday() {
	u.Age = u.Age + 1
}`

### 一次基础语法测试(junior_test)

- 使用切片，指针，循环，判断逻辑，方法函数进行能力测试

### 学习接口（interface）的使用方法和原理

- 初步理解interface的概念，明天需要复习

## 26-01-22

### 复习interface和map(map&interface_payment)

- 利用interface和map编写一段简单的代码

### 学习错误处理机制(error_bank_system)

- 结合之前学习的知识搭建一个电子银行ATM程序
- for循环外可以添加lable标签方便跳出

### 学习文件读写操作

- 升级上一个电子银行程序拥有存储数据和读取数据的功能。
- 数据的序列化和反序列化  
- `encoding/json`
- `jsonData, err := json.MarshalIndent(data, "", "    ")`  
- `err = json.Unmarshal(fileData, &data)`  
- 文件写入和文件读取
- `0644`为权限等级  
- 第一位 0：特殊位。在绝大多数 Go 开发场景中，第一位固定为 0。它代表特殊权限（如 SetUID），普通开发者几乎用不到。
- 后三位 644：核心权限。这三位分别代表三个不同的用户身份 
1. (6)：文件所有者 (Owner) —— 就是运行 Go 程序的你。
2. (4)：同组用户 (Group) —— 和你在同一个用户组的人。
3. (4)：其他人 (Others) —— 系统里的任何陌生人。  
- 使用os读写文件
- `os.WriteFile(DataFileName, jsonData, 0644)`  
- `os.ReadFile(DataFileName)`

## 26-01-23

### 学习switch，分包(switch_defer_packages)

- 使用分包，switch重构电子银行程序 
- 构建go项目`go mod init <项目名>