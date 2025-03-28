package room

import (
	"github.com/Jhon714-cooler/G_shopping/internal/good"
	"github.com/Jhon714-cooler/G_shopping/internal/service/rule"
	"github.com/Jhon714-cooler/G_shopping/internal/service/user"
)



type Room struct {
	Rule rule.Rule
	RoomId int
    Users []*user.User
	//商品id
	Good *good.ProductDetail

}

func NewRoom(rule rule.Rule, roomId int,good *good.ProductDetail ) *Room {
	return &Room{
		Rule: rule,
		RoomId: roomId,
		Good: good,
		Users: []*user.User{},	
	}	
}