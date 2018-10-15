# Golang AlbumArt
Simple album art grabbing library that grabs album artwork from iTunes Store

## Get the package
`go get github.com/arjndr/albumart.go`

## Examples

```golang
package main

import (
  "github.com/arjndr/albumart.go"
  "fmt"
)

func main(){
  artWork := albumart.GetArtwork("System Of A Down", "Toxicity")
  fmt.Printf(artWork)
}

// logs https://is3-ssl.mzstatic.com/image/thumb/Music128/v4/f7/7d/78/f77d7847-9831-fead-de75-5e17468f14ac/source/1200x1200bb.jpg
```

## API
`albumart.GetArtwork(ArtistName, AlbumSongName)` - `func` : Returns string (URL) or "Not Found" on 0 results
* `ArtistName` - String : Artist's name
* `AlbumSongName` - String : Album/Song name

## Related
* [itunes-albumart](https://github.com/arjndr/itunes-albumart) - Node.js alternative for this library
* [albumart.js](https://github.com/arjndr/albumart.js) - Browser compatible library using Last.fm API

## License
This library uses Apple's [search API](https://affiliate.itunes.apple.com/resources/documentation/itunes-store-web-service-search-api/). You should read Apple's legal [terms](https://www.apple.com/legal/internet-services/terms/site.html) for more.

[MIT](http://opensource.org/licenses/MIT) &copy; 2018 Akash Rajendra
