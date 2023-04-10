1. io.Copy(w Writer,r Reader): 把读拷贝到写

   有时候我们可能需要先读数据，然后在把数据输出。io.Copy 省去了很多reader 到writer 的过程，而且根据测试，使用Copy 很快而且很省内存。尤其是在大文件的情况下，reader 和 writer 可能会导致内存溢出

2. 示例：

   ![image](../../assets/292.jpg)