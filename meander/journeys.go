package meander

type j struct{
	Name string
	PlaceType []string
}

var Journeys = []interface{}{
	&j{Name:"ロマンティック",PlaceType:[]string{"park", "bar", "movie_theater", "restaurant", "florist", "taxi_stand"}},
	&j{Name:"ショッピング",PlaceType:[]string{"park", "bar", "movie_theater", "restaurant", "florist", "taxi_stand"}},
	&j{Name:"ナイトライフ",PlaceType:[]string{"park", "bar", "clothing_cafe", "restaurant", "florist", "taxi_stand"}},
	&j{Name:"カルチャー",PlaceType:[]string{"museum", "cafe", "cemetery", "restaurant", "florist", "taxi_stand"}},
	&j{Name:"リラックス",PlaceType:[]string{"hair_care", "beauty_salon", "cafe", "spa", "florist", "taxi_stand"}},
}
