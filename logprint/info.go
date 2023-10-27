// 1.最好是小写字母，并且最好不要带下划线，见名知义
package logprint

import (
	"fmt"
	"time"
)

func Info(msg interface{}) {
	//接受任何信息
	//标记当前时间
	t := time.Now()
	fmt.Printf("[info]%s: %s \n", t.Format("2006-01-02 15:04:05.000"), msg)

}
