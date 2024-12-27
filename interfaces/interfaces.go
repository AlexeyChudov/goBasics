package interfaces

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

// import "fmt" // пакет используется для проверки ответа, не удаляйте его
type battery struct {
	energy   uint
	capacity uint
}

func (b battery) String() string {
	var res string
	res += "["
	fmt.Print(b.energy)
	for i := uint(0); i < b.capacity; i++ {
		if i < b.capacity-b.energy {
			res += "0"
		} else {
			res += "1"
		}
	}
	res += "]"
	fmt.Println(res)
	return res
}
func interface_example() {
	// batteryForTest - не забывайте об имени

	var str string = "1000001011"
	var batteryForTest battery
	batteryForTest.capacity = uint(10)
	fmt.Scan(str)
	for _, ch := range str {
		if ch == '1' {
			fmt.Println("хуй")
			batteryForTest.energy++
		}
	}
	fmt.Println(batteryForTest.energy)
	fmt.Println(batteryForTest)

}

func WriterInterfaces() {
	message := "message bez napryaga"
	buffer := &bytes.Buffer{}
	file, _ := os.OpenFile("msg.txt", os.O_CREATE|os.O_RDWR, 0644)
	WriteEncodedMessage(buffer, message)
	WriteEncodedMessage(file, message)
	fmt.Println(buffer)
}

func WriteEncodedMessage(writer io.Writer, message string) {
	encodedMessage := base64.StdEncoding.EncodeToString([]byte(message))
	writer.Write([]byte(encodedMessage))
}

func WebServer() {

}
