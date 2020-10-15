package txtdb

import (
	"encoding/json"
	"errors"
    "fmt"
	"log"
	"strings"
	"time"

    "github.com/andersonlira/item-api/domain"
	"github.com/andersonlira/goutils/io"
	"github.com/andersonlira/goutils/str"
	"sort"
)

//GetItemList return all items 
func GetItemList() []domain.Item {
	list := []domain.Item{}
    fileName := fmt.Sprintf("bd/%ss.json", "Item");
	listTxt, _ := io.ReadFile(fileName)
	json.Unmarshal([]byte(listTxt), &list)
	sort.Slice(list, func(i, j int) bool {
		return list[i].Name < list[j].Name
	})
	return list
}

//GetItemByID return item by its id
func GetItemByID(ID string) (domain.Item, error) {
	list := GetItemList()
	for idx, _ := range list {
		if(list[idx].ID == ID){
			return list[idx],nil
		}
	}
	return domain.Item{}, errors.New("NOT_FOUND")
}



//SaveItem saves a Item object
func SaveItem(it domain.Item) domain.Item {
	list := GetItemList()
	it.ID = str.NewUUID()
	it.Initial = strings.ToUpper(it.Initial)
	it.Name = strings.Title(it.Name)
	it.CreatedAt = time.Now()
	list = append(list, it)
	writeItem(list)
	return it
}

//UpdateItem( updates a Item object
func UpdateItem(ID string, it domain.Item) domain.Item{
	list := GetItemList()
	for idx, _ := range list {
		if(list[idx].ID == ID){
			list[idx] = it
			list[idx].ID = ID
			list[idx].Initial = strings.ToUpper(it.Initial)
			list[idx].Name = strings.Title(it.Name)
			list[idx].UpdatedAt = time.Now()
			writeItem(list)
			return list[idx]
		}
	}
	return it
}

//DeleteItem delete object by giving ID
func DeleteItem(ID string) bool {
	list := GetItemList()
	for idx, _ := range list {
		if(list[idx].ID == ID){
			list = append(list[:idx], list[idx+1:]...)
			writeItem(list)
			return true
		}
	}
	return false
}

func writeItem(list []domain.Item) {
	b, err := json.Marshal(list)
	if err != nil {
		log.Println("Error while writiong file items")
		return
	}
	io.WriteFile(fmt.Sprintf("bd/%ss.json", "Item"), string(b))
}

