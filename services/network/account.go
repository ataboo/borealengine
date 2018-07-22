package network

import (
	"github.com/ataboo/gowssrv/models"
	. "github.com/ataboo/borealengine/config"
	"gopkg.in/mgo.v2"
	"sync"
	"gopkg.in/mgo.v2/bson"
)

var accounts AccountDao

func init() {
	if Config.Environment() != EnvTest {
		accounts = &MgoAccountDAO{}
	} else {
		accounts = &FakeAccountDao{}
	}
}

type AccountDao interface {
	GetByToken(token string) (user models.User, err error)
}

type MgoAccountDAO struct {
	lock sync.Mutex
	acctCol *mgo.Collection
}

func (a *MgoAccountDAO) GetByToken(token string) (user models.User, err error) {
	if a.acctCol == nil {
		err = a.connect()
		if err != nil {
			return user, err
		}
	}

	a.lock.Lock()
	defer a.lock.Unlock()

	err = a.acctCol.Find(bson.M{"session_id": token}).One(&user)
	if err != nil {
		Logger.Debugf("failed to find session with token %s", token)
	}

	return user, err
}

func (a *MgoAccountDAO) connect() error {
	session, err := mgo.Dial(Config.Mongo.AccountServer)
	if err != nil {
		Logger.Errorf("failed to reach account server with: %s", err)
		a.acctCol = nil
		return err
	}

	a.acctCol = session.DB(Config.Mongo.AccountDB).C("users")
	return nil
}

type tokenToUser func(token string) (models.User, error)

type FakeAccountDao struct {
	FakeTTU tokenToUser
}

func (a *FakeAccountDao) GetByToken(token string) (user models.User, err error) {
	return a.FakeTTU(token)
}

