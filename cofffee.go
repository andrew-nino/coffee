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
	Id   int    `json:"id,omitempty" db:"id,omitempty"`
	Guid string `json:"guidCategory,omitempty" db:"guid,omitempty"`
	Name string `json:"name,omitempty" db:"name,omitempty"`
}

type SubCategory struct {
	Id         int    `json:"id,omitempty" db:"id,omitempty"`
	ParentGuid string `json:"parent_guid,omitempty" db:"parent_guid,omitempty"`
	Guid       string `json:"guidSubCategory,omitempty" db:"guid,omitempty"`
	Name       string `json:"name,omitempty" db:"name,omitempty"`
}

type Item struct {
	Id            int    `json:"id,omitempty" db:"id,omitempty"`
	CatGuid       string `json:"cat_guid,omitempty" db:"cat_guid,omitempty"`
	SubCatGuid    string `json:"sub_cat_guid,omitempty" db:"sub_cat_guid,omitempty"`
	Guid          string `json:"guidItem,omitempty" db:"guid,omitempty"`
	Name          string `json:"name,omitempty" db:"name,omitempty"`
	Description   string `json:"description,omitempty" db:"description,omitempty"`
	ThimbnailsPic string `json:"thimbnails_pic,omitempty" db:"thimbnails_pic,omitempty"`
}

type Type struct {
	Id         int    `json:"id,omitempty"`
	ParentGuid string `json:"parent_guid,omitempty" db:"parent_guid,omitempty"`
	Guid       string `json:"guidType,omitempty" db:"guid,omitempty"`
	Name       string `json:"name,omitempty" db:"name,omitempty"`
	Price      int    `json:"price,omitempty" db:"price,omitempty"`
	TypePic    string `json:"type_pic,omitempty" db:"type_pic,omitempty"`
}
