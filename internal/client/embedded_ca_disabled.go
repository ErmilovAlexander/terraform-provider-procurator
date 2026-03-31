//go:build !embedded_ca

package client

// Если binary собран без тега embedded_ca,
// встроенного сертификата нет.
var embeddedRootCA []byte
