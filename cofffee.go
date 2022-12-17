package coffee

type Category struct {
	Id   int    `json:"id,omitempty" db:"id"`
	Guid string `json:"guidCategory" db:"guid"`
	Name string `json:"name" db:"name"`
}

type SubCategory struct {
	Id         int    `json:"id,omitempty" db:"id"`
	ParentGuid string `json:"parent_guid" db:"parent_guid"`
	Guid       string `json:"guidSubCategory" db:"guid"`
	Name       string `json:"name" db:"name"`
}

type Item struct {
	Id            int    `json:"id,omitempty" db:"id"`
	CatGuid       string `json:"cat_guid" db:"cat_guid"`
	SubCatGuid    string `json:"sub_cat_guid" db:"sub_cat_guid"`
	Guid          string `json:"guidItem" db:"guid"`
	Name          string `json:"name" db:"name"`
	Description   string `json:"description" db:"description"`
	ThimbnailsPic string `json:"thimbnails_pic" db:"thimbnails_pic"`
	Types         []Type `json:"types" db:"types"`
}

type Type struct {
	Id         int    `json:"id,omitempty"`
	ParentGuid string `json:"parent_guid" db:"parent_guid"`
	Guid       string `json:"guidType" db:"guid"`
	Name       string `json:"name" db:"name"`
	Price      int    `json:"price" db:"price"`
	TypePic    string `json:"type_pic" db:"type_pic"`
}

type Action struct {
	Id               int    `json:"id,omitempty" db:"id"`
	ActionGuid       string `json:"action_guid" db:"guid"`
	ActionName       string `json:"action_name" db:"name"`
	ActionStartDate  string `json:"action_start_date" db:"start_date"`
	ActionExpiryDate string `json:"action_expiry_date" db:"expiry_date"`
	Description      string `json:"description" db:"description"`
	Picture          string `json:"action_picture" db:"picture"`
}
