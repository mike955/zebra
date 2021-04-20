package data

type CreateRequest struct {
	Username  string
	Level     uint64
	QQ        string
	Wechat    string
	Cellphone string
	Email     string
}

type CreateResponse struct {
}

type DeletesRequest struct {
	Ids []uint64
}

type DeletesResponse struct {
}

type UpdateRequest struct {
	Id uint64
}

type UpdateResponse struct {
}

type GetsRequest struct {
	Ids       []uint64
	Username  string
	Level     uint64
	QQ        string
	Wechat    string
	Cellphone string
	Email     string
}

type GetsResponse struct {
}

type AuthRequest struct {
}

type AuthResponse struct {
}
