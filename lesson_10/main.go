package main 

import(
	"fmt"
	m "github.com/walonCode/math/math"
	"github.com/walonCode/math/hello"
)


//when creating a package we need the following
// initialize go mod init github.com/username/packagename in the root  file
// next make a folder which will be the package name aand make a package name.go file in it
func main(){
	sum := m.Add(2,4)
	fmt.Println(sum)
	hello.Hello("walon")
}