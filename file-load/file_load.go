package fileload


import "github.com/Myriad-Dreamin/go-py"


var FileLoad GoPy.RefPyObject


func LoadJson(pathToJson string) GoPy.PyDict {
	return GoPy.InvokeMemberFunction(FileLoad, "getopintents", GoPy.PyString(pathToJson))
}

func init() {
	FileLoad = GoPy.RequireObject("uiputils.ethtools.loadfile", "FileLoad")

	GoPy.RegisterAtExitFunc(func(){
		GoPy.DecRef(&FileLoad)
	})
}


