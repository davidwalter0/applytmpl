package main

var Version string = "Version option not set" // from the build ldflag options
var Build string = "Build option not set"     // from the build ldflag options
var Commit string = "Commit option not set"   // from the build ldflag options

// func init() {
// 	for _, arg := range os.Args {
// 		if arg == "version" {
// 			array := strings.Split(os.Args[0], "/")
// 			me := array[len(array)-1]
// 			fmt.Println(me, "Build:", Build, "Commit:", Commit)
// 			os.Exit(0)
// 		}
// 	}
// }
