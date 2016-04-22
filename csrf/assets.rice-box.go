package main

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    `balance.html`,
		FileModTime: time.Unix(1459543356, 0),
		Content:     string([]byte{0x3c, 0x21, 0x64, 0x6f, 0x63, 0x74, 0x79, 0x70, 0x65, 0x20, 0x68, 0x74, 0x6d, 0x6c, 0x3e, 0xa, 0x3c, 0x21, 0x2d, 0x2d, 0xa, 0x43, 0x6f, 0x70, 0x79, 0x72, 0x69, 0x67, 0x68, 0x74, 0x20, 0x32, 0x30, 0x31, 0x36, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x20, 0x49, 0x6e, 0x63, 0x2e, 0x20, 0x41, 0x6c, 0x6c, 0x20, 0x52, 0x69, 0x67, 0x68, 0x74, 0x73, 0x20, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x2e, 0xa, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x64, 0x20, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x41, 0x70, 0x61, 0x63, 0x68, 0x65, 0x20, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2c, 0x20, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x32, 0x2e, 0x30, 0x20, 0x28, 0x74, 0x68, 0x65, 0x20, 0x22, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x3b, 0xa, 0x79, 0x6f, 0x75, 0x20, 0x6d, 0x61, 0x79, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x75, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x66, 0x69, 0x6c, 0x65, 0x20, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x20, 0x69, 0x6e, 0x20, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x74, 0x68, 0x65, 0x20, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0xa, 0x59, 0x6f, 0x75, 0x20, 0x6d, 0x61, 0x79, 0x20, 0x6f, 0x62, 0x74, 0x61, 0x69, 0x6e, 0x20, 0x61, 0x20, 0x63, 0x6f, 0x70, 0x79, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x20, 0x61, 0x74, 0xa, 0x20, 0x20, 0x20, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x2f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x2d, 0x32, 0x2e, 0x30, 0xa, 0x55, 0x6e, 0x6c, 0x65, 0x73, 0x73, 0x20, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x6c, 0x61, 0x77, 0x20, 0x6f, 0x72, 0x20, 0x61, 0x67, 0x72, 0x65, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x69, 0x6e, 0x20, 0x77, 0x72, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x2c, 0x20, 0x73, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0xa, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x20, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x20, 0x69, 0x73, 0x20, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x20, 0x6f, 0x6e, 0x20, 0x61, 0x6e, 0x20, 0x22, 0x41, 0x53, 0x20, 0x49, 0x53, 0x22, 0x20, 0x42, 0x41, 0x53, 0x49, 0x53, 0x2c, 0xa, 0x57, 0x49, 0x54, 0x48, 0x4f, 0x55, 0x54, 0x20, 0x57, 0x41, 0x52, 0x52, 0x41, 0x4e, 0x54, 0x49, 0x45, 0x53, 0x20, 0x4f, 0x52, 0x20, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x20, 0x4f, 0x46, 0x20, 0x41, 0x4e, 0x59, 0x20, 0x4b, 0x49, 0x4e, 0x44, 0x2c, 0x20, 0x65, 0x69, 0x74, 0x68, 0x65, 0x72, 0x20, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x20, 0x6f, 0x72, 0x20, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x2e, 0xa, 0x53, 0x65, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x20, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x20, 0x67, 0x6f, 0x76, 0x65, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x20, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x20, 0x61, 0x6e, 0x64, 0xa, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x20, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0xa, 0x2d, 0x2d, 0x3e, 0xa, 0x3c, 0x68, 0x74, 0x6d, 0x6c, 0x3e, 0xa, 0x3c, 0x62, 0x6f, 0x64, 0x79, 0x3e, 0xa, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x74, 0x72, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x74, 0x68, 0x3e, 0x44, 0x61, 0x74, 0x65, 0x3c, 0x2f, 0x74, 0x68, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x74, 0x68, 0x3e, 0x52, 0x65, 0x63, 0x65, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x3c, 0x2f, 0x74, 0x68, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x74, 0x68, 0x3e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x3c, 0x2f, 0x74, 0x68, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x74, 0x72, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7b, 0x7b, 0x20, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x20, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x20, 0x7d, 0x7d, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x74, 0x72, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x74, 0x64, 0x3e, 0x7b, 0x7b, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x7d, 0x7d, 0x3c, 0x2f, 0x74, 0x64, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x74, 0x64, 0x3e, 0x7b, 0x7b, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x7d, 0x7d, 0x3c, 0x2f, 0x74, 0x64, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x74, 0x64, 0x3e, 0x7b, 0x7b, 0x2e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x7d, 0x7d, 0x3c, 0x2f, 0x74, 0x64, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x74, 0x72, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7b, 0x7b, 0x20, 0x65, 0x6e, 0x64, 0x20, 0x7d, 0x7d, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x3e, 0xa, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x20, 0x20, 0x3c, 0x66, 0x6f, 0x72, 0x6d, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x22, 0x20, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3d, 0x22, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x22, 0x20, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x3d, 0x22, 0x50, 0x4f, 0x53, 0x54, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x68, 0x33, 0x3e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x73, 0x6f, 0x6d, 0x65, 0x20, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x3c, 0x2f, 0x68, 0x33, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x3a, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x22, 0x74, 0x65, 0x78, 0x74, 0x22, 0x20, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x3d, 0x22, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x22, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x22, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x22, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x22, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3d, 0x22, 0x53, 0x65, 0x6e, 0x64, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x3c, 0x2f, 0x66, 0x6f, 0x72, 0x6d, 0x3e, 0xa, 0x20, 0x20, 0x3c, 0x61, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x22, 0x2f, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x22, 0x3e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x3c, 0x2f, 0x61, 0x3e, 0xa, 0x3c, 0x2f, 0x62, 0x6f, 0x64, 0x79, 0x3e, 0xa, 0x3c, 0x2f, 0x68, 0x74, 0x6d, 0x6c, 0x3e, 0xa}), //++ TODO: optimize? (double allocation) or does compiler already optimize this?
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    `login.html`,
		FileModTime: time.Unix(1459543359, 0),
		Content:     string([]byte{0x3c, 0x21, 0x64, 0x6f, 0x63, 0x74, 0x79, 0x70, 0x65, 0x20, 0x68, 0x74, 0x6d, 0x6c, 0x3e, 0xa, 0x3c, 0x21, 0x2d, 0x2d, 0xa, 0x43, 0x6f, 0x70, 0x79, 0x72, 0x69, 0x67, 0x68, 0x74, 0x20, 0x32, 0x30, 0x31, 0x36, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x20, 0x49, 0x6e, 0x63, 0x2e, 0x20, 0x41, 0x6c, 0x6c, 0x20, 0x52, 0x69, 0x67, 0x68, 0x74, 0x73, 0x20, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x2e, 0xa, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x64, 0x20, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x41, 0x70, 0x61, 0x63, 0x68, 0x65, 0x20, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2c, 0x20, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x32, 0x2e, 0x30, 0x20, 0x28, 0x74, 0x68, 0x65, 0x20, 0x22, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x3b, 0xa, 0x79, 0x6f, 0x75, 0x20, 0x6d, 0x61, 0x79, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x75, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x66, 0x69, 0x6c, 0x65, 0x20, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x20, 0x69, 0x6e, 0x20, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x74, 0x68, 0x65, 0x20, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0xa, 0x59, 0x6f, 0x75, 0x20, 0x6d, 0x61, 0x79, 0x20, 0x6f, 0x62, 0x74, 0x61, 0x69, 0x6e, 0x20, 0x61, 0x20, 0x63, 0x6f, 0x70, 0x79, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x20, 0x61, 0x74, 0xa, 0x20, 0x20, 0x20, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x2f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x2d, 0x32, 0x2e, 0x30, 0xa, 0x55, 0x6e, 0x6c, 0x65, 0x73, 0x73, 0x20, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x6c, 0x61, 0x77, 0x20, 0x6f, 0x72, 0x20, 0x61, 0x67, 0x72, 0x65, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x69, 0x6e, 0x20, 0x77, 0x72, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x2c, 0x20, 0x73, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0xa, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x20, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x20, 0x69, 0x73, 0x20, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x20, 0x6f, 0x6e, 0x20, 0x61, 0x6e, 0x20, 0x22, 0x41, 0x53, 0x20, 0x49, 0x53, 0x22, 0x20, 0x42, 0x41, 0x53, 0x49, 0x53, 0x2c, 0xa, 0x57, 0x49, 0x54, 0x48, 0x4f, 0x55, 0x54, 0x20, 0x57, 0x41, 0x52, 0x52, 0x41, 0x4e, 0x54, 0x49, 0x45, 0x53, 0x20, 0x4f, 0x52, 0x20, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x20, 0x4f, 0x46, 0x20, 0x41, 0x4e, 0x59, 0x20, 0x4b, 0x49, 0x4e, 0x44, 0x2c, 0x20, 0x65, 0x69, 0x74, 0x68, 0x65, 0x72, 0x20, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x20, 0x6f, 0x72, 0x20, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x2e, 0xa, 0x53, 0x65, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x20, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x20, 0x67, 0x6f, 0x76, 0x65, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x20, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x20, 0x61, 0x6e, 0x64, 0xa, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x20, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0xa, 0x2d, 0x2d, 0x3e, 0xa, 0x3c, 0x68, 0x74, 0x6d, 0x6c, 0x3e, 0xa, 0x3c, 0x62, 0x6f, 0x64, 0x79, 0x3e, 0xa, 0x20, 0x20, 0x3c, 0x66, 0x6f, 0x72, 0x6d, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x22, 0x20, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3d, 0x22, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x22, 0x20, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x3d, 0x22, 0x50, 0x4f, 0x53, 0x54, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x68, 0x33, 0x3e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x3c, 0x2f, 0x68, 0x33, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x3a, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x22, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x22, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x22, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3d, 0x22, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x3c, 0x2f, 0x66, 0x6f, 0x72, 0x6d, 0x3e, 0xa, 0x3c, 0x2f, 0x62, 0x6f, 0x64, 0x79, 0x3e, 0xa, 0x3c, 0x2f, 0x68, 0x74, 0x6d, 0x6c, 0x3e, 0xa}), //++ TODO: optimize? (double allocation) or does compiler already optimize this?
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   ``,
		DirModTime: time.Unix(1459336030, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // balance.html
			file3, // login.html

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`assets`, &embedded.EmbeddedBox{
		Name: `assets`,
		Time: time.Unix(1460977598, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"balance.html": file2,
			"login.html":   file3,
		},
	})
}
