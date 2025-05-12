// Package testcert contains a test-only localhost certificate.
package testcert

import "strings"

// LocalhostCert is a PEM-encoded TLS cert with SAN IPs
// "127.0.0.1" and "[::1]", expiring at Jan 29 16:00:00 2084 GMT.
// Removed for clean Traefik build.
var LocalhostCert = []byte("")

// LocalhostKey is the private key for LocalhostCert.
// Removed for clean Traefik build.
var LocalhostKey = []byte("")

// IsLocalhost returns true if host is "localhost", "127.0.0.1" or "::1".
func IsLocalhost(host string) bool {
    host = strings.ToLower(host)
    return host == "localhost" || host == "127.0.0.1" || host == "::1"
}
