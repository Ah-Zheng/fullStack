package utils

import (
	"encoding/base64"
	"log"
	"time"

	"gocv.io/x/gocv"
)

func CaptureImg() string {
	webcam, err := gocv.OpenVideoCapture(0)

	if err != nil {
		log.Fatal(err)
	}

	defer webcam.Close()

	img := gocv.NewMat()
	time.Sleep(time.Millisecond * 100)
	webcam.Read(&img)

	data, err := gocv.IMEncode(".png", img)

	if err != nil {
		log.Fatal(err)
	}

	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(data.GetBytes())
}
