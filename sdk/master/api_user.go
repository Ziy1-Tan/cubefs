package master

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/cubefs/cubefs/proto"
	"github.com/cubefs/cubefs/util/ump"
)

type UserAPI struct {
	mc *MasterClient
}

func (api *UserAPI) CreateUser(param *proto.UserCreateParam, clientIDKey string) (userInfo *proto.UserInfo, err error) {
	request := newAPIRequest(http.MethodPost, proto.UserCreate)
	request.addParam("clientIDKey", clientIDKey)
	var reqBody []byte
	if reqBody, err = json.Marshal(param); err != nil {
		return
	}
	request.addBody(reqBody)
	var data []byte
	if data, err = api.mc.serveRequest(request); err != nil {
		return
	}
	userInfo = &proto.UserInfo{}
	if err = json.Unmarshal(data, userInfo); err != nil {
		return
	}
	return
}

func (api *UserAPI) DeleteUser(userID string, clientIDKey string) (err error) {
	request := newAPIRequest(http.MethodPost, proto.UserDelete)
	request.addParam("user", userID)
	request.addParam("clientIDKey", clientIDKey)
	if _, err = api.mc.serveRequest(request); err != nil {
		return
	}
	return
}

func (api *UserAPI) UpdateUser(param *proto.UserUpdateParam, clientIDKey string) (userInfo *proto.UserInfo, err error) {
	request := newAPIRequest(http.MethodPost, proto.UserUpdate)
	request.addParam("clientIDKey", clientIDKey)
	var reqBody []byte
	if reqBody, err = json.Marshal(param); err != nil {
		return
	}
	request.addBody(reqBody)
	var data []byte
	if data, err = api.mc.serveRequest(request); err != nil {
		return
	}
	userInfo = &proto.UserInfo{}
	if err = json.Unmarshal(data, userInfo); err != nil {
		return
	}
	return
}

func (api *UserAPI) GetAKInfo(accesskey string) (userInfo *proto.UserInfo, err error) {
	localIP, _ := ump.GetLocalIpAddr()
	request := newAPIRequest(http.MethodGet, proto.UserGetAKInfo)
	request.addParam("ak", accesskey)
	request.addParam("ip", localIP)
	var data []byte
	if data, err = api.mc.serveRequest(request); err != nil {
		return
	}
	userInfo = &proto.UserInfo{}
	if err = json.Unmarshal(data, userInfo); err != nil {
		return
	}
	return
}

func (api *UserAPI) AclOperation(volName string, localIP string, op uint32) (aclInfo *proto.AclRsp, err error) {
	request := newAPIRequest(http.MethodGet, proto.AdminACL)
	request.addParam("name", volName)
	request.addParam("ip", localIP)
	request.addParam("op", strconv.Itoa(int(op)))
	var data []byte
	if data, err = api.mc.serveRequest(request); err != nil {
		fmt.Fprintf(os.Stdout, "AclOperation serveRequest err %v\n", err)
		return
	}
	aclInfo = &proto.AclRsp{}
	if err = json.Unmarshal(data, aclInfo); err != nil {
		fmt.Fprintf(os.Stdout, "AclOperation Unmarshal err %v\n", err)
		return
	}

	return
}

func (api *UserAPI) UidOperation(volName string, uid string, op uint32, val string) (uidInfo *proto.UidSpaceRsp, err error) {
	request := newAPIRequest(http.MethodGet, proto.AdminUid)
	request.addParam("name", volName)
	request.addParam("uid", uid)
	request.addParam("op", strconv.Itoa(int(op)))
	request.addParam("capacity", val)
	var data []byte
	if data, err = api.mc.serveRequest(request); err != nil {
		fmt.Fprintf(os.Stdout, "UidOperation serveRequest err %v\n", err)
		return
	}
	uidInfo = &proto.UidSpaceRsp{}
	if err = json.Unmarshal(data, uidInfo); err != nil {
		fmt.Fprintf(os.Stdout, "UidOperation Unmarshal err %v\n", err)
		return
	}

	return
}

func (api *UserAPI) GetUserInfo(userID string) (userInfo *proto.UserInfo, err error) {
	request := newAPIRequest(http.MethodGet, proto.UserGetInfo)
	request.addParam("user", userID)
	var data []byte
	if data, err = api.mc.serveRequest(request); err != nil {
		return
	}
	userInfo = &proto.UserInfo{}
	if err = json.Unmarshal(data, userInfo); err != nil {
		return
	}
	return
}

func (api *UserAPI) UpdatePolicy(param *proto.UserPermUpdateParam, clientIDKey string) (userInfo *proto.UserInfo, err error) {
	request := newAPIRequest(http.MethodPost, proto.UserUpdatePolicy)
	request.addParam("clientIDKey", clientIDKey)
	var reqBody []byte
	if reqBody, err = json.Marshal(param); err != nil {
		return
	}
	request.addBody(reqBody)
	var data []byte
	if data, err = api.mc.serveRequest(request); err != nil {
		return
	}
	userInfo = &proto.UserInfo{}
	if err = json.Unmarshal(data, userInfo); err != nil {
		return
	}
	return
}

func (api *UserAPI) RemovePolicy(param *proto.UserPermRemoveParam, clientIDKey string) (userInfo *proto.UserInfo, err error) {
	request := newAPIRequest(http.MethodPost, proto.UserRemovePolicy)
	request.addParam("clientIDKey", clientIDKey)
	var reqBody []byte
	if reqBody, err = json.Marshal(param); err != nil {
		return
	}
	request.addBody(reqBody)
	var data []byte
	if data, err = api.mc.serveRequest(request); err != nil {
		return
	}
	userInfo = &proto.UserInfo{}
	if err = json.Unmarshal(data, userInfo); err != nil {
		return
	}
	return
}

func (api *UserAPI) DeleteVolPolicy(vol, clientIDKey string) (err error) {
	request := newAPIRequest(http.MethodPost, proto.UserDeleteVolPolicy)
	request.addParam("name", vol)
	request.addParam("clientIDKey", clientIDKey)
	if _, err = api.mc.serveRequest(request); err != nil {
		return
	}
	return
}

func (api *UserAPI) TransferVol(param *proto.UserTransferVolParam, clientIDKey string) (userInfo *proto.UserInfo, err error) {
	request := newAPIRequest(http.MethodPost, proto.UserTransferVol)
	request.addParam("clientIDKey", clientIDKey)
	var reqBody []byte
	if reqBody, err = json.Marshal(param); err != nil {
		return
	}
	request.addBody(reqBody)
	var data []byte
	if data, err = api.mc.serveRequest(request); err != nil {
		return
	}
	userInfo = &proto.UserInfo{}
	if err = json.Unmarshal(data, userInfo); err != nil {
		return
	}
	return
}

func (api *UserAPI) ListUsers(keywords string) (users []*proto.UserInfo, err error) {
	request := newAPIRequest(http.MethodGet, proto.UserList)
	request.addParam("keywords", keywords)
	var data []byte
	if data, err = api.mc.serveRequest(request); err != nil {
		return
	}
	users = make([]*proto.UserInfo, 0)
	if err = json.Unmarshal(data, &users); err != nil {
		return
	}
	return
}

func (api *UserAPI) ListUsersOfVol(vol string) (users []string, err error) {
	request := newAPIRequest(http.MethodGet, proto.UsersOfVol)
	request.addParam("name", vol)
	var data []byte
	if data, err = api.mc.serveRequest(request); err != nil {
		return
	}
	users = make([]string, 0)
	if err = json.Unmarshal(data, &users); err != nil {
		return
	}
	return
}
