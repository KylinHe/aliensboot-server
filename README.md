目录结构

bin  常用的工程脚本文件
  |- gencode.sh 通过protobuf 协议生成接口代码文件
  |- gendata.sh 把data目录的配置文件生成为代码文件
data 存放业务逻辑配置文件
src 工程代码
  |-公司名
  	  |-工程名
  	  	  |- constant 常量、代码生成的常量代码
  	  	  |- data 自动生成的配置数据模型代码
  	  	  |- dispatch 调用其他服务的接口
  	  	  |- module 模块目录
  	  	  |- protocol 协议文件
  	  	  |- public 服务启动入口文件
templates gencode脚本会根据代码模板生成对应的代码文件
project.yml 工程相关描述文件用于一些工程相关参数读取

#游戏服务端框架