package torrent

import (
	"crypto/sha1"
	"errors"
	"os"
	"torrent-client/internals/bencode"
)

type TorrentFile struct {
	Announce string
	InfoHash [20]byte
	PieceHashes [][20]byte
	PieceLength int
	Length int
	Name string
}

func Open(path string) (*TorrentFile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	decoded, err := bencode.Decode(data)
	if err != nil {
		return nil, err
	}

	root := decoded.(map[string]interface{})
	tf := &TorrentFile{}

	announce, ok := root["announce"].(string)
	if !ok {
		return nil, errors.New("missing announce")
	}

	tf.Announce = announce

	info, ok := root["info"].(map[string]interface{})
	if !ok {
		return nil, errors.New("missing dict info")
	}

	name := info["name"].(string)
	length := info["length"].(int)
	pieceLength := info["pieceLength"].(int)

	tf.Name = name
	tf.Length = length
	tf.PieceLength = pieceLength

	piecesStr := info["pieces"].(string)

	hashes, err := splitPieceHashes([]byte(piecesStr))
	if err != nil {
		return nil, err
	}

	tf.PieceHashes = hashes
	tf.InfoHash = sha1.Sum(extractInfoBytes(data))

	return tf, nil
}

func splitPieceHashes(buf []byte) ([][20]byte, error) {
	hashLen := 20

	if len(buf) % hashLen != 0 {
		return nil, errors.New("invalid piece hashes")
	}

	numHashes := len(buf) / hashLen
	hashes := make([][20]byte, numHashes)

	for i := range numHashes {
		copy(hashes[i][:], buf[i*20:(i+1)*20])
	}

	return hashes, nil
}

func extractInfoBytes(data []byte) []byte {
	start := -1
	count := 0

	for i := 0;i < len(data)-4;i++ {
		if string(data[i:i+4]) == "4:info" {
			start = i+6
			break
		}
	}

	if start == -1 {
		return nil
	}

	for i := start;i < len(data);i++ {
		switch data[i] {
			case 'd', 'l':
				count++
			case 'e':
				count--

				if count == 0 {
					return data[start:i+1]
				}
		}
	}

	return nil
}

