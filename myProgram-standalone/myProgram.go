package main 

import (
	"ethos/syscall"
	"ethos/ethos"
        "ethos/efmt"
	"log"
)
import "math"

func main () {
	me := syscall.GetUser()
	path := "/user/" + me + "/myDir/"
	fd, status := ethos.OpenDirectoryPath(path)
	if status != syscall.StatusOk {
		log.Fatalf ("Error opening %v: %v\n", path, status)
	}


	data    := MyType { 1, 4, 9, 16 }

	data.Write(fd)
	data.WriteVar(path +"foobar")
        efmt.Println("val:",data.x1,data.y1,data.x2,data.y2)
        data.x1 = math.Sqrt(data.x1)
        data.y1 = math.Sqrt(data.y1)
        data.x2 = math.Sqrt(data.x2)
        data.y2 = math.Sqrt(data.y2)
        efmt.Println("val:", data.x1,data.y1,data.x2,data.y2)

}
