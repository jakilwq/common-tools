package common_tools

//my common tools
//20210806
import (
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func GetNowTimeStr() string {
	return time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
}

func GetHostName() string {
	name, _ := os.Hostname()
	return name
}

func GetLocalIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		log.Printf("get local addr err:%v", err)
		return ""
	} else {
		localIp := strings.Split(conn.LocalAddr().String(), ":")[0]
		conn.Close()
		return localIp
	}
}

//func main() {
//	fmt.Println(GetNowTimeStr())
//	fmt.Println(GetHostName())
//	fmt.Println(GetLocalIp())
//}
