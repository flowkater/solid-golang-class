// +build !dynamic
// +build !musl

// This file was auto-generated by librdkafka_vendor/bundle-import.sh, DO NOT EDIT.

package kafka

// #cgo CFLAGS: -DUSE_VENDORED_LIBRDKAFKA -DLIBRDKAFKA_STATICLIB
// #cgo LDFLAGS: ${SRCDIR}/librdkafka_vendor/librdkafka_glibc_linux.a  -lm -ldl -lpthread -lrt
import "C"

// LibrdkafkaLinkInfo explains how librdkafka was linked to the Go client
const LibrdkafkaLinkInfo = "static glibc_linux from librdkafka-static-bundle-v1.8.2.tgz"
