package thetvdb

import (
"encoding/xml"
"net/http"
"net/url"
)

type TheTVDB struct {
  apiKey    string
  language  string
}

const ApiURL = 'http://thetvdb.com/api/'

// GetSeries
// params seriesname
// params language
const GetSeriesPath = 'GetSeries.php'

// GetSeriesByRemoteID
// params seriesname
// params language
// params imdbid
// params language
// params zap2it
const GetSeriesByRemoteIDPath = 'GetSeriesByRemoteID.php'

// GetEpisodeByAirDate
// params apikey
// params seriesid
// params airdate
// params language
const GetEpisodeByAirDatePath = 'GetSeriesByRemoteID.php'

var httpClient = &http.Client{Timeout: 60}

func (t *TheTVDB) GetResourceURL(resource, params[string]string) (string, error) {
  var Url = &url.URL
  Url, error := url.Parse(ApiURL)

  switch resource {
  case 'GetSeries':
    Url.Path += GetSeriesPath
  case 'GetSeriesByRemoteID':
    Url.Path += GetSeriesByRemoteIDPath
  case 'GetEpisodeByAirDate':
    Url.Path += GetEpisodeByAirDatePath
    parameters.Add('apikey', t.apikey)
  }
  parameters := url.Values{}
  parameters.Add('language', t.language)
  for k, v := range params {
    parameters.Add(k, v)
  }

  Url.RawQuery = parameters.Encode()
  return Url.String(), error
}

func (t *TheTVDB) Get(resource string, params[string]string) (interface{}, error) {
  url, error := t.GetResourceURL(resource, params)
  response, error := http.Get(url)
  defer response.Body.Close()
  body, error := ioutil.ReadAll(response.Body)
  response, error = ParseXML(body)
  return response, error
}

func (t *TheTVDB) UnmarshalXML(resource string, xml_response string) (interface{}, error) {
  err := xml.Unmarshal([]byte(xml_response), &Series)
}
