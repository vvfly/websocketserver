package schedule

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"runtime/pprof"
	"time"

	. "github.com/luckyweiwei/base/logger"
	"github.com/luckyweiwei/base/utils"
)

func GetProgName() string {
	fullPath, _ := exec.LookPath(os.Args[0])
	fname := filepath.Base(fullPath)

	return fname
}

func StartHeapCheck() {
	Log.Debug("enter...")

	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)

	Log.Debugf("MemStats=%v", utils.FormatStruct(&ms))

	heapInuse := ms.HeapInuse
	if heapInuse > 5*1024*1024*1024 { // 5G
		// 生成堆内存报告
		progname := GetProgName()
		dstr := time.Now().Format(utils.TIME_FORMAT_COMPACT)

		f, err := os.OpenFile("./pprof/"+progname+"_"+dstr+".prof", os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			Log.Errorf("err = %v", err)
			return
		}
		defer f.Close()

		err = pprof.WriteHeapProfile(f)
		if err != nil {
			Log.Errorf("err = %v", err)
			return
		}
	}

}
