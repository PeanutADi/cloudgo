## Golang web app

---
## 概述

设计一个 web 小应用，展示静态文件服务、js 请求支持、模板输出、表单处理、Filter 中间件设计等方面的能力。

## 实现效果

### 1. 运行go文件启动服务器

![启动.png](https://i.loli.net/2019/01/18/5c41cee41a2fa.png)

之后再8080端口查看这个web app

### 2. 静态文件访问

![访问静态文件.png](https://i.loli.net/2019/01/18/5c41cee41a160.png)

### 2. Ajax获取当前时间

上部分为表单提交，下半部分显示了当前时间

![LoginInterface.png](https://i.loli.net/2019/01/18/5c41cee4187a6.png)

### 3. 提交表单并输出

![InputUser.png](https://i.loli.net/2019/01/18/5c41cee3f29ac.png)

点击Login提交表单，然后进入signin界面查看表单

![产生表单.png](https://i.loli.net/2019/01/18/5c41cee41857e.png)

### 4. 对未处理返回501错误

在服务器后端，即cmd中，查看服务器信息

![501.png](https://i.loli.net/2019/01/18/5c41cee3c8812.png)

在web界面的信息

![InDevelop.png](https://i.loli.net/2019/01/18/5c41cee3c8b5f.png)
