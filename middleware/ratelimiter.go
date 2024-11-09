package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type clientInfo struct {
	failures int
	banTime  time.Time
}

// In-memory storage for client attempts
var clientStore = struct {
	sync.RWMutex
	clients map[string]*clientInfo
}{clients: make(map[string]*clientInfo)}

// limit rate for prevent from brute attack
func RateLimiterMiddleware(limit int, window time.Duration, banDuration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()

		clientStore.Lock()
		client, exists := clientStore.clients[clientIP]
		if !exists {
			client = &clientInfo{failures: 0, banTime: time.Time{}}
			clientStore.clients[clientIP] = client
		}
		clientStore.Unlock()

		if time.Now().Before(client.banTime) {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()

		if c.Writer.Status() == http.StatusUnauthorized {
			clientStore.Lock()
			client.failures++
			if client.failures > limit {
				client.banTime = time.Now().Add(banDuration)
			}
			clientStore.Unlock()
		}

		// Reset the count after the window
		go func() {
			time.Sleep(window)
			clientStore.Lock()
			client.failures = 0
			clientStore.Unlock()
		}()
	}
}
