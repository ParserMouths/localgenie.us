package storyblok

import (
	"bytes"
	"encoding/json"
	"fmt"
	domain_storyblok "htf/src/internal/domain/storyblok"
	"htf/src/utils"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	orderedmap "github.com/wk8/go-ordered-map"
)

const (
	assetURL       = "https://mapi.storyblok.com/v1/spaces/%s/assets/"
	storiesURL     = "https://mapi.storyblok.com/v1/spaces/%s/stories/"
	updateStoryURL = "https://mapi.storyblok.com/v1/spaces/%s/stories/%s"
)

type assetPayload struct {
	Filename string `json:"filename"`
}

func Upload(client *http.Client, url string, values *orderedmap.OrderedMap) (err error) {
	// Prepare a form that you will submit to that URL.
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for pair := values.Oldest(); pair != nil; pair = pair.Next() {
		fmt.Println("adding ", pair.Key)
		var fw io.Writer
		if x, ok := pair.Value.(io.Closer); ok {
			defer x.Close()
		}
		// Add an image file
		if x, ok := pair.Value.(*os.File); ok {
			if fw, err = w.CreateFormFile(fmt.Sprint(pair.Key), x.Name()); err != nil {
				return
			}
		} else {
			// Add other fields
			if fw, err = w.CreateFormField(fmt.Sprint(pair.Key)); err != nil {
				return
			}
		}
		if fmt.Sprint(pair.Key) == "file" {
			if _, err = io.Copy(fw, pair.Value.(*os.File)); err != nil {
				return err
			}
		} else {
			fmt.Println(fmt.Sprint(pair.Value))
			if _, err = io.Copy(fw, pair.Value.(*strings.Reader)); err != nil {
				return err
			}
		}
	}

	w.Close()

	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		log.Panic(err)
		return
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())

	// Submit the request
	res, err := client.Do(req)
	if err != nil {
		log.Panic(err)
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	fmt.Println("response body stage 2")
	fmt.Println(string(body))

	fmt.Println("Status code", res.StatusCode)
	return
}

func mustOpen(f string) *os.File {
	r, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	fmt.Println("opened file")
	return r
}

func CreateAsset(config *utils.Config, localPath string) (string, error) {
	pathArr := strings.Split(localPath, "/")
	payload := &assetPayload{
		Filename: pathArr[len(pathArr)-1],
	}
	json_data, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("https://mapi.storyblok.com/v1/spaces/%s/assets/", config.SpaceID), bytes.NewBuffer(json_data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", config.StoryBlokOAuth)

	c := *http.DefaultClient
	r, err := c.Do(req)

	if err != nil {
		log.Panic(err)
		return "", err
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	fmt.Println("response body")
	var response domain_storyblok.SingleResponseObject
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Panic(err)
		return "", err
	}

	values := orderedmap.New()
	values.Set("key", strings.NewReader(response.Fields.Key))
	values.Set("acl", strings.NewReader(response.Fields.Acl))
	values.Set("Expires", strings.NewReader(response.Fields.Expires))
	values.Set("Cache-Control", strings.NewReader(response.Fields.CacheControl))
	values.Set("Content-Type", strings.NewReader(response.Fields.ContentType))
	values.Set("policy", strings.NewReader(response.Fields.Policy))
	values.Set("X-Amz-Credential", strings.NewReader(response.Fields.Xamzcredential))
	values.Set("X-Amz-Algorithm", strings.NewReader(response.Fields.Xamzalgorithm))
	values.Set("X-Amz-Date", strings.NewReader(response.Fields.Xamzdate))
	values.Set("X-Amz-Signature", strings.NewReader(response.Fields.Xamzsignature))
	values.Set("file", mustOpen(localPath)) // lets assume its this fil)

	err = Upload(&c, response.PostUrl, values)
	if err != nil {
		log.Panic(err)
		return "", err
	}
	req, err = http.NewRequest("GET", fmt.Sprintf(assetURL, config.SpaceID)+string(response.Id)+"/finish_upload", bytes.NewBuffer(json_data))
	req.Header.Set("Authorization", config.StoryBlokOAuth)

	fmt.Println("https://a.storyblok.com/" + response.Fields.Key + " uploaded!")

	// response.Fields
	return fmt.Sprintf("https://a.storyblok.com/%s", response.Fields.Key), nil
}

func CreateStore(config *utils.Config, payload *domain_storyblok.StoryPayload) (string, error) {
	json_data, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf(storiesURL, config.SpaceID), bytes.NewBuffer(json_data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", config.StoryBlokOAuth)

	if err != nil {
		log.Fatal(err)
		return "", err
	}
	c := http.DefaultClient
	resp, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		return "", err
	}

	var resObject domain_storyblok.StoryPayload
	err = json.Unmarshal(r, &resObject)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return fmt.Sprint(resObject.Story.StoryID), nil
}

func UpdateStory(config *utils.Config, storyId string, payload *domain_storyblok.StoryPayload) (string, error) {
	json_data, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	fmt.Println(fmt.Sprintf(updateStoryURL, config.SpaceID, storyId))
	req, err := http.NewRequest("PUT", fmt.Sprintf(updateStoryURL, config.SpaceID, storyId), bytes.NewBuffer(json_data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", config.StoryBlokOAuth)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	c := http.DefaultClient
	resp, err := c.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	fmt.Println(string(r))

	if err != nil {
		return "", err
	}

	return string(r), nil
}
