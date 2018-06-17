package albumart

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"github.com/buger/jsonparser"
)

func GetArtwork(ArtistName string, AlbumSongName string) string {
	artworkUrlToReturn := ""
	whitespaceRegEx, _ := regexp.Compile(`\s`)
	ArtistName = whitespaceRegEx.ReplaceAllString(ArtistName, "+")
	AlbumSongName = whitespaceRegEx.ReplaceAllString(AlbumSongName, "+")
	response, err := http.Get("https://itunes.apple.com/search?term=" + ArtistName + "+" + AlbumSongName + "&limit=1&entity=song")
  if err != nil {
    fmt.Printf("%s", err)
  } else {
		defer response.Body.Close()
		data, readErr := ioutil.ReadAll(response.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}
		resultCount, _ := jsonparser.GetInt(data, "resultCount")
		if resultCount == 0 {
			artworkUrlToReturn = "Not Found"
		} else {
			artworkUrlToReturn, _ = jsonparser.GetString(data, "results", "[0]", "artworkUrl100")
			makeLarge, _ := regexp.Compile("100x100bb")
  		artworkUrlToReturn = makeLarge.ReplaceAllString(artworkUrlToReturn, "1200x1200bb")
		}
	}
	return artworkUrlToReturn
}
