package model

type Session struct {
	SessionID string
	UserName  string
	UserID    int
	Order     *Order
	Orders    []*Order
}