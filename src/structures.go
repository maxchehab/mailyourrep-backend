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
	Kind            string `json:"kind"`
	NormalizedInput struct {
		Line1 string `json:"line1"`
		City  string `json:"city"`
		State string `json:"state"`
		Zip   string `json:"zip"`
	} `json:"normalizedInput"`
	Divisions struct {
		OcdDivisionCountryUs struct {
			Name          string `json:"name"`
			OfficeIndices []int  `json:"officeIndices"`
		} `json:"ocd-division/country:us"`
		OcdDivisionCountryUsStateOr struct {
			Name          string `json:"name"`
			OfficeIndices []int  `json:"officeIndices"`
		} `json:"ocd-division/country:us/state:or"`
		OcdDivisionCountryUsStateOrCd2 struct {
			Name          string `json:"name"`
			OfficeIndices []int  `json:"officeIndices"`
		} `json:"ocd-division/country:us/state:or/cd:2"`
		OcdDivisionCountryUsStateOrCircuitCourt11 struct {
			Name string `json:"name"`
		} `json:"ocd-division/country:us/state:or/circuit_court:11"`
		OcdDivisionCountryUsStateOrCountyDeschutes struct {
			Name          string `json:"name"`
			OfficeIndices []int  `json:"officeIndices"`
		} `json:"ocd-division/country:us/state:or/county:deschutes"`
		OcdDivisionCountryUsStateOrSldl54 struct {
			Name          string `json:"name"`
			OfficeIndices []int  `json:"officeIndices"`
		} `json:"ocd-division/country:us/state:or/sldl:54"`
		OcdDivisionCountryUsStateOrSldu27 struct {
			Name          string `json:"name"`
			OfficeIndices []int  `json:"officeIndices"`
		} `json:"ocd-division/country:us/state:or/sldu:27"`
	} `json:"divisions"`
	Offices []struct {
		Name            string   `json:"name"`
		DivisionID      string   `json:"divisionId"`
		Levels          []string `json:"levels,omitempty"`
		Roles           []string `json:"roles,omitempty"`
		OfficialIndices []int    `json:"officialIndices"`
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
		Party    string   `json:"party"`
		Phones   []string `json:"phones"`
		Urls     []string `json:"urls,omitempty"`
		PhotoURL string   `json:"photoUrl,omitempty"`
		Channels []struct {
			Type string `json:"type"`
			ID   string `json:"id"`
		} `json:"channels,omitempty"`
		Emails []string `json:"emails,omitempty"`
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
