package vegetable

import "PACKAGESTUDY/logprint"

func MakeDish(name string) {
	msg := "做的素菜：" + name

	logprint.Info(msg)
}
