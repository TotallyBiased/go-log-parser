package main

	
import (
    // "bufio"
    "fmt"
    // "io"
    "io/ioutil"
    // "os"
    "strings"
    "regexp"
)
var logReg = regexp.MustCompile(`^(\S+) (\S+) (\S+) \[([\w:/]+\s[+\-]\d{4})\] "(\S+)\s?(\S+)?\s?(\S+)?" (\d{3}|-) (\d+|-)\s?"?([^"]*)"?\s?"?([^"]*)?"?`)
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
  dat, err := ioutil.ReadFile("./example-data.log")
  check(err)
  logs := strings.Split(string(dat), "\n") 
  ipSet := make(map[string][][][]string)
  for _, s := range logs {
    if s != "" {
      matched := logReg.FindAllStringSubmatch(s, -1)
      if _, ok := ipSet[string(matched[0][1])]; !ok {
        var list [][][]string
        result := append(list, matched[0])
        ipSet[string(matched[0][1])] = result
      } else {
        fmt.Println(matched[0][1])
      }
    }
    
  }


  // f, err := os.Open("./example-data.log")
  // check(err)

  // b1 := make([]byte, 5)
  // n1, err := f.Read(b1)
  // check(err)
  // fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

  // o2, err := f.Seek(6, 0)
  // check(err)
  // b2 := make([]byte, 2)
  // n2, err := f.Read(b2)
  // check(err)
  // fmt.Printf("%d bytes @ %d: ", n2, o2)
  // fmt.Printf("%v\n", string(b2[:n2]))

  // o3, err := f.Seek(6, 0)
  // check(err)
  // b3 := make([]byte, 2)
  // n3, err := io.ReadAtLeast(f, b3, 2)
  // check(err)
  // fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

  // _, err = f.Seek(0, 0)
  // check(err)

  // r4 := bufio.NewReader(f)
  // b4, err := r4.Peek(5)
  // check(err)
  // fmt.Printf("5 bytes: %s\n", string(b4))

  // f.Close()
}