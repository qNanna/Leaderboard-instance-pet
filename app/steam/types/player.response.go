package types

type SteamPlayersResponse struct {
	Response struct {
		Players []SteamPlayer `json:"players"`
	} `json:"response"`
}

type SteamPlayer struct {
	SteamID             string `json:"steamid"`
	CommunityVisibility int    `json:"communityvisibilitystate"`
	ProfileState        int    `json:"profilestate"`
	PersonaName         string `json:"personaname"`
	CommentPermission   int    `json:"commentpermission,omitempty"` // Add omitempty tag
	ProfileURL          string `json:"profileurl"`
	Avatar              string `json:"avatar"`
	AvatarMedium        string `json:"avatarmedium"`
	AvatarFull          string `json:"avatarfull"`
	AvatarHash          string `json:"avatarhash"`
	LastLogOff          int64  `json:"lastlogoff"` // Change to int64
	PersonaState        int    `json:"personastate"`
	PrimaryClanID       string `json:"primaryclanid"`
	TimeCreated         int64  `json:"timecreated"` // Change to int64
	PersonaStateFlags   int    `json:"personastateflags"`
}
