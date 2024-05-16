package temporal

import (
	"context"
	"fmt"
)

func ComposeGreeting(ctx context.Context, name string) (string, error) {
	greeting := fmt.Sprintf("Hello %s!", name)
	fmt.Printf("Hello %s!", name)
	return greeting, nil
}
func WriteMessage(ctx context.Context, message string) (string, error) {
	fmt.Printf("send message is%v!", message)
	return "", nil
}
func ReadMessage(ctx context.Context, message string) (string, error) {
	fmt.Printf("read message is%v!", message)
	return "", nil
}
