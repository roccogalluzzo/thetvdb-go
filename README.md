# thetvdb-go
Go wrapper for thetvdb.com API

## Examples

Get Series info by IMDB id
```
  tvdb := &TheTVDB{ApiKey: "api_key", Language: "en"}
  series := tvdb.GetSeriesByRemoteID("tt0903747")
```

Get Series by name
```
  tvdb := &TheTVDB{ApiKey: "api_key", Language: "en"}
  series := tvdb.GetSeries("Breaking Bad")
```
