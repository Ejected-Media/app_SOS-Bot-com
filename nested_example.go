// Golang program to illustrate
// the nested structure
package main
 
import "fmt"
 
// Creating structure
type Student struct {
    name   string
    branch string
    year   int
}
 
// Creating nested structure
type Teacher struct {
    name    string
    subject string
    exp     int
    details Student
}
 
func main() {
 
    // Initializing the fields
    // of the structure
    result := Teacher{
        name:    "Suman",
        subject: "Java",
        exp:     5,
        details: Student{"Bongo", "CSE", 2},
    }
 
    // Display the values
    fmt.Println("Details of the Teacher")
    fmt.Println("Teacher's name: ", result.name)
    fmt.Println("Subject: ", result.subject)
    fmt.Println("Experience: ", result.exp)
 
    fmt.Println("\nDetails of Student")
    fmt.Println("Student's name: ", result.details.name)
    fmt.Println("Student's branch name: ", result.details.branch)
    fmt.Println("Year: ", result.details.year)
}


 

 
type Address struct {
    Street     string
    City       string
    State      string
    PostalCode string
}
 
type Person struct {
    FirstName string
    LastName  string
    Age       int
    Address   Address
}
 
func main2() {
    p := Person{
        FirstName: "John",
        LastName:  "Doe",
        Age:       30,
        Address: Address{
            Street:     "123 Main St",
            City:       "Anytown",
            State:      "CA",
            PostalCode: "12345",
        },
    }
 
    fmt.Println(p.FirstName, p.LastName)
    fmt.Println("Age:", p.Age)
    fmt.Println("Address:")
    fmt.Println("Street:", p.Address.Street)
    fmt.Println("City:", p.Address.City)
    fmt.Println("State:", p.Address.State)
    fmt.Println("Postal Code:", p.Address.PostalCode)
}