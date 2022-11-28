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

type Categories struct {
	Id   int    `json:"id"`
	Guid string `json:"guid"`
	Name string `json:"name"`
}

type SubCategories struct {
	Id         int    `json:"id"`
	ParentGuid string `json:"parent_guid"`
	Guid       string `json:"guid"`
	Name       string `json:"name"`
}

type Items struct {
	Id          int    `json:"id"`
	CatGuid     string `json:"cat_guid"`
	Guid        string `json:"guid"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Types struct {
	Id         int    `json:"id"`
	ParentGuid string `json:"parent_guid"`
	Guid       string `json:"guid"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	TypePic    string `json:"type_pic"`
}
