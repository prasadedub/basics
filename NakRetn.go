package main 
import "fmt"

func main() {
     id, err := randId()
     if err != nil {
	     fmt.Printf("Error = %s\n", err)
}
     fmt.Printf("print ID %d\n", id)
}
func randId() (int, error) {
	id := 10
	return id, nil
}
