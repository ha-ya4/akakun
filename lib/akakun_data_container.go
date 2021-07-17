package lib

import (
	"path/filepath"

	"github.com/syndtr/goleveldb/leveldb"

	"github.com/ha-ya4/hlib"
)

type AkakunDataContainer struct {
	DB      *leveldb.DB
	PrjRoot string
	Groups  []AkakunAccount `json:"groups"`
}

type AkakunAccount struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

const akakunGroupPath = "./akakun_group.json"

func CreateDataContainer() (*AkakunDataContainer, error) {
	d := &AkakunDataContainer{}

	err := d.ReadGroupFromFile()
	// ファイルがなければ作成
	if err != nil {
		if err.Error() == "open ./akakun_group.json: no such file or directory" {
			return d, d.SaveGroup()
		}
	}

	d.PrjRoot, err = filepath.Abs(".")
	if err != nil {
		return d, err
	}

	return d, err
}

func (adc AkakunDataContainer) CloseDB() error {
	if adc.DB == nil {
		return nil
	}
	return adc.DB.Close()
}

func (adc *AkakunDataContainer) ReadGroupFromFile() error {
	return hlib.JSONUnmarshalFromFile(akakunGroupPath, &adc.Groups)
}

func (adc *AkakunDataContainer) SaveGroup() error {
	return hlib.WriteFileJSONPretty(adc.Groups, akakunGroupPath, 0666)
}

func (adc *AkakunDataContainer) RegisterGroup(account AkakunAccount) error {
	adc.Groups = append(adc.Groups, account)
	return adc.SaveGroup()
}
