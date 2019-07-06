package main

import (
	"context"
	"fmt"
	"time"
)

type User struct {
	Name  string
	Pass  string
	Count uint8
}

func userContext(ctx context.Context, username string) context.Context {
	myUser := User{username, "123456", 0}
	return context.WithValue(ctx, myUser.Name, &myUser)
}

func userRotine(ctx context.Context, username string) {
	myUser := ctx.Value(username).(*User)

	for {
		myUser.Count++
		fmt.Printf("for loop: user %s %d \n", myUser.Name, myUser.Count)
		select {
		case <-ctx.Done():
			fmt.Printf("  User %s is DONE! \n", myUser.Name)
			return
		default:
			fmt.Printf("  still working on user %s\n", myUser.Name)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

var usernameChannel = make(chan string)

func Completed(ctx context.Context) {
	username := <-usernameChannel
	usr := ctx.Value(username).(*User)
	fmt.Println(usr)
	<-ctx.Done()
	fmt.Println(usr)
}

func main() {
	username := "boby"
	usernameChannel <- username

	fmt.Println(<-usernameChannel)
	fmt.Println(<-usernameChannel)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	time.AfterFunc(3*time.Second, cancel)
	// defer cancel()

	ctx = userContext(ctx, username)
	go userRotine(ctx, username)

	go Completed(ctx)
	go Completed(ctx)
	go Completed(ctx)

	// <-ctx.Done()
	fmt.Printf("Err(): %s\n", ctx.Err())
	// fmt.Printf("Done(): %s\n", <-ctx.Done())
	// fmt.Printf("Err(): %s\n", ctx.Err())
	// fmt.Println("Finish!")
}
