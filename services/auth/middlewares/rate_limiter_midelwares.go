package middlewares

import (
	"context"
	"net/http"
	"time"
	"your_project/db"
	"your_project/services"
	"github.com/sirupsen/logrus"
)

func RateLimiterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		clientIP := r.RemoteAddr

		limiter := services.NewRateLimiter(db.GetRedisClient(), clientIP, 10, 0.5)
		blocker := services.NewIPBlocker(db.GetRedisClient())

		allowed, err := limiter.AllowRequest(ctx)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if !allowed {
			blocker.BlockIP(ctx, clientIP, time.Hour)
			logrus.WithFields(logrus.Fields{
				"ip":  clientIP,
				"time": time.Now().Format(time.RFC3339),
			}).Warn("IP blocked due to excessive requests")

			http.Error(w, "Too Many Requests - Your IP is temporarily blocked", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
