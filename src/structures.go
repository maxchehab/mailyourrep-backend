package main

//RequestData structure
type RequestData struct {
	Name        string   `json:"name"`
	Zipcode     string   `json:"addressZipcode"`
	Line1       string   `json:"addressLine1"`
	Line2       string   `json:"addressLine2"`
	City        string   `json:"addressCity"`
	State       string   `json:"addressState"`
	Country     string   `json:"addressCountry"`
	Message     string   `json:"message"`
	StripeToken string   `json:"stripeToken"`
	Reps        []string `json:"reps"`
}

//GoogleCivic structure
type GoogleCivic struct {
	Kind    string `json:"kind"`
	Offices []struct {
		Name            string `json:"name"`
		OfficialIndices []int  `json:"officialIndices"`
	} `json:"offices"`
	Officials []struct {
		Name    string `json:"name"`
		Address []struct {
			Line1 string `json:"line1"`
			Line2 string `json:"line2"`
			City  string `json:"city"`
			State string `json:"state"`
			Zip   string `json:"zip"`
		} `json:"address"`
	} `json:"officials"`
}

//Reps is a collection of the Rep structure
type Reps []Rep

//Rep structure
type Rep struct {
	Name    string `json:"name"`
	Zipcode string `json:"addressZipcode"`
	Line1   string `json:"addressLine1"`
	Line2   string `json:"addressLine2"`
	City    string `json:"addressCity"`
	State   string `json:"addressState"`
	Country string `json:"addressCountry"`
}
