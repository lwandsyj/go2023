1. 对于底层网络应用开发者而言，几乎所有网络编程都是Socket，因为大部分底层网络的编程都离不开Socket编程。HTTP编程、Web开发、IM通信、视频流传输的底层都是Socket编程

   Socket编程主要是面向OSI模型的第三层和第四层协议，再具体点就是主要针对IP协议、TCP协议和UDP协议，

   + TCP: 流式 Socket 是一种面向连接的Socket

   + UDP: 数据报式Socket是一种无连接的Socket，对应无连接的UDP服务应用。

2. 网络进程：`网络层的“IP地址”可以唯一标识网络中的主机，而传输层的“协议+端口”可以唯一标识主机中的应用程序（进程）`

+ IP 地址

+ 协议

+ 端口号

