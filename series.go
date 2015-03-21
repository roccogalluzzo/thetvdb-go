package thetvdb


type Series struct {
  ID            uint64   `xml:"id"`
  SeriesID      string   `xml:"seriesid"`
  Banner        string   `xml:"banner"`
  Zap2ItID      string   `xml:"zap2it_id"`
  FirstAired    string   `xml:"FirstAired"`
  ImdbID        string   `xml:"IMDB_ID"`
  Language      string   `xml:"language"`
  Network       string   `xml:"Network"`
  Overview      string   `xml:"Overview"`
  SeriesName    string   `xml:"SeriesName"`
  AliasNames    string   `xml:"AliasNames"`
}

type SeriesList struct {
  Series []*Series `xml:"Series"`
}
type PipeList []string

func (t *TheTVDB) GetSeriesByRemoteID(seriesid string) (*Series) {

  params := make(map[string]string)
  params["imdbid"] = seriesid

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
