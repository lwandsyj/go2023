1. recover是一个内建的函数，可以让进入令人恐慌的流程中的goroutine恢复过来。如果当前的goroutine陷入恐慌，调用recover可以捕获到panic的输入值，并且恢复正常的执行

   >`recover 仅用于defer 函数中有效`

   ```go

   func test(){
      defer func(){
        // if 可以有且仅有一条初始化子语句
        if r:=recover();r!=nil{
            log.Printf("捕获到的异常：%v",r)
        }
      }
   }