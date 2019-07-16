package opintent

// import "github.com/Myriad-Dreamin/go-py"
//
// var mModule GoPy.PyModule
// var OpIntent GoPy.PyClass
//
//
// func Jsonize(mObj GoPy.RefPyObject) string {
//     return GoPy.GoString(GoPy.InvokeMemberFunction(mObj, "jsonize"))
// }
//
// func BuildGraph(pDict GoPy.PyDict) GoPy.PyTuple {
//     return GoPy.InvokeMemberFunction(OpIntent, "build_graph", pDict)
// }
//
// func GoBuildGraph(pDict GoPy.PyDict) GoPy.PyTuple {
//     return GoPy.InvokeMemberFunction(OpIntent, "build_graph", pDict)
// }
//
// func init() {
// 	mModule = GoPy.RequireModule("uiputils.op_intents")
// 	OpIntent = GoPy.GetAttr(mModule, "OpIntent")
//
// 	GoPy.RegisterAtExitFunc(func(){
// 		GoPy.DecRef(&mModule)
// 		GoPy.DecRef(&OpIntent)
// 	})
// }
