package models

type BasicResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Room struct {
	IdRoom   int    `json:"id_room"`
	RoomName string `json:"room_name"`
	IdGame   int    `json:"id_game"`
}

type Participant struct {
	Id        int    `json:"id"`
	IdAccount int    `json:"id_account"`
	Username  string `json:"username"`
}

type GetAllRoomsRequest struct {
	IdGame int `json:"id_game"`
}

type GetAllRoomsResponse struct {
	Status int             `json:"status"`
	Data   GetAllRoomsItem `json:"data"`
}

type GetAllRoomsItem struct {
	Rooms []Room `json:"rooms"`
}

type GetRoomDetailRequest struct {
	IdRoom int `json:"id_room"`
}

type GetRoomDetailResponse struct {
	Status int               `json:"status"`
	Data   GetRoomDetailItem `json:"data"`
}

type GetRoomDetailItem struct {
	Room RoomDetail `json:"room"`
}

type RoomDetail struct {
	Id           int           `json:"id"`
	RoomName     string        `json:"room_name"`
	Participants []Participant `json:"participants"`
}

type EnterRoomRequest struct {
	IdRoom    int `json:"id_room"`
	IdAccount int `json:"id_account"`
}

type LeaveRoomRequest struct {
	IdRoom    int `json:"id_room"`
	IdAccount int `json:"id_account"`
}
