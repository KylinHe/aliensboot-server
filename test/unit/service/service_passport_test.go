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
	"github.com/KylinHe/aliensboot-server/dispatch/rpc"
	"github.com/KylinHe/aliensboot-server/protocol"
	"testing"
)

func TestUserLogin(t *testing.T) {
	ret, _ := rpc.Passport.UserLoginAuth(1, "", &protocol.UserLogin{
		Username: "test1",
		Password: "test1",
	})
	if ret.Result == protocol.LoginResult_loginSuccess {
		t.Errorf("invalid user login %+v", ret)
	}
}