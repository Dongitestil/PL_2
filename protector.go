package main

import (
	"math/rand"
	"strconv"
	"strings"
	"unicode"
)

type Protector struct {
	currentKey string
	hash string
}

func (p *Protector) set_hash(hash string) {
	p.hash = hash
}

func (p *Protector) set_key(key string) {
	p.currentKey = key
}

func (p *Protector) get_key () string {
	return p.currentKey
}

func (p *Protector) get_hash () string {
	return p.hash
}

func (p *Protector) get_hash_str() string {
	var hash string =""
	for i := 0; i < 5; i++ {
		r := '0' + int(6*rand.Float64() + 1)
		hash += string(r)
	}
	return hash
}

func (p *Protector) get_session_key() string {
	var  result string = ""
	for i := 1; i < 10; i++ {
		r := strconv.Itoa(int(9*rand.Float64() + 1))[0]
		result += string(r)
	}
	return result
}

func (p *Protector) verify_hash() bool{
	if len(p.hash) != 5 {
		return false
	}
	for _, char := range p.hash {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func (p *Protector) next_session_key() string{
	if !p.verify_hash(){return "Error"}
	result:=0
	for _,char := range p.hash{
		result += p.calc_hash(char)
	}
	str_result:=strconv.Itoa(result)
	nkey:=strings.Repeat("0",10)+ str_result[:9]
	nkey = nkey[len(nkey)-9:]
	p.currentKey = nkey
	return nkey
}

func (p *Protector) calc_hash(val rune) int{
	if val == 1 {
		keySliceInt, _ := strconv.Atoi(p.currentKey[:5])
		keySliceStr := strconv.Itoa(keySliceInt % 97)
		keyPartStr := "00" + keySliceStr
		keyPart, _ := strconv.Atoi(keyPartStr[len(keyPartStr) - 2:])
		return keyPart
	}else if val == 2{
		result:=""
		for i := 1; i < 10; i++ {
			result += string(p.currentKey[10 -i]) //
		}
		keyPart, _ := strconv.Atoi(result + string(p.currentKey[0]))
		return keyPart
	}else if val == 3{
		keySlicesSum := p.currentKey[10 - 5:] + p.currentKey[:5]
		keyPart, _ := strconv.Atoi(keySlicesSum)
		return keyPart
	}else if val == 4{
		num:=0
		for i:=1;i<9;i++{
			ch:=p.currentKey[i]
			num += int(ch-'0')+41
		}
		return num
	}else if val == 5{
		var ch rune
		num:=0
		for _, char := range p.currentKey {
			ch = rune(int(char) ^ 43)
			if !unicode.IsDigit(ch) {
				ch = rune(int(ch))
			}
			chInt := int(ch)
			num += chInt
		}
		return num
	}else {
		keyInt, _ := strconv.Atoi(p.currentKey)
		keyPart := keyInt + int(val - '0')
		return keyPart
	}
}