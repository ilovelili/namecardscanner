package naturallanguage

import (
	"io/ioutil"
	"log"
	"namecardscanner/util"
	"net/http"
	"path"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	language "google.golang.org/api/language/v1"
)

var (
	client *http.Client
)

// ParsedItem including name and string field
type ParsedItem struct {
	Name string
	Type string
}

func init() {
	serviceAccountJSONFile := path.Join(util.GetWorkingDirectory(), "serviceaccount.json")
	dat, err := ioutil.ReadFile(serviceAccountJSONFile)
	if err != nil {
		log.Fatalf("Unable to read service account file %v", err)
		panic(err)
	}

	conf, err := google.JWTConfigFromJSON(dat, language.CloudLanguageScope)
	if err != nil {
		log.Fatalf("Unable to acquire generate config: %v", err)
	}

	client = conf.Client(oauth2.NoContext)
}

// Parse code function to parse natural language
func Parse(text string) (*language.AnalyzeEntitiesResponse, error) {
	service, err := language.New(client)
	if err != nil {
		return nil, err
	}

	req := &language.AnalyzeEntitiesRequest{
		Document: &language.Document{
			Content: text,
			Type:    string(PlainText),
		},
		EncodingType: string(Utf8),
	}

	return service.Documents.AnalyzeEntities(req).Do()
}
