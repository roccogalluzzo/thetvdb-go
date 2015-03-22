package thetvdb

type Episode struct {
	ID                    uint64 `xml:"id"`
	CombinedEpisodeNumber string `xml:"Combined_episodenumber"`
	CombinedSeason        uint64 `xml:"Combined_season"`
	DvdChapter            string `xml:"DVD_chapter"`
	DvdDiscID             string `xml:"DVD_discid"`
	DvdEpisodeNumber      string `xml:"DVD_episodenumber"`
	DvdSeason             string `xml:"DVD_season"`
	Director              string `xml:"Director"`
	EpImgFlag             string `xml:"EpImgFlag"`
	EpisodeName           string `xml:"EpisodeName"`
	EpisodeNumber         uint64 `xml:"EpisodeNumber"`
	FirstAired            string `xml:"FirstAired"`
	GuestStars            string `xml:"GuestStars"`
	ImdbID                string `xml:"IMDB_ID"`
	Language              string `xml:"language"`
	Overview              string `xml:"Overview"`
	ProductionCode        string `xml:"ProductionCode"`
	Rating                string `xml:"Rating"`
	SeasonNumber          uint64 `xml:"SeasonNumber"`
	Writer                string `xml:"Writer"`
	AbsoluteNumber        string `xml:"absolute_number"`
	Filename              string `xml:"filename"`
	LastUpdated           string `xml:"lastupdated"`
	SeasonID              uint64 `xml:"seasonid"`
	SeriesID              uint64 `xml:"seriesid"`
}

type EpisodeList struct {
	Episodes []*Episode `xml:"Episode"`
}

func (t *TheTVDB) GetEpisodeById(episode_id string) *Episode {

	params := make(map[string]string)
	params["episode_id"] = episode_id

	url, _ := t.GetResourceURL("GetEpisodeById", params)
	response := t.Get(url)
	var list EpisodeList
	UnmarshalXML(response, &list)
	return list.Episodes[0]
}
