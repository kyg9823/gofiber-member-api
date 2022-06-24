package model

import (
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	Id        int32      `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name"`
	Email     string     `json:"email" gorm:"unique"`
	Favorites []Favorite `json:"favorites" gorm:"foreignKey:Id"`
}

type MemberResponse struct {
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type MemberDetailResponse struct {
	Id        int32  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Favorites string `json:"favorites"`
}

type MemberRequest struct {
	Id        int32    `json:"id"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Favorites []string `json:"favorites"`
}

func (member *Member) ConvertFromRequest(memberRequest *MemberRequest) {
	member.Name = memberRequest.Name
	member.Email = memberRequest.Email

	favorites := new([]Favorite)
	for _, favorite := range memberRequest.Favorites {
		*favorites = append(*favorites, Favorite{
			Id:   member.Id,
			Item: favorite,
		})
	}
	member.Favorites = *favorites
}
