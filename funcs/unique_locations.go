package funcs

func Unique(locations LocationsDATA) LocationsDATA {
	MyMap := make(map[string]struct{})

	var UniqueLocationsHere []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
	}

	for _, TheStructs := range locations.Index {

		result := []string{}

		for _, location := range TheStructs.Locations {
			if isunique(location, MyMap) {
				MyMap[location] = struct{}{}
				result = append(result, location)
			} else {
				continue
			}
		}

		UniqueLocationsHere = append(UniqueLocationsHere, struct {
			ID        int      `json:"id"`
			Locations []string `json:"locations"`
		}{ID: TheStructs.ID, Locations: result})
	}

	locations.Index = UniqueLocationsHere
	return locations
}

func isunique(Location string, Map map[string]struct{}) bool {
	_, See := Map[Location]

	return !See
}
