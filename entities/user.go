package entities

type User struct {
	Id string `json,gorethink:"id"`
	Username string `json,gorethink:"username" form:"username"`
	Password string `form:"password"`
	Email string `yml:"email" form:"email"`
	Image string `yml:"Image" form:"image"`
	Role Role `json,gorethink:"role"`
	Token Token
	Cart Cart `json,gorethink:"cart"`
	PastCarts []Cart `json,gorethink:"past_carts"`
}

type Token struct {
	Token string `json,gorethink:"token"`
	Expire int64 `json,gorethink:"expire"`
	RefreshToken string `json,gorethink:"refresh_token"`
}

type Cart struct {
	Items []Item `json,gorethink:"items"`
}

type PostItem struct {
	ItemId string `json,gorethink:"item_id"`
}
