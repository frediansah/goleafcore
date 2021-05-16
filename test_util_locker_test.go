package goleafcore

import (
	"testing"
	"time"

	"github.com/frediansah/goleafcore/glinit"
	"github.com/frediansah/goleafcore/glutil"
	"github.com/sirupsen/logrus"
)

var value int = 0
var value2 int = 10

func TestLocker(t *testing.T) {
	glinit.InitLog()
	db := glinit.InitDb()
	defer db.Close()

	n := 4000
	valueFinal := value + n
	valueFinal2 := value2 + n
	logrus.Debug("Value intiial : ", value)

	for i := 0; i < n; i++ {
		go testLockerInc()
		go testLockerInc2()
	}

	time.Sleep(1 * time.Second)
	logrus.Debug("Value final : ", value, " is it = ", valueFinal)
	logrus.Debug("Value final2 : ", value2, " is it = ", valueFinal2)

}

func testLockerInc() {
	glutil.Lock.Get("val").Lock()
	defer glutil.Lock.Get("val").Unlock()
	value = value + 1
}

func testLockerInc2() {
	glutil.Lock.Get("val2").Lock()
	defer glutil.Lock.Get("val2").Unlock()
	value2 = value2 + 1
}
