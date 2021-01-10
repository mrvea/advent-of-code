package blank

import (
	"log"
	"os"
)

const template = `
package {{packageName}}
const (
	day           = "{{.Day}}"
	basePath      = "2020/day-" + day + "/"
	inputName     = "input.txt"
	testInputName = "test_" + inputName
)

func SubMain(args ...string){
	fmt.Println("vacation" + {{.Day}} + ".{{.Part}} starting exectution...")
	filePath := basePath + inputName
	if len(args) > 0 && args[0] == "test" {
		fmt.Println("Running test file...")
		filePath = basePath + testInputName
	}
	fmt.Printf("filePath: %s\n", filePath)
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	total := 0
	for s.Scan() {
		_ := s.Text()
	}

	fmt.Println("total: ", total)
}

`

func SubMain(args ...string) {
	l := len(args)
	if l == 0 {
		log.Fatalf("please specify which day you want to make")
	}
	data = struct {
		Day         string
		Part        string
		PackageName string
	}{
		Day: args[0],
	}
	day := args[0]
	name := "main"
	if l > 1 {
		name = args[1]
	}
	fName := "day-" + day
	_, err := os.Stat(fName)

	if os.IsExist(err) {
		log.Fatalf("folder with name %s already exists", fName)
		return
	}
	errDir := os.MkdirAll(fName, 0755)
	if errDir != nil {
		log.Fatal(err)
	}

	makePartFolders()

}
