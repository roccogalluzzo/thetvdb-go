package thetvdb

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type TheTVDB struct {
	ApiKey   string
	Language string
}

const (
	ApiURL = "http://thetvdb.com/api/"

	// GetSeries
	// params seriesname
	// params language
	GetSeriesPath = "GetSeries.php"

	// GetSeriesByRemoteID
	// params seriesname
	// params language
	// params imdbid
	// params language
	// params zap2it
	GetSeriesByRemoteIDPath = "GetSeriesByRemoteID.php"

	// GetEpisodeByAirDate
	// params apikey
	// params seriesid
	// params airdate
	// params language
	GetEpisodeByAirDatePath = "GetSeriesByRemoteID.php"

	// GetSeriesById
	// apikey, series_id, language
	GetBaseSeriesByIdPath = "%v/series/%v/%v.xml"
	GetFullSeriesByIdPath = "%v/series/%v/all/%v.xml"

	// GetEpisodeById
	// apikey, episode_id, language
	GetEpisodeByIdPath = "%v/episodes/%v/%v.xml"
)

var httpClient = &http.Client{Timeout: 60}

func (t TheTVDB) GetResourceURL(resource string, params map[string]string) (string, error) {

	var Url = &url.URL{}
	Url, error := url.Parse(ApiURL)
	parameters := url.Values{}

	switch resource {
	case "GetSeries":
		Url.Path += GetSeriesPath
	case "GetBaseSeriesById":
		Url.Path += fmt.Sprintf(GetBaseSeriesByIdPath, t.ApiKey, params["series_id"], t.Language)
	case "GetFullSeriesById":
		Url.Path += fmt.Sprintf(GetFullSeriesByIdPath, t.ApiKey, params["series_id"], t.Language)
	case "GetEpisodeById":
		Url.Path += fmt.Sprintf(GetEpisodeByIdPath, t.ApiKey, params["episode_id"], t.Language)
	case "GetSeriesByRemoteID":
		Url.Path += GetSeriesByRemoteIDPath
	case "GetEpisodeByAirDate":
		Url.Path += GetEpisodeByAirDatePath
		parameters.Add("apikey", t.ApiKey)
	}

	if (resource == "GetEpisodeByAirDate") || (resource == "GetSeriesByRemoteID") || (resource == "GetSeries") {
		parameters.Add("language", t.Language)
		for k, v := range params {
			parameters.Add(k, v)
		}
		Url.RawQuery = parameters.Encode()
	}

	return Url.String(), error
}

func (t TheTVDB) Get(url string) []byte {
	response, _ := http.Get(url)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return body
}

func UnmarshalXML(response []byte, resource interface{}) {
	xml.Unmarshal(response, &resource)
}
