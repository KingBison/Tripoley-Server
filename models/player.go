package models

type Player struct {
	Name      string
	Ready     bool
	Money     int
	Bid       int
	BidStatus string
	MyTurn    bool
	Dealer    bool
}
