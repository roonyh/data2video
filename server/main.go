package main

import (
	"bytes"
	"data2video/localutil"
	"data2video/video/column"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os/exec"
	"time"

	"net/http"

	"github.com/fogleman/gg"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/create", create)
	http.Handle("/videos/", http.StripPrefix("/videos/", http.FileServer(http.Dir("./videos"))))
	http.ListenAndServe(":7000", nil)
}

type Input struct {
	Data []float64 `json:"data"`
}

func create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	var i *Input
	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	name := render(i.Data)
	fmt.Fprint(w, name)
}

func render(data []float64) string {
	dc := gg.NewContext(1280, 720)
	xLabels := []string{"A", "B", "C", "D", "E", "F"}
	b := column.New(dc, 1280, 720, 40, 20, 50.0, data, xLabels)

	name := localutil.RandStringBytes(8)
	in, _, cmd := startRecording(name)

	totalFrames := 30
	for i := 0; i <= totalFrames; i++ {
		step := float64(i) / float64(totalFrames)
		b.Render(step)
		dc.EncodePNG(in)
	}
	for i := 0; i <= 2*totalFrames; i++ {
		step := 1.0
		b.Render(step)
		dc.EncodePNG(in)
	}
	for i := totalFrames; i >= 0; i-- {
		step := float64(i) / float64(totalFrames)
		b.Render(step)
		dc.EncodePNG(in)
	}
	in.Close()
	cmd.Wait()
	return name + ".webm"
}

func startRecording(id string) (io.WriteCloser, *bytes.Buffer, *exec.Cmd) {
	// ffmpeg -f image2pipe -vcodec png -r 10 -i - -f webm -
	cmd := exec.Command("ffmpeg", "-f", "image2pipe", "-vcodec", "png", "-r", "30", "-i", "-", "-vcodec", "libvpx-vp9", "-r", "30", "-b:v", "10M", "-f", "webm", "./videos/"+id+".webm")
	fmt.Print()
	fmt.Println(cmd.ProcessState)

	stdin, err := cmd.StdinPipe()
	fmt.Println(err)

	var out bytes.Buffer
	cmd.Stderr = &out

	err = cmd.Start()
	if err != nil {
		fmt.Println(err)
	}

	return stdin, &out, cmd
}
