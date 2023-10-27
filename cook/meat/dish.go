package meat

import "PACKAGESTUDY/logprint"

func MakeDish(name string) {
	msg := "做的荤菜:" + name

	logprint.Info(msg)
}
