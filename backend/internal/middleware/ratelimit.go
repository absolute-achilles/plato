package middleware

import (
	"sync"
	"time"

	"github.com/absolute-achilles/plato/pkg/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// NewLimiter returns a new Limiter that allows events up to rate r and permits
const (
	// Limit defines the maximum frequency of some events.
	// Limit is represented as number of events per second.
	// A zero Limit allows no events.
	RATE_LIMIT rate.Limit = 5

	// bursts of at most b tokens.
	BURST_TOKEN int = 10
)

// client holds the rate limiter and the last time it was seen
type client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	mu      sync.Mutex
	clients = make(map[string]*client)
)

func init() {
	// Background routine to clear idle clients every 3 minute
	go cleanupClients()
}

func cleanupClients() {
	for {
		time.Sleep(2 * time.Minute)
		mu.Lock()
		for ip, cl := range clients {
			// Delete entries inactive for more than 3 minutes
			if time.Since(cl.lastSeen) > 3*time.Minute {
				delete(clients, ip)
			}
		}
		mu.Unlock()
	}
}

func RateLimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		ip := c.ClientIP()

		mu.Lock()
		v, exists := clients[ip]
		if !exists {
			// Allow 'r' requests per second with a burst capacity of 'b'
			v = &client{limiter: rate.NewLimiter(RATE_LIMIT, BURST_TOKEN)}
			clients[ip] = v
		}
		v.lastSeen = time.Now()
		mu.Unlock()

		if !v.limiter.Allow() {
			response.TooManyRequests(c, "message", "Limit exceeded")
			return
		}

		c.Next()
	}
}
