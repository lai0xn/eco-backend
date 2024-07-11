package sse

import (
	"context"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/lai0xn/squid-tech/pkg/redis"
	r "github.com/redis/go-redis/v9"
)

type Notifier struct {
	client *r.Client
}

func NewNotifier() *Notifier {
	return &Notifier{
		client: redis.GetClient(),
	}
}

func (n *Notifier) NotificationHandler(c echo.Context) error {

	c.Response().Header().Set("Content-Type", "text/event-stream")
	c.Response().Header().Set("Cache-Control", "no-cache")
	c.Response().Header().Set("Connection", "keep-alive")
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	id := c.QueryParam("userID")

	// Ensure the Content-Length is not set, as SSE is a streaming response
	c.Response().Header().Del("Content-Length")

	// Create a new Redis Pub/Sub subscriber
	pubsub := redis.GetClient().Subscribe(context.Background(), "notifs:"+id)
	defer pubsub.Close()

	_, err := pubsub.Receive(context.Background())
	if err != nil {
		log.Printf("Failed to subscribe: %v", err)
		return err
	}

	// Start a goroutine to receive messages from Redis
	ch := pubsub.Channel()

	for {
		select {
		case msg := <-ch:
			fmt.Fprintf(c.Response(), "data: %s\n\n", msg.Payload)
			c.Response().Flush()
		case <-c.Request().Context().Done():
			return nil
		}
	}
}
