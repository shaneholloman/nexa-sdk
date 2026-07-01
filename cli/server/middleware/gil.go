// Copyright 2024-2026 Qualcomm Technologies, Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-3-Clause

package middleware

import (
	"sync"

	"github.com/gin-gonic/gin"
)

var lock sync.Mutex

func GIL(c *gin.Context) {
	// Block and wait for lock instead of immediately failing
	// This prevents 429 errors when requests queue up briefly
	lock.Lock()
	defer lock.Unlock()

	c.Next()
}
