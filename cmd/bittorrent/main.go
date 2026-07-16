package main

import ( 
	"fmt"
	"bittorrent-client/internal/bencode"
)

/*
Features I would like to add:
- Bencode Parser: Decodes whatever bencode data we recieve
- .torrent File Parsing: Extracting Tracker Url, Info Hash, Piece Length and Piece Hashes
- Peer Discovery: Communicating with HTTP trackers to fetch a list of available peers
- Peer Handshaking: Establishes TCP connections and performs BitTorrent handshakes
- Piece Downloading: Requesting and Downloading specific pieces of a file.
- Concurrent Downloading: Downloading multiple pieces simultaneously from different peers.
- Magnet Link Support
*/

func main() {
	data := map[string]any{
    	"announce": "http://tracker.example.com",
    	"info": map[string]any{
        	"name": "example.txt",
        	"length": 12345,
		"piece length": 16384,
        	"pieces": "abcdefgh",
    	},
	}

	result, err := bencode.EncodeBencode(data)
	fmt.Println(result)
	fmt.Println(err)
}
