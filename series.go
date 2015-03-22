package thetvdb

//import "fmt"

// some code taked from https://github.com/garfunkel/go-tvdb/blob/master/tvdb.go
type Series struct {
  ID            uint64   `xml:"id"`
  Actors        PipeList `xml:"Actors"`
  AirsDayOfWeek string   `xml:"Airs_DayOfWeek"`
  AirsTime      string   `xml:"Airs_Time"`
  ContentRating string   `xml:"ContentRating"`
  FirstAired    string   `xml:"FirstAired"`
  Genre         PipeList `xml:"Genre"`
  ImdbID        string   `xml:"IMDB_ID"`
  Language      string   `xml:"language"`
  Network       string   `xml:"Network"`
  NetworkID     string   `xml:"NetworkID"`
  Overview      string   `xml:"Overview"`
  Rating        string   `xml:"Rating"`
  RatingCount   string   `xml:"RatingCount"`
  Runtime       string   `xml:"Runtime"`
  SeriesID      string   `xml:"seriesid"`
  SeriesName    string   `xml:"SeriesName"`
  Status        string   `xml:"Status"`
  Added         string   `xml:"added"`
  AddedBy       string   `xml:"addedBy"`
  Banner        string   `xml:"banner"`
  Fanart        string   `xml:"fanart"`
  LastUpdated   string   `xml:"lastupdated"`
  Poster        string   `xml:"poster"`
  Zap2ItID      string   `xml:"zap2it_id"`
  Seasons       map[uint64][]*Episode
}

type SeriesList struct {
  Series []*Series `xml:"Series"`
}

// PipeList type representing pipe-separated string values.
type PipeList []string

func (t *TheTVDB) GetBaseSeriesById(series_id string) (*Series) {

  params := make(map[string]string)
  params["series_id"] = series_id

  url, _ := t.GetResourceURL("GetBaseSeriesById", params)
  response := t.Get(url)
  var list SeriesList
  UnmarshalXML(response, &list)
  return list.Series[0]
}

func (t *TheTVDB) GetFullSeriesById(series_id string) (*Series) {

  params := make(map[string]string)
  params["series_id"] = series_id

  url, _ := t.GetResourceURL("GetFullSeriesById", params)
  response := t.Get(url)
  var list SeriesList
  UnmarshalXML(response, &list)
  return list.Series[0]
}

func (t *TheTVDB) GetSeriesByRemoteID(imdbid string) (*Series) {

  params := make(map[string]string)
  params["imdbid"] = imdbid

  url, _ := t.GetResourceURL("GetSeriesByRemoteID", params)
  response := t.Get(url)

  var list SeriesList
  UnmarshalXML(response, &list)
  return list.Series[0]
}

func (t *TheTVDB) GetSeries(seriesName string) (SeriesList) {
  params := make(map[string]string)
  params["seriesname"] = seriesName

  url, _ := t.GetResourceURL("GetSeries", params)
  response := t.Get(url)

  var list SeriesList
  UnmarshalXML(response, &list)
  return list
}
