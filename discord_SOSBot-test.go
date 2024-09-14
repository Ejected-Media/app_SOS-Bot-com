package main
 
import "fmt"

// Creating structure
type discordRoom_name struct {
    name   string
    branch string
    areaID   int
}
 
// Creating nested structure
type discordRoom_area struct {
    name    string
    subject string
    roomID     int
    details discordRoom_name
}

func main() {
	
discordRoom_list := discordRoom_area{
        name: "Welcome Zone",
        subject:  "Doe",
        roomID:       30,
        details: discordRoom_name{
            name:     "Welcome Board",
            branch:       "Anytown",
            areaID: 12345,
        },
    }
    
    fmt.Println(discordRoom_list)
    fmt.Println("roomID:", discordRoom_list.roomID)
    
    }