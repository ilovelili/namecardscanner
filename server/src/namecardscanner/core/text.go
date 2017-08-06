// Package core core logic handling business card parsing
package core

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"path"

	"namecardscanner/util"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	vision "google.golang.org/api/vision/v1"
)

var (
	client *http.Client
)

func init() {
	serviceAccountJSONFile := path.Join(util.GetWorkingDirectory(), "serviceaccount.json")
	dat, err := ioutil.ReadFile(serviceAccountJSONFile)
	if err != nil {
		log.Fatalf("Unable to read service account file %v", err)
		panic(err)
	}

	conf, err := google.JWTConfigFromJSON(dat, vision.CloudPlatformScope)
	if err != nil {
		log.Fatalf("Unable to acquire generate config: %v", err)
	}

	client = conf.Client(oauth2.NoContext)
}

// VisionResponse response provided by google vision API
type VisionResponse struct {
	Text    string
	Success bool
}

// DetectText detect text
func DetectText(buffer bytes.Buffer) *VisionResponse {
	res := &VisionResponse{}

	service, err := vision.New(client)
	if err != nil {
		res.Text = err.Error()
		res.Success = false
		return res
	}

	imagestream := buffer.Bytes()
	req := &vision.AnnotateImageRequest{
		// Apply image which is encoded by base64
		Image: &vision.Image{
			Content: base64.StdEncoding.EncodeToString(imagestream),
		},
		// Apply features to indicate what type of image detection
		Features: []*vision.Feature{
			{
				Type: string(TextDetection),
			},
		},
	}

	// BatchAnnotateImagesRequest: Multiple image annotation requests are batched into a single service call.
	batch := &vision.BatchAnnotateImagesRequest{
		Requests: []*vision.AnnotateImageRequest{req},
	}

	// Annotate: Run image detection and annotation for a batch of images. Do executes the "vision.images.annotate" call
	response, err := service.Images.Annotate(batch).Do()
	if err != nil {
		res.Text = err.Error()
		res.Success = false
		return res
	}

	// TextAnnotations: If present, text detection completed successfully
	// Description: Entity textual description
	if annotations := response.Responses[0].TextAnnotations; len(annotations) > 0 {
		res.Text = annotations[0].Description
		res.Success = true
		return res
	}

	res.Text = ""
	res.Success = false
	return res
}
