// Copyright 2018 Changkun Ou. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package sched

import (
	"time"
)

// Task interface for sched
type Task interface {
	GetID() (id string)
	SetID(id string)
	IsValidID() bool
	GetExecution() (execute time.Time)
	SetExecution(new time.Time) (old time.Time)
	GetTimeout() (lockTimeout time.Duration)
	GetRetryTime() (execute time.Time)
	Execute() (retry bool, fail error)
}