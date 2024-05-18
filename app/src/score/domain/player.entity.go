package entity

import "time"

type SteamPlayer struct {
	SteamID             string    `json:"steamid"`
	CommunityVisibility int       `json:"communityvisibilitystate"`
	ProfileState        int       `json:"profilestate"`
	PersonaName         string    `json:"personaname"`
	CommentPermission   int       `json:"commentpermission"`
	ProfileURL          string    `json:"profileurl"`
	Avatar              string    `json:"avatar"`
	AvatarMedium        string    `json:"avatarmedium"`
	AvatarFull          string    `json:"avatarfull"`
	AvatarHash          string    `json:"avatarhash"`
	LastLogOff          time.Time `json:"lastlogoff"`
	PersonaState        int       `json:"personastate"`
	PrimaryClanID       string    `json:"primaryclanid"`
	TimeCreated         time.Time `json:"timecreated"`
	PersonaStateFlags   int       `json:"personastateflags"`
}
