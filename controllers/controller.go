package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"uts/models"
)

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var request models.GetAllRoomsRequest
	err := decoder.Decode(&request)
	if err != nil {
		sendErrorResponse(w, 400, "empty request body")
		return
	}

	id_game := request.IdGame
	if id_game == 0 {
		sendErrorResponse(w, 400, "id_game required")
		return
	}

	rows, err := db.Query("SELECT * FROM `rooms` WHERE `id_game`=?", id_game)
	if err != nil {
		sendErrorResponse(w, 500, "internal server error")
		return
	}

	var room models.Room
	var rooms []models.Room
	for rows.Next() {
		if err := rows.Scan(&room.IdRoom, &room.RoomName, &room.IdGame); err != nil {
			sendErrorResponse(w, 500, "internal server error")
			return
		} else {
			rooms = append(rooms, room)
		}
	}

	w.Header().Set("Content-Type", "application-json")
	var response models.GetAllRoomsResponse
	response.Status = 200
	response.Data.Rooms = rooms
	json.NewEncoder(w).Encode(response)
}
func GetRoomDetail(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var request models.GetRoomDetailRequest
	err := decoder.Decode(&request)
	if err != nil {
		sendErrorResponse(w, 400, "empty request body")
		return
	}

	id_room := request.IdRoom
	if id_room == 0 {
		sendErrorResponse(w, 400, "id_room required")
		return
	}

	rows, err := db.Query("SELECT `rooms`.`id` AS `id_room`, `rooms`.`room_name`, `participants`.`id` AS `id_participant`, `accounts`.`id` AS `id_account`, `accounts`.`username` FROM `rooms` JOIN `participants` ON `rooms`.`id` = `participants`.`id_room` JOIN `accounts` ON `participants`.`id_account` = `accounts`.`id` WHERE `id_room`=?;", id_room)
	if err != nil {
		sendErrorResponse(w, 500, "internal server error")
		return
	}

	var room models.RoomDetail
	var participant models.Participant
	var participants []models.Participant
	for rows.Next() {
		if err_query := rows.Scan(&room.Id, &room.RoomName, &participant.Id, &participant.IdAccount, &participant.Username); err_query != nil {
			sendErrorResponse(w, 400, "unable to retrieve room info")
			return
		} else {
			participants = append(participants, participant)
		}
	}

	w.Header().Set("Content-Type", "application-json")
	var response models.GetRoomDetailResponse
	response.Status = 200
	response.Data.Room = room
	response.Data.Room.Participants = participants
	json.NewEncoder(w).Encode(response)
}
func EnterRoom(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var request models.EnterRoomRequest
	err := decoder.Decode(&request)
	if err != nil {
		sendErrorResponse(w, 400, "empty request body")
		return
	}

	id_room := request.IdRoom
	if id_room == 0 {
		sendErrorResponse(w, 400, "id_room required")
		return
	}

	id_account := request.IdAccount
	if id_account == 0 {
		sendErrorResponse(w, 400, "id_account required")
		return
	}

	// check whether user already in a room
	var count int
	_ = db.QueryRow("SELECT COUNT(*) FROM `participants` WHERE `id_account`=?;", id_account).Scan(&count)
	if count > 0 {
		sendErrorResponse(w, 409, "user already in a room")
		return
	}

	// check room capacity
	var account_in_room_count int
	_ = db.QueryRow("SELECT COUNT(*) FROM `participants` WHERE `id_room`=?;", id_room).Scan(&account_in_room_count)
	var max_account_in_room int
	_ = db.QueryRow("SELECT max_player FROM `rooms` WHERE `id`=?;", id_room).Scan(&max_account_in_room)

	if account_in_room_count >= max_account_in_room {
		sendErrorResponse(w, 429, "room full")
		return
	}

	_, err_query := db.Query("INSERT INTO `participants` VALUE (NULL, ?, ?)", id_room, id_account)
	if err_query != nil {
		sendErrorResponse(w, 500, "internal server error")
		return
	}

	sendSuccessResponse(w, "room entered succesfully")
}
func LeaveRoom(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var request models.LeaveRoomRequest
	err := decoder.Decode(&request)
	if err != nil {
		sendErrorResponse(w, 400, "empty request body")
		return
	}

	id_room := request.IdRoom
	if id_room == 0 {
		sendErrorResponse(w, 400, "id_room required")
		return
	}

	id_account := request.IdAccount
	if id_account == 0 {
		sendErrorResponse(w, 400, "id_account required")
		return
	}

	// Check whether user in the room
	var id_participant int
	err_check_participant := db.QueryRow("SELECT id FROM `participants` WHERE `id_room`=? AND `id_account`=?;", id_room, id_account).Scan(&id_participant)

	if err_check_participant == sql.ErrNoRows {
		sendErrorResponse(w, 404, "user not in the room")
		return
	}

	_, err_query := db.Query("DELETE FROM `participants` WHERE `id`=?", id_participant)
	if err_query != nil {
		sendErrorResponse(w, 500, "internal server error")
		return
	}

	sendSuccessResponse(w, "room exited succesfully")
}

func sendSuccessResponse(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application-json")
	var response models.BasicResponse
	response.Status = 200
	response.Message = message
	json.NewEncoder(w).Encode(response)
}

func sendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application-json")
	var response models.ErrorResponse
	response.Status = statusCode
	response.Message = message
	json.NewEncoder(w).Encode(response)
}
