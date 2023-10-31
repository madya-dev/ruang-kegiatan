package repository

import (
	"context"
	"fmt"
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/model"
)

func (r *ReservationRepositoryImpl) InsertTrackReservation(id int) {
	coll := r.MongoDB.Database("ruang-kegiatan").Collection("reservations")
	reservation := dto.ReservationResponse{}
	result := r.DB.Model(&model.Reservation{}).
		Select("reservations.id as id, room_name, activity, start_time, end_time, reservations.study_program as study_program, class_of, document, pic, phone, reservations.updated_at").
		Joins("left join users on users.username = reservations.pic").
		Joins("left join rooms on rooms.id = reservations.room_id").
		Where("reservations.id = ?", id).
		First(&reservation)

	if result.Error != nil {
		fmt.Printf("Error getting data reservation to insert into MongoDB: %v", result.Error)
	} else {
		result, err := coll.InsertOne(context.TODO(), reservation)

		if err != nil {
			fmt.Printf("Error inserting reservation into MongoDB: %v", err)
		} else {
			fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
		}
	}

}
