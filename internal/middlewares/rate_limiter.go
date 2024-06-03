package middlewares

import (
	"net"
	"net/http"

	"github.com/Survialander/rate-limitter/configs"
	infra "github.com/Survialander/rate-limitter/internal/infra/providers"
	"github.com/Survialander/rate-limitter/internal/service"
)

func RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		redisprovider := infra.GetLimitterRedisProvider()
		limiterService := service.GetLimitterService(configs.GetConfig(), redisprovider)

		var blockRequest bool
		token := r.Header.Get("API_KEY")

		if token != "" {
			blockRequest = limiterService.CheckRequestLimit(token)
		} else {
			ip, _, _ := net.SplitHostPort(r.RemoteAddr)
			blockRequest = limiterService.CheckRequestLimit(ip)
		}

		if blockRequest {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("you have reached the maximum number of requests or actions allowed within a certain time frame"))
			return
		}

		next.ServeHTTP(w, r)
	})
}
