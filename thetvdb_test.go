package thetvdb

import "testing"
//import "fmt"

func Test_GetResourceURL(t *testing.T) {
  tvdb := &TheTVDB{ApiKey: "A76837DFE272D3F8", Language: "en"}
  params := make(map[string]string)
  params["seriesname"] = "breaking bad"
  bb, _ := tvdb.GetResourceURL("GetSeries", params)

  if (bb != "http://thetvdb.com/api/GetSeries.php?language=en&seriesname=breaking+bad") {
    t.Error(".GetResourceURL did not work as expected.")
  }
}

func Test_GetSeriesByRemoteID(t *testing.T) {
  tvdb := &TheTVDB{ApiKey: "A76837DFE272D3F8", Language: "en"}
  series := tvdb.GetSeriesByRemoteID("tt0903747")
  //fmt.Printf("%+v\n", series)

  if (series.SeriesName != "Breaking Bad") {
    t.Error(".GetSeriesByRemoteID did not work as expected.")
  }
}

func Test_GetSeries(t *testing.T) {
  tvdb := &TheTVDB{ApiKey: "A76837DFE272D3F8", Language: "en"}
  series := tvdb.GetSeries("Breaking Bad").Series[0]
  //fmt.Printf("%+v\n", series)

  if (series.SeriesName != "Breaking Bad") {
    t.Error(".GetSeriesByRemoteID did not work as expected.")
  }
}

func Test_GetBaseSeriesById(t *testing.T) {
  tvdb := &TheTVDB{ApiKey: "A76837DFE272D3F8", Language: "en"}
  series := tvdb.GetBaseSeriesById("81189")

  if (series.SeriesName != "Breaking Bad") {
    t.Error(".GetSeriesByRemoteID did not work as expected.")
  }
}

func Test_GetFullSeriesById(t *testing.T) {
  tvdb := &TheTVDB{ApiKey: "A76837DFE272D3F8", Language: "en"}
  series := tvdb.GetFullSeriesById("81189")

  if (series.SeriesName != "Breaking Bad") {
    t.Error(".GetSeriesByRemoteID did not work as expected.")
  }
}


