package main
import "os"
import "fmt"

func main(){
	// f,err:=os.Open("Hello.txt")
	// if err!=nil{
	// 	panic(err)
	// }
	// fileinfo, err:=f.Stat()
	// 	if err!=nil{
	// 		panic(err)
	// 	}
	// 	fmt.Println(fileinfo.Name())
	// 	fmt.Println(fileinfo.IsDir())
	// 	fmt.Println(fileinfo.Size())

	// Read file

	f,err:=os.ReadFile("Hello.txt")
	if err!=nil{
		panic(err)
	}
	defer f.Close()
	fmt.Println(string(f))


	
}
