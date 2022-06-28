package model

import (
	"strings"
)

type Member struct {
	// gorm.Model
	Id        string     `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name"`
	Email     string     `json:"email" gorm:"unique"`
	Favorites []Favorite `json:"favorites" gorm:"foreignKey:Id"`
}

type MemberResponse struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type MemberDetail struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Favorites string `json:"favorites"`
}

type MemberDetailResponse struct {
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Favorites []string `json:"favorites"`
}

type MemberRequest struct {
	Id        string   `json:"id"`
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

func (member *MemberDetailResponse) ConvertToResponse(memberDetail *MemberDetail) {
	member.Id = memberDetail.Id
	member.Name = memberDetail.Name
	member.Email = memberDetail.Email
	member.Favorites = strings.Split(memberDetail.Favorites, ",")
}
