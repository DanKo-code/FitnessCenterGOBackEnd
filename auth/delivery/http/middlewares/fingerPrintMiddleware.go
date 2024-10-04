package middlewares

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"strings"
)

func FingerprintMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get User-Agent and Accept headers
		userAgent := c.GetHeader("User-Agent")
		acceptHeaders := c.GetHeader("Accept")

		// Get client IP
		clientIP := c.ClientIP()

		// Concatenate the parameters for fingerprinting
		fingerprintSource := strings.Join([]string{userAgent, acceptHeaders, clientIP}, "")

		// Create a hash of the fingerprint source (MD5 hash here)
		hash := md5.New()
		hash.Write([]byte(fingerprintSource))
		fingerprint := hex.EncodeToString(hash.Sum(nil))

		// Set fingerprint in request context or headers
		c.Set("fingerprint", fingerprint)
		//c.Writer.Header().Set("X-Fingerprint", fingerprint)

		c.Next()
	}
}
