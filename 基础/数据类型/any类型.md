1. any 是任意类型，any 是 interface{} 的别名,是go1.18 以后开始加的

        type any interface{}

2. 声明为any 类型的变量，可以赋值任意类型的值

        var a any

        a=1;
        a="hello"
        a=true