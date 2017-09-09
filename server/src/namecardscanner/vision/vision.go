// Package vision core logic handling business card image parsing
package vision

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"

	"namecardscanner/naturallanguage"
	"namecardscanner/util"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	vision "google.golang.org/api/vision/v1"
)

var (
	client         *http.Client
	requiredFields = []string{"ORGANIZATION", "PERSON", "LOCATION"}
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
	ParsedItems []*naturallanguage.ParsedItem
	Success     bool
}

// DetectTextByBase64 detect text by base64 encoded image content
func DetectTextByBase64(content string) *VisionResponse {
	req := &vision.AnnotateImageRequest{
		// Apply image which is encoded by base64
		Image: &vision.Image{
			Content: content,
		},
		// Apply features to indicate what type of image detection
		Features: []*vision.Feature{
			{
				Type: string(TextDetection),
			},
		},
	}

	return detectText(req)
}

// DetectTextByImageStream detect text by image stream uploaded
func DetectTextByImageStream(buffer bytes.Buffer) *VisionResponse {

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

	return detectText(req)
}

func detectText(req *vision.AnnotateImageRequest) *VisionResponse {
	res := &VisionResponse{}

	service, err := vision.New(client)
	if err != nil {
		res.Success = false
		return res
	}

	// BatchAnnotateImagesRequest: Multiple image annotation requests are batched into a single service call.
	batch := &vision.BatchAnnotateImagesRequest{
		Requests: []*vision.AnnotateImageRequest{req},
	}

	// Annotate: Run image detection and annotation for a batch of images. Do executes the "vision.images.annotate" call
	response, err := service.Images.Annotate(batch).Do()
	if err != nil {
		res.Success = false
		return res
	}

	// TextAnnotations: If present, text detection completed successfully
	// Description: Entity textual description
	if annotations := response.Responses[0].TextAnnotations; len(annotations) > 0 {
		text := annotations[0].Description

		fmt.Println(text)

		parsedEntities, err := naturallanguage.Parse(text)
		if err != nil {
			res.Success = false
			return res
		}

		items := make([]*naturallanguage.ParsedItem, 0)
		for _, entity := range parsedEntities.Entities {
			if contains(requiredFields, entity.Type) {
				items = append(items, &naturallanguage.ParsedItem{Name: entity.Name, Type: entity.Type})
			}
		}

		res.ParsedItems = items
		res.Success = len(res.ParsedItems) > 0
		return res
	}

	res.Success = false
	return res
}

// Array.contains helper
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
