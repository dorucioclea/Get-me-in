package models

type Advert struct {
	Uuid      	string `json:"id"`
	AccountId	string `json:"accountid"`
	Title   	string `json:"title"`
	//Content     Content`json:"content"`
	MaxUsers 	string `json:"maxusers"`
	Premium		bool   `json:"premium"`
	ValidFrom   string `json:"validfrom"`
	ValidTill	string `json:"validtill"`
	Company  	string `json:"company"`
	Description string `json:"description"`
}

/*type Content struct {
	Company  	string `json:"company"`
	Description string `json:"description"`
}*/


