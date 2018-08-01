package adapter

import (
	"encoding/json"
	entity "github.com/diegosanchez/clean_arch_loyalty/entity"
	"io/ioutil"
	"os"
)

type TestItemGateway struct {
}

func NewTestItemGateway() *TestItemGateway {
	return new(TestItemGateway)
}

func (this *TestItemGateway) ItemById(itemId *entity.ItemId) *entity.Item {
	itemFromApi := this.itemFromApiFromFile(itemId)

	return entity.NewItem(itemId, itemFromApi.BasePrice)
}

func (this *TestItemGateway) itemFromApiFromFile(itemId *entity.ItemId) *ItemFromApi {
	dir, _ := os.Getwd()

	fixtureFile := itemId.RenderPattern(dir + "/../fixtures/items/item_{site}_{nro}.json")

	// Open our jsonFile
	jsonFile, err := os.Open(fixtureFile)

	// if we os.Open returns an error then handle it
	if err != nil {
		panic(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var itemFromApi ItemFromApi

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &itemFromApi)

	return &itemFromApi
}
