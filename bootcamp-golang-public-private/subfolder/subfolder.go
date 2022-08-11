package subfolder

import "fmt"

type privateStruct struct {
	Name string
}

type PublicStruct struct {
	privateProperty string
	PublicProperty  string
}

func IniPublicFunction() {
	fmt.Println("ini public function")
}

func iniPrivateFunction() {
	fmt.Println("ini private function")
}
