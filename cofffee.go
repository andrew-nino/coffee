package coffee

type CoffeeList struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type CoffeeItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}

type Category struct {
	Id   int    `json:"id" db:"id"`
	Guid string `json:"guidCategory,omitempty" db:"guid,omitempty"`
	Name string `json:"name,omitempty" db:"name,omitempty"`
}

type SubCategory struct {
	Id         int    `json:"id" db:"id"`
	ParentGuid string `json:"parent_guid,omitempty" db:"parent_guid,omitempty"`
	Guid       string `json:"guidSubCategory,omitempty" db:"guid,omitempty"`
	Name       string `json:"name,omitempty" db:"name,omitempty"`
}

type Item struct {
	Id            int    `json:"id db:"id"`
	CatGuid       string `json:"cat_guid,omitempty" db:"cat_guid,omitempty"`
	SubCatGuid    string `json:"sub_cat_guid,omitempty" db:"sub_cat_guid,omitempty"`
	Guid          string `json:"guidItem,omitempty" db:"guid,omitempty"`
	Name          string `json:"name,omitempty" db:"name,omitempty"`
	Description   string `json:"description,omitempty" db:"description,omitempty"`
	ThimbnailsPic string `json:"thimbnails_pic,omitempty" db:"thimbnails_pic,omitempty"`
	Types         []Type `json:"types,omitempty" db:"types,omitempty"`
}

type Type struct {
	Id         int    `json:"id"`
	ParentGuid string `json:"parent_guid,omitempty" db:"parent_guid,omitempty"`
	Guid       string `json:"guidType,omitempty" db:"guid,omitempty"`
	Name       string `json:"name,omitempty" db:"name,omitempty"`
	Price      int    `json:"price,omitempty" db:"price,omitempty"`
	TypePic    string `json:"type_pic,omitempty" db:"type_pic,omitempty"`
}

type Action struct {
	Id               int    `json:"id" db:"id"`
	ActionGuid       string `json:"action_guid" db:"action_guid"`
	ActionName       string `json:"action_name" db:"action_name"`
	ActionStartDate  string `json:"action_start_date" db:"action_start_date"`
	ActionExpiryDate string `json:"action_expiry_date,omitempty" db:"action_expiry_date,omitempty"`
	Description      string `json:"description,omitempty" db:"description,omitempty"`
	Picture          string `json:"action_picture,omitempty" db:"action_picture,omitempty"`
}