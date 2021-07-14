package lib

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"

	"github.com/ha-ya4/hlib"
)

type AkakunDataContainer struct {
	db    *leveldb.DB
	Group []AkakunAccount `json:"group"`
}

type AkakunAccount struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

const akakunGroupPath = "./akakun_group.json"

func CreateDataContainer() (AkakunDataContainer, error) {
	d := AkakunDataContainer{}
	err := d.ReadGroupFromFile()
	// ファイルがなければ作成
	if err != nil {
		if err.Error() == "open ./akakun_group.json: no such file or directory" {
			fmt.Println("ssss")
			return d, d.SaveGroup()
		}
	}
	return d, err
}

func (adc AkakunDataContainer) ReadGroupFromFile() error {
	return hlib.JSONUnmarshalFromFile(akakunGroupPath, &adc.Group)
}

func (adc AkakunDataContainer) SaveGroup() error {
	return hlib.WriteFileJSONPretty(adc.Group, akakunGroupPath, 0666)
}
