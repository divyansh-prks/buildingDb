
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

type Column struct {
	Name string 
	Datatype string 
}

type Table struct {
    Name    string
    Columns []Column
    Rows    map[int][]interface{}  // A map where key is row ID, and value is a slice of column values
}


var usersTable = Table{
	Name : "users", 
	Columns : []Column{
		{Name : "id" , Datatype : "int"}, 
		{Name : "name" , Datatype : "string"},
	}, 
	Rows : make(map[int][]interface{}),
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

			newRow := []interface{}{1 , "John"}
			usersTable.Rows[1] = newRow


			//Place holder : Handle insert logic (store into array /db /file )

			fmt.Println("Inserting into the database")

			fmt.Printf("New user added : Id = %d , Name = %s \n" , 1 , "John")
			//inserting into the persistent storage 
		}else if line == "EXIT"{
			fmt.Println("Exiting Program")
			break
		} else {
			fmt.Printf("Command not recognized :%s\n" , line )
		}
	}
}