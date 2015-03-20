package thetvdb

type Series struct {
  ID            uint64   'xml:"id"'
  Actors        string   'xml:"Actors"'
  AirsDayOfWeek string   'xml:"Airs_DayOfWeek"'
  AirsTime      string   'xml:"Airs_Time"'
  ContentRating string   'xml:"ContentRating"'
  FirstAired    string   'xml:"FirstAired"'
  Genre         string   'xml:"Genre"'
  ImdbID        string   'xml:"IMDB_ID"'
  Language      string   'xml:"Language"'
  Network       string   'xml:"Network"'
  NetworkID     string   'xml:"NetworkID"'
  Overview      string   'xml:"Overview"'
  Rating        string   'xml:"Rating"'
  RatingCount   string   'xml:"RatingCount"'
  Runtime       string   'xml:"Runtime"'
  SeriesID      string   'xml:"SeriesID"'
  SeriesName    string   'xml:"SeriesName"'
  Status        string   'xml:"Status"'
  Added         string   'xml:"added"'
  AddedBy       string   'xml:"addedBy"'
  Banner        string   'xml:"banner"'
  Fanart        string   'xml:"fanart"'
  LastUpdated   string   'xml:"lastupdated"'
  Poster        string   'xml:"poster"'
  Zap2ItID      string   'xml:"zap2it_id"'
  Seasons       map[uint64][]*Episode
}


func (t *TheTVDB) GetSeriesById(seriesId int) (*Series, error) {}

func (t *TheTVDB) GetSeries(seriesName string) ([]Series, error) {}
