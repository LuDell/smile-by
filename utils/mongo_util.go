package utils

import (
	"gopkg.in/mgo.v2"
	"time"
)

var Session *mgo.Session

func init()  {
	Session = globalSession()
}

func globalSession() *mgo.Session {
	dialInfo := &mgo.DialInfo{
		Addrs:     []string{"45.77.82.85:27017"},
		Direct:    false,
		Timeout:   time.Second * 1,
		Username:  "admin",
		Password:  "1q2w3e4r",
		PoolLimit: 4096, // Session.SetPoolLimit
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if nil != err {
		panic(err)
	}
	session.SetMode(mgo.Monotonic,true)
	return session
}

func ShowAdminDB() *mgo.Database {
	session := Session.Copy();
	return session.DB("admin")
}

func ShowTestDB() *mgo.Database {
	session := Session.Copy();
	return session.DB("test")
}
