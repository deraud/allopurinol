package entity

type UserInfo struct {
	Id    int64
	Email string
	Name  string
	Role  string
}

type UserInfoOptions struct {
	SearchBy    string
	SearchValue string
	Page        int
	Limit       int
	TotalRows   int
}
