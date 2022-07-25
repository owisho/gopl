package tmp

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestString(t *testing.T) {
	s := "123456"
	bs := []byte(s)
	fmt.Printf("[%d %d %d %d %d %d]\n", s[0], s[1], s[2], s[3], s[4], s[5])
	fmt.Println(bs)
	bs[0] = 2
	fmt.Println(bs)
	fmt.Println(s)
}

func TestByteArr(t *testing.T) {
	br := [...]byte{1, 2, 3, 4, 5}
	b1 := br[1:4]
	b1[0] = 1
	fmt.Println(b1)
	fmt.Println(br)
}

func TestMemcache(t *testing.T) {
	client := memcache.New("39.106.64.162:11211")
	client.Timeout = 100 * time.Second
	err := client.Ping()
	if err != nil {
		fmt.Printf("ping err = %v\n", err)
		return
	}
	fmt.Println("success")
	f, err := os.Open("/Users/wangyang/Desktop/test.svg")
	if err != nil {
		fmt.Printf("file open error=%v\n", err)
		return
	}
	br, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("file read error=%v\n", err)
		return
	}
	err = client.Set(&memcache.Item{
		Key:        "test",
		Value:      br,
		Flags:      0,
		Expiration: 0,
	})
	if err != nil {
		fmt.Printf("set memcacehe error %v", err)
		return
	}
	item, err := client.Get("test")
	if err != nil {
		fmt.Printf("get err = %v\n", err)
		return
	}
	f1, err := os.Create("/Users/wangyang/Desktop/jietu_0722.png")
	if err != nil {
		fmt.Printf("file create error %v\n", err)
		return
	}
	n, err := f1.Write(item.Value)
	if err != nil {
		fmt.Printf("file write error %v\n", err)
		return
	}
	fmt.Printf("write %d byte to file %s", n, f1.Name())
}

type Employee struct {
	ID            int
	Name, Address string
	DoB           time.Time
	Position      string
	Salary        int
	ManagerID     int
}

func EmployeeByID(id int) Employee {
	return Employee{}
}

func TestEmployeeByID(t *testing.T) {
	fmt.Println(EmployeeByID(1).ID)
	employee := EmployeeByID(1)
	employee.Name = "test"
}
