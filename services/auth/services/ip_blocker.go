package services

import (
	"context"
	"encoding/json"
	"net"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type IPBlocker struct {
	client *redis.Client
	logger *logrus.Logger
}


func NewIPBlocker(client *redis.Client) *IPBlocker {
	logger := logrus.New()


	conn, err := net.Dial("tcp", "logstash:5000")
	if err == nil {
		logger.SetOutput(conn)
	} else {
		logger.Warn("Failed to connect to Logstash, logging to stdout")
	}

	logger.SetFormatter(&logrus.JSONFormatter{})
	return &IPBlocker{client: client, logger: logger}
}


func (ib *IPBlocker) BlockIP(ctx context.Context, ip string, duration time.Duration) error {
	err := ib.client.Set(ctx, "blocked_ip:"+ip, "blocked", duration).Err()
	if err == nil {
		logData := map[string]interface{}{
			"ip":         ip,
			"blockedAt":  time.Now().Format(time.RFC3339),
			"duration":   duration.String(),
			"event":      "IP Blocked",
			"timestamp":  time.Now().Format(time.RFC3339),
		}

		jsonLog, _ := json.Marshal(logData)
		ib.logger.Info(string(jsonLog))
	}
	return err
}


func (ib *IPBlocker) IsIPBlocked(ctx context.Context, ip string) (bool, error) {
	exists, err := ib.client.Exists(ctx, "blocked_ip:"+ip).Result()
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}
