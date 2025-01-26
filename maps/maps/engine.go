package maps

type MapSource map[string]interface{}

func (that MapSource) Fetch() (map[string]interface{}, error) {
	return that, nil
}
