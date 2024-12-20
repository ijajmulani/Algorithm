// The Observer Pattern is a behavioral design pattern used to establish a one-to-many relationship between objects.
// When one object (the subject) changes state, all its dependent objects (observers) are notified and updated automatically.
// This pattern is particularly useful in scenarios where state changes in one object should trigger updates in other objects, without tightly coupling them.

package main

import "fmt"

// Observer interface
type Subscriber interface {
	Notification(string, string)
}

type User struct {
	Name string
}

func (u *User) Notification(blogName string, title string) {
	fmt.Printf("%s got notification from %s blog, Post title: %s \n", u.Name, blogName, title)
}

type BlogChannel struct {
	Name        string
	subscribers []Subscriber
}

func (b *BlogChannel) AddSubscriber(subscriber Subscriber) {
	b.subscribers = append(b.subscribers, subscriber)
}

func (b *BlogChannel) RemoveSubscriber(sub Subscriber) {
	for i, s := range b.subscribers {
		if sub == s {
			b.subscribers = append(b.subscribers[:i], b.subscribers[i+1:]...)
			break
		}
	}
}

func (b *BlogChannel) Post(postTitle string) {
	for _, s := range b.subscribers {
		s.Notification(b.Name, postTitle)
	}
}

func main() {
	user1 := &User{Name: "Ijaj"}
	user2 := &User{Name: "Arhaan"}
	b := &BlogChannel{Name: "Tech_Blog"}
	b.AddSubscriber(user1)
	b.AddSubscriber(user2)
	b.Post("Heap Algorithm Explanation")
}
