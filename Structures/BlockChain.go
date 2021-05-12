package Structures

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Index        int
	Date         string
	Data         string
	Nonce        int
	PreviousHash []byte
	Hash         []byte
}

func (this *Block) GenerateBlock(merkleU *MerkleTreeUsers, merkleO *MerkleTreeOrders, merkleP *MerkleTreeProducts, merkleS *MerkleTreeShops) *Block {
	if this == nil {
		this = &Block{}
	}

	if this.Index != 0 {
		this.PreviousHash = this.Hash
	}
	var text string
	valido := false
	this.Date = time.Now().String()
	//this.Data = string(merkleU.Root.Hash) + string(merkleO.Root.Hash) + string(merkleP.Root.Hash) + string(merkleS.Root.Hash)
	this.Data = merkleU.GetHash() + merkleO.GetHash() + merkleP.GetHash() + merkleS.GetHash()
	for valido {
		//text = strconv.Itoa(this.Index) + this.Date + hex.EncodeToString(this.PreviousHash) + this.Data + strconv.Itoa(this.Nonce)
		text = strconv.Itoa(this.Index) + this.Date + base64.URLEncoding.EncodeToString(this.PreviousHash) + this.Data + strconv.Itoa(this.Nonce)
		text = strings.ReplaceAll(text, " ", "")
		text = strings.ReplaceAll(text, "\n", "")

		h := sha256.New()
		h.Write([]byte(text))
		hash := h.Sum(nil)

		this.Nonce++

		if hash[0] <= 15 {
			fmt.Println(hash)
			valido = true
			this.Hash = hash
		}
	}

	text = strconv.Itoa(this.Index) + "\n" + this.Date + "\n" + this.Data + "\n" + strconv.Itoa(this.Nonce) + "\n" + base64.URLEncoding.EncodeToString(this.PreviousHash) + "\n" + base64.URLEncoding.EncodeToString(this.Hash)
	fmt.Println(text)
	data := []byte(text)
	ioutil.WriteFile(strconv.Itoa(this.Index), data, 0644)

	this.Nonce = 0

	return this
}
