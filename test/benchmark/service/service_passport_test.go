/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2019/12/14
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package service

import (
	"github.com/KylinHe/aliensboot-core/common/util"
	"github.com/KylinHe/aliensboot-server/dispatch/rpc"
	"github.com/KylinHe/aliensboot-server/protocol"
	"math/rand"
	"testing"
)

func BenchmarkParallelUserLogin(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			userId := rand.Int63n(10000000)
			req := &protocol.UserLogin{
				Username: "user_" + util.Int64ToString(userId),
				Password: "1111111",
			}
			rpc.Passport.UserLogin("", req)
		}
	})
}