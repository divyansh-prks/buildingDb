
//MY try 
// package main

// import (
//     "bufio"
//     "fmt"
//     "log"
//     "os"
// )

// func main() {
//     fmt.Println("input text:")
//     reader := bufio.NewReader(os.Stdin)
//     line, err := reader.ReadString('\n')

// 	if (line == `INSERT users 1 John`) {
// 		fmt.Printf("Logged into insert Command")

// 		//db changes insert into the cloud storage or we will store into disk or fake array db 


// 		//prints the result / output 
// 	}else {
// 		main()
// 	}
//     if err != nil {
//         log.Fatal(err)
//     }
//     fmt.Printf("read line: %s-\n", line)
// }


//changes 

package main 
import (

	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Table struct {
	Name strings
	Column
	Rows map(int ,0,3) 
}

func main(){
	fmt.Println("Input Text :")

	reader := bufio.NewReader(os.Stdin)

	for {
		line , err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		line = strings.TrimSpace(line)

		if line == "INSERT users 1 John" {
			fmt.Println("Logged into insert Command")

			//Place holder : Handle insert logic (store into array /db /file )

			fmt.Println("Inserting into the database")
		}else if line == "EXIT"{
			fmt.Println("Exiting Program")
			break
		} else {
			fmt.Printf("Command not recognized :%s\n" , line )
		}
	}
}