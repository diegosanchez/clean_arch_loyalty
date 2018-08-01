package entity

type Databag struct {
	data map[string]interface{}
}

func NewDatabag() *Databag {
	result := new(Databag)

	result.data = make(map[string]interface{})

	return result
}

func (this *Databag) set(key string, value interface{}) *Databag {
	(this.data)[key] = value
	return this
}

func (this *Databag) AsMap(dumpArea *map[string]interface{}) *Databag {
	for k, v := range this.data {
		(*dumpArea)[k] = v
	}
	return this
}
