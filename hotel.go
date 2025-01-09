package main

import (
	"fmt"
)

type Room struct {
	Number      int
	Type        string
	Price       float64
	IsAvailable bool
}

func main() {
	var hotel []Room
	var choice int

	fmt.Println("Hotel Room Management System")
	fmt.Println("============================")
	fmt.Println("1. Add a Room")
	fmt.Println("2. View All Rooms")
	fmt.Println("3. Book a Room")
	fmt.Println("4. Exit\n")

	for {
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scan(&choice)
		if err != nil || choice < 1 || choice > 4 {
			fmt.Println("Invalid choice. Please enter a number between 1 and 4.\n")
			continue
		}

		switch choice {
		case 1:
			// Add a new room
			var room Room
			fmt.Print("Enter room number: ")
			fmt.Scan(&room.Number)
			fmt.Print("Enter room type (e.g., Single, Double, Suite): ")
			fmt.Scan(&room.Type)
			fmt.Print("Enter room price per night: ")
			fmt.Scan(&room.Price)
			room.IsAvailable = true
			hotel = append(hotel, room)
			fmt.Println("Room added successfully!\n")

		case 2:
			// View all rooms
			if len(hotel) == 0 {
				fmt.Println("No rooms available.\n")
				continue
			}
			fmt.Println("Available Rooms:")
			fmt.Println("================")
			for _, room := range hotel {
				status := "Available"
				if !room.IsAvailable {
					status = "Booked"
				}
				fmt.Printf("Room Number: %d, Type: %s, Price: $%.2f, Status: %s\n", room.Number, room.Type, room.Price, status)
			}
			fmt.Println()

		case 3:
			// Book a room
			var roomNumber int
			fmt.Print("Enter room number to book: ")
			fmt.Scan(&roomNumber)

			roomFound := false
			for i := range hotel {
				if hotel[i].Number == roomNumber {
					roomFound = true
					if hotel[i].IsAvailable {
						hotel[i].IsAvailable = false
						fmt.Printf("Room %d has been successfully booked!\n\n", roomNumber)
					} else {
						fmt.Printf("Room %d is already booked.\n\n", roomNumber)
					}
					break
				}
			}

			if !roomFound {
				fmt.Printf("Room %d does not exist.\n\n", roomNumber)
			}

		case 4:
			// Exit
			fmt.Println("Exiting the system. Goodbye!")
			return
		}
	}
}
