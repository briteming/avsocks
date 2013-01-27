#Notice
##newsocks已改名avsocks，现在项目在 https://github.com/avplayer/avsocks ，使用C++改写，效率更高，bug更少

#README

##起源
前几天受到V2EX的[clowwindy](http://www.v2ex.com/t/32777)启发，想用GO实现一个这样的加密代理，就当是练手。

##优点
- 。。。。（你懂的）
- 效率（？？还没有测试，不过据说go的效率还行）
- 保密性：使用了AES加密，默认为128位

##缺点
- 需要在客户端和服务器端都部署
- 需要有国外服务器才能。。。。（你懂的）

##安装
windows客户端可以和linux服务器端通信

##配置(config.ini)

	[main]
	#这里可以写client和server，本地填写client，服务器填写server
	mode = client 
	
	[server]
	#服务器监听的地址，在mode为server时起作用
	listen = 0.0.0.0:1082
	
	[client]
	#服务器的地址，填写运行server mode的服务器ip和端口
	server = 199.193.249.182:1082
	#本地socks5代理的ip和端口
	listen = 127.0.0.1:1080
	#开启http代理,不开启就留空
	http_proxy = 127.0.0.1:808
	
	[encrypto]
	#用于加密的客户端密钥和服务器密钥
	#客户端的client-key和服务器的client-key必须相同
	#客户端的server-key和服务器的server-key也必须相同
	client-key = 1234567890qwerty 
	server-key = poiuytrewqasdfgh
	#最后要有空行
	

###windows用户
可以直接在Download页面下载已编译版本，包含newsocks.exe,按照本页面给出的配置格式,保存为config.ini,放在与主程序同一目录


###linux/mac用户
首先需要有**Google go**环境

	git clone http://github.com/hyqhyq3/newsocks
	cd newsocks  
	go build
	;;;编辑config.ini
	nohup ./newsocks &	

##重要提示
配置文件最后一行必须要是空行，否则最后一行不会被解析
