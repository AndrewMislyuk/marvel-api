package marvel

type ComicDataWrapper struct {
	Code            int                `json:"code"`
	Status          string             `json:"status"`
	Copyright       string             `json:"copyright"`
	AttributionText string             `json:"attributionText"`
	AttributionHTML string             `json:"attributionHTML"`
	Data            ComicDataContainer `json:"data"`
	Etag            string             `json:"etag"`
}

type ComicDataContainer struct {
	Offset  int     `json:"offset"`
	Limit   int     `json:"limit"`
	Total   int     `json:"total"`
	Count   int     `json:"count"`
	Results []Comic `json:"results"`
}

type Comic struct {
	Id                 int            `json:"id"`
	DigitalId          int            `json:"digitalId"`
	Title              string         `json:"title"`
	IssueNumber        float64        `json:"issueNumber"`
	VariantDescription string         `json:"variantDescription"`
	Description        string         `json:"description"`
	Modified           string         `json:"modified"`
	ISBN               string         `json:"isbn"`
	UPC                string         `json:"upc"`
	DiamondCode        string         `json:"diamondCode"`
	EAN                string         `json:"ean"`
	ISSN               string         `json:"issn"`
	Format             string         `json:"format"`
	PageCount          int            `json:"pageCount"`
	TextObjects        []TextObject   `json:"textObjects"`
	ResourceURI        string         `json:"resourceURI"`
	URLS               []Url          `json:"urls"`
	Series             SeriesSummary  `json:"series"`
	Variants           []ComicSummary `json:"variants"`
	Collections        []ComicSummary `json:"collections"`
	CollectedIssues    []ComicSummary `json:"collectedIssues"`
	Dates              []ComicDate    `json:"dates"`
	Prices             []ComicPrice   `json:"prices"`
	Thumbnail          Image          `json:"thumbnail"`
	Images             []Image        `json:"images"`
	Creators           CreatorList    `json:"creators"`
	Characters         CharacterList  `json:"characters"`
	Stories            StoryList      `json:"stories"`
	Events             EventList      `json:"events"`
}

type TextObject struct {
	Type     string `json:"type"`
	Language string `json:"language"`
	Text     string `json:"text"`
}

type Url struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type SeriesSummary struct {
	ResourceURI string `json:"resourceURI"`
	Name        string `json:"name"`
}

type ComicSummary struct {
	ResourceURI string `json:"resourceURI"`
	Name        string `json:"name"`
}

type ComicDate struct {
	Type string `json:"type"`
	Date string `json:"date"`
}

type ComicPrice struct {
	Type  string  `json:"type"`
	Price float64 `json:"price"`
}

type Image struct {
	Path      string `json:"path"`
	Extension string `json:"extension"`
}

type CreatorList struct {
	Available     int              `json:"available"`
	Returned      int              `json:"returned"`
	CollectionURI string           `json:"collectionURI"`
	Items         []CreatorSummary `json:"items"`
}

type CreatorSummary struct {
	ResourceURI string `json:"resourceURI"`
	Name        string `json:"name"`
	Role        string `json:"role"`
}

type CharacterList struct {
	Available     int                `json:"available"`
	Returned      int                `json:"returned"`
	CollectionURI string             `json:"collectionURI"`
	Items         []CharacterSummary `json:"items"`
}

type CharacterSummary struct {
	ResourceURI string `json:"resourceURI"`
	Name        string `json:"name"`
	Role        string `json:"role"`
}

type StoryList struct {
	Available     int            `json:"available"`
	Returned      int            `json:"returned"`
	CollectionURI string         `json:"collectionURI"`
	Items         []StorySummary `json:"items"`
}

type StorySummary struct {
	ResourceURI string `json:"resourceURI"`
	Name        string `json:"name"`
	Type        string `json:"type"`
}

type EventList struct {
	Available     int            `json:"available"`
	Returned      int            `json:"returned"`
	CollectionURI string         `json:"collectionURI"`
	Items         []EventSummary `json:"items"`
}

type EventSummary struct {
	ResourceURI string `json:"resourceURI"`
	Name        string `json:"name"`
}
