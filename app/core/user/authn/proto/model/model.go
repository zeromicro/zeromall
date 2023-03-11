package model

import "user/authn/rpc/pb"

func (m *UserAuthn) From(in *pb.UserRegisterReq) *UserAuthn {
	m.AuthType = int64(in.RegisterType)
	if in.RegisterType == pb.AuthType_Mobile {
		m.AuthUnique = in.MobileNo + in.MobileCountry
		m.AuthField1 = in.MobileNo
		m.AuthField2 = in.MobileCountry
	} else if in.RegisterType == pb.AuthType_Email {
		m.AuthUnique = in.Email
		m.AuthField1 = in.Email
	} else if in.RegisterType == pb.AuthType_Username {
		m.AuthUnique = in.Username
		m.AuthField1 = in.Username
	}
	return m
}
