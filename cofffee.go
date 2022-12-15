package coffee

type Category struct {
	Id   int    `json:"id,omitempty" db:"id"`
	Guid string `json:"guidCategory,omitempty" db:"guid,omitempty"`
	Name string `json:"name,omitempty" db:"name,omitempty"`
}

type SubCategory struct {
	Id         int    `json:"id,omitempty" db:"id"`
	ParentGuid string `json:"parent_guid,omitempty" db:"parent_guid,omitempty"`
	Guid       string `json:"guidSubCategory,omitempty" db:"guid,omitempty"`
	Name       string `json:"name,omitempty" db:"name,omitempty"`
}

type Item struct {
	Id            int    `json:"id,omitempty" db:"id"`
	CatGuid       string `json:"cat_guid,omitempty" db:"cat_guid,omitempty"`
	SubCatGuid    string `json:"sub_cat_guid,omitempty" db:"sub_cat_guid,omitempty"`
	Guid          string `json:"guidItem,omitempty" db:"guid,omitempty"`
	Name          string `json:"name,omitempty" db:"name,omitempty"`
	Description   string `json:"description,omitempty" db:"description,omitempty"`
	ThimbnailsPic string `json:"thimbnails_pic,omitempty" db:"thimbnails_pic,omitempty"`
	Types         []Type `json:"types,omitempty" db:"types,omitempty"`
}

type Type struct {
	Id         int    `json:"id,omitempty"`
	ParentGuid string `json:"parent_guid,omitempty" db:"parent_guid,omitempty"`
	Guid       string `json:"guidType,omitempty" db:"guid,omitempty"`
	Name       string `json:"name,omitempty" db:"name,omitempty"`
	Price      int    `json:"price,omitempty" db:"price,omitempty"`
	TypePic    string `json:"type_pic,omitempty" db:"type_pic,omitempty"`
}

type Action struct {
	Id               int    `json:"id,omitempty" db:"id"`
	ActionGuid       string `json:"action_guid" db:"guid"`
	ActionName       string `json:"action_name" db:"name"`
	ActionStartDate  string `json:"action_start_date" db:"start_date"`
	ActionExpiryDate string `json:"action_expiry_date,omitempty" db:"expiry_date,omitempty"`
	Description      string `json:"description,omitempty" db:"description,omitempty"`
	Picture          string `json:"action_picture,omitempty" db:"picture,omitempty"`
}