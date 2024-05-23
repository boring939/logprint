package logprint

import (
	"fmt"
	"time"
)

// var t = time.Now()

func Errlog(msg interface{}) {
	t := time.Now()
	fmt.Printf("[Err]%s,%s\n", t.Format("2006-01-02 15:04:05"), msg)

}
func Debuglog(msg interface{}) {
	t := time.Now()
	fmt.Printf("[Debug]%s,%s\n", t.Format("2006-01-02 15:04:05"), msg)

}
func Infolog(msg interface{}) {
	t := time.Now()
	fmt.Printf("[info]%s,%s\n", t.Format("2006-01-02 15:04:05"), msg)

}
