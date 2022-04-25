package main

import (
	"context"
	"encoding/binary"
	"github.com/docker/docker/api/types"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/coci/dumitor/docker"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func homePage(w http.ResponseWriter, r *http.Request) {
	ctnList := docker.ContainerList()

	tmpl := template.Must(template.ParseFiles("./static/list.html"))
	err := tmpl.Execute(w, ctnList)
	if err != nil {
		return 
	}
}

func logs(w http.ResponseWriter, r *http.Request) {

	conn, _ := upgrader.Upgrade(w, r, nil)
	
	id := strings.TrimPrefix(r.URL.Path, "/logs/")
	id = strings.Replace(id, "/", "", -1)
	
	cli := docker.DockerClient()
	containerLog, _ := cli.ContainerLogs(context.Background(), id, types.ContainerLogsOptions{
		ShowStderr: true,
		ShowStdout: true,
		Timestamps: false,
		Follow:     true,
		Tail:       "500",
	})

	header := make([]byte, 8)
	for {
		_, err := containerLog.Read(header)
		if err == io.EOF {
			break
		}


		count := binary.BigEndian.Uint32(header[4:])
		data := make([]byte, count)

		_, _ = containerLog.Read(data)

		err = conn.WriteMessage(1,data)
		if err != nil {
			return
		}

	}


}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/logs/", logs)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
