package main

import (
	"fmt"
	"io/ioutil"
)

func main()  {
	const filename="abc.txt"
	if contents,err := ioutil.ReadFile(filename) ; err != nil {
		fmt.Printf("%s \n",err)
	} else{
		fmt.Printf("%s \n",contents)

	}

}
