package middlewares

import (
	"context"
	"net/http"
	"your_project/db"
	"your_project/services"
)

func IPBlockerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		clientIP := r.RemoteAddr

		blocker := services.NewIPBlocker(db.GetRedisClient())
		blocked, err := blocker.IsIPBlocked(ctx, clientIP)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if blocked {
			blocker.Logger().WithFields(logrus.Fields{
				"ip":  clientIP,
				"time": time.Now().Format(time.RFC3339),
			}).Warn("Attempted request from blocked IP")
			http.Error(w, "Your IP is temporarily blocked", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
