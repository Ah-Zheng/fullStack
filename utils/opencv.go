package myUtils

import (
	"encoding/base64"
	"log"

	"gocv.io/x/gocv"
)

func CaptureImg() string {
	webcam, err := gocv.OpenVideoCapture(0)

	if err != nil {
		log.Fatal(err)
	}

	defer webcam.Close()

	img := gocv.NewMat()
	webcam.Read(&img)

	data, err := gocv.IMEncode(".png", img)

	if err != nil {
		log.Fatal(err)
	}

	n := base64.StdEncoding.EncodedLen(data.Len())
	dst := make([]byte, n)
	base64.StdEncoding.Encode(dst, data.GetBytes())

	return "data:image/png;base64," + string(dst)
}
