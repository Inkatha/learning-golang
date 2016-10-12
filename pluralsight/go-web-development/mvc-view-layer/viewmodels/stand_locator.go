package viewmodels

type StandLocator struct {
	Title  string
	Active string
}

func GetStandLocator() StandLocator {
	result := StandLocator{
		Title:  "Lemonade Stand Supply - Stand Locator",
		Active: "stand_locator",
	}

	return result
}

type StandLocation struct {
	Lat   float32
	Lng   float32
	Title string
}

func GetStandLocations() []StandLocation {
	result := []StandLocation{
		StandLocation{
			Lat:   37.4217,
			Lng:   -122.075,
			Title: "Matthew's stand",
		},
		StandLocation{
			Lat:   37.4206,
			Lng:   -122.08,
			Title: "Alice's stand",
		},
		StandLocation{
			Lat:   37.4205,
			Lng:   -122.083,
			Title: "Kara's stand",
		},
		StandLocation{
			Lat:   37.414,
			Lng:   -122.09,
			Title: "Fred's stand",
		},
		StandLocation{
			Lat:   37.412,
			Lng:   -122.09,
			Title: "Jake's stand",
		},
		StandLocation{
			Lat:   37.41,
			Lng:   -122.093,
			Title: "Wallace's stand",
		},
		StandLocation{
			Lat:   37.416,
			Lng:   -122.095,
			Title: "Gromit's stand",
		},
		StandLocation{
			Lat:   37.41,
			Lng:   -122.1,
			Title: "Kirk's stand",
		},
		StandLocation{
			Lat:   37.41,
			Lng:   -122.102,
			Title: "Lorelei's stand",
		},
		StandLocation{
			Lat:   37.412,
			Lng:   -122.099,
			Title: "Rebecca's stand",
		},
		StandLocation{
			Lat:   37.407,
			Lng:   -122.1025,
			Title: "Chris's stand",
		},
		StandLocation{
			Lat:   37.423,
			Lng:   -122.1025,
			Title: "Carson's stand",
		},
	}
	return result
}
