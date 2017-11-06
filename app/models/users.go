package models

import (
	"fmt"
	"time"
)

// User struct
type User struct {
	userID        int
	email         string
	name          string
	password      string
	emailVerified string
	createdAt     time.Time
	updatedAt     time.Time
}

/**
Create new User
*/
func (b *User) PostUser() error {
	// var (
	// 	obj interface{}
	// 	err error
	// )

	fmt.Printf("Get post %#v\n", b)
	// obj, err = exe.Get(User{}, b.UserId)
	// if err != nil {
	// 	return fmt.Errorf("Error loading a booking's user (%d): %s", b.UserId, err)
	// }

	// b.User = obj.(*User)

	// obj, err = exe.Get(Hotel{}, b.HotelId)
	// if err != nil {
	// 	return fmt.Errorf("Error loading a booking's hotel (%d): %s", b.HotelId, err)
	// }
	// b.Hotel = obj.(*Hotel)

	// if b.CheckInDate, err = time.Parse(SQL_DATE_FORMAT, b.CheckInStr); err != nil {
	// 	return fmt.Errorf("Error parsing check in date '%s':", b.CheckInStr, err)
	// }
	// if b.CheckOutDate, err = time.Parse(SQL_DATE_FORMAT, b.CheckOutStr); err != nil {
	// 	return fmt.Errorf("Error parsing check out date '%s':", b.CheckOutStr, err)
	// }
	return nil
}
