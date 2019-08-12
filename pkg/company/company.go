package company

type Company struct {
	UID                        int    `csv:"U-ID"`
	Firmierung                 string `csv:"Firmierung"`
	Strasse                    string `csv:"Strasse"`
	PLZ                        string `csv:"Plz"`
	Ort                        string `csv:"Ort"`
	GeschaeftsanschriftStrasse string `csv:"Geschäftsanschrift-Strasse"`
	GeschaeftsanschriftPLZ     string `csv:"Geschäftsanschrift-PLZ"`
	GeschaeftsanschriftOrt     string `csv:"Geschäftsanschrift-Ort"`
}
