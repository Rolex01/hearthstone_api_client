package models

type Cards struct {
	Cards     []Card `json:"cards"`
	CardCount int    `json:"cardCount"`
	PageCount int    `json:"pageCount"`
	Page      int    `json:"page"`
}

type Card struct {
	Id            int           `json:"id"`
	Collectible   int           `json:"collectible"`
	Slug          string        `json:"slug"`
	ClassId       int           `json:"classId"`
	MultiClassIds []int         `json:"multiClassIds"`
	CardTypeId    int           `json:"cardTypeId"`
	CardSetId     int           `json:"cardSetId"`
	RarityId      int           `json:"rarityId"`
	ArtistName    string        `json:"artistName"`
	Health        string        `json:"health"`
	ManaCost      int           `json:"manaCost"`
	Name          string        `json:"name"`
	Text          string        `json:"text"`
	Image         string        `json:"image"`
	ImageGold     string        `json:"imageGold"`
	FlavorText    string        `json:"flavorText"`
	CropImage     string        `json:"cropImage"`
	ChildIds      []int         `json:"childIds"`
	KeywordIds    []int         `json:"keywordIds"`
	Battlegrounds Battlegrounds `json:"battlegrounds"`
}

type Battlegrounds struct {
	Tier      int    `json:"tier"`
	Hero      bool   `json:"childIds"`
	UpgradeId int    `json:"upgradeId"`
	Image     string `json:"image"`
	ImageGold string `json:"imageGold"`
}
