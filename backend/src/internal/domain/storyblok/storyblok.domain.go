package domain_storyblok

type StoryContent struct {
	Component   string `json:"component"`
	StallID     string
	OwnerID     string
	StallName   string
	IsOpen      int
	CreatedAt   string
	Rating      string
	Offering    string
	AboutVendor string
	LastActive  string
	Latitude    float32
	Longitude   float32
	Assets      []string
}

type StoryUpdateContent struct {
}

type StoryStruct struct {
	Name    string        `json:"name"`
	Slug    string        `json:"slug"`
	Content *StoryContent `json:"content"`
	StoryID int           `json:"id"`
}

type StoryPayload struct {
	Story   *StoryStruct `json:"story"`
	Publish int          `json:"publish"`
}

type SroFields struct {
	Key            string `json:"key"`
	Acl            string `json:"acl"`
	Expires        string `json:"Expires"`
	CacheControl   string `json:"Cache-Control"`
	ContentType    string `json:"Content-Type"`
	Policy         string `json:"policy"`
	Xamzcredential string `json:"x-amz-credential"`
	Xamzalgorithm  string `json:"x-amz-algorithm"`
	Xamzdate       string `json:"x-amz-date"`
	Xamzsignature  string `json:"x-amz-signature"`
}

type SingleResponseObject struct {
	Id        int        `json:"id"`
	PrettyUrl string     `json:"pretty_url"`
	PublicUrl string     `json:"public_url"`
	Fields    *SroFields `json:"fields"`
	PostUrl   string     `json:"post_url"`
}

type AssetPayload struct {
	Filename string `json:"filename"`
}
