package main
import (
    "fmt"
    "bufio"
    "os"
)

func main() {
	fmt.Println("Hello, World.")
	//objective can achieved using both reader and scanner types
    // reader := bufio.NewReader(os.Stdin)
    // text, _ := reader.ReadString('\n')
    // fmt.Println(text)
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        fmt.Println(scanner.Text()) // Println will add back the final '\n'
    }
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input:", err)
    }
}