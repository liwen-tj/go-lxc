package main
/* 执行代码并获取输出 */
import (
	"log"
	"fmt"
	// "runtime"
	// "os"
	// "bytes"
	"os/exec"
)


// func main() {
//     cmd := exec.Command("python3", "echo.py")
//     out, err := cmd.CombinedOutput()
//     if err != nil {
//         log.Fatalf("cmd.Run() failed with %s\n", err)
//     }
//     fmt.Printf("combined out:\n%s\n", string(out))
// }


// 可以输出到标准屏幕上
// func main() {
// 	cmd := exec.Command("ls", "-lah")
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
// 	err := cmd.Run()
// 	if err != nil {
// 		log.Fatalf("cmd.Run() failed with %s\n", err)
// 	}
// 	fmt.Printf("Done")
// }

// 可以把打印结果放到某个buffer里面之后再打印出来
// func main() {
// 	cmd := exec.Command("ls", "-lah")
//     var stdout, stderr bytes.Buffer
//     cmd.Stdout = &stdout
//     cmd.Stderr = &stderr
//     err := cmd.Run()
//     if err != nil {
//         log.Fatalf("cmd.Run() failed with %s\n", err)
//     }
//     outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
//     fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
// }


// 执行了，但是没有输出
func main() {
	cmd := exec.Command("ls", "-lah")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("Done")
}
