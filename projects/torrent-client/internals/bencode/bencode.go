package bencode

import (
	"errors"
	"strconv"
)

type Parser struct {
	data []byte
	pos int
}

func Decode(data []byte) (interface{}, error) {
	p := &Parser{
		data: data,
		pos: 0,
	}

	return p.parse()
}

func (p *Parser) parse() (interface{}, error) {
	if p.pos >= len(p.data) {
		return nil, errors.New("End of input before passing value")
	}

	exp := p.data[p.pos]

	switch exp {
	case 'i':
		return p.parseInt()
	case 'l':
		return p.parseList()
	case 'd':
		return p.parseDict()
	default:
		if p.data[p.pos] >= '0' && p.data[p.pos] <= '9' {
			return p.parseString()
		}
	}

	return nil, errors.New("Invalid bencode\n")
}

func (p *Parser) parseInt() (int, error) {
	p.pos++

	start := p.pos
	for p.data[p.pos] != 'e' {
		p.pos++;
	}

	numStr := string(p.data[start:p.pos])
	p.pos++

	return strconv.Atoi(numStr)
}

func (p *Parser) parseList() ([]interface{}, error) {
	p.pos++

	var list []interface{}

	for p.data[p.pos] != 'e' {
		v, err := p.parse()
		if err != nil {
			return nil, err
		}

		list = append(list, v)
	}

	p.pos++

	return list, nil
}

func (p *Parser) parseDict() (map[string]interface{}, error) {
	p.pos++

	dict := make(map[string]interface{})

	for p.data[p.pos] != 'e' {
		key, err := p.parseString()
		if err != nil {
			return nil, err
		}

		val, err := p.parse()
		if err != nil {
			return nil, err
		}

		dict[key] = val
	}

	p.pos++

	return dict, nil
}

func (p *Parser) parseString() (string, error) {
	start := p.pos

	for p.data[p.pos] != ':' {
		p.pos++
	}

	strLen := string(p.data[start:p.pos])

	length, err := strconv.Atoi(strLen)
	if err != nil {
		return "", err
	}

	p.pos++
	str := string(p.data[p.pos : p.pos+length])
	p.pos += length

	return str, nil
}

