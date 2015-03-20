package thetvdb

import (
"encoding/xml"
"net/http"
"net/url"
)

type TheTVDB struct {
  apiKey    string
  lang      string
}

const ApiURL = 'http://thetvdb.com/api'

var httpClient = &http.Client{Timeout: 60}


func (t *TheTVDB) Call(path string) (*http.Request, error) {

}

func (t *TheTVDB) ParseXML(req *http.Request) (response) {

}
