package mongo

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type UserService struct {
	collection *mgo.Collection
}

func NewUserService(session *mgo.Session, config *root.MongoConfig) *UserService {
	collection := session.DB(config.DbName).C("user")
	collection.EnsureIndex(userModelIndex())
	return &UserService{collection}
}

func (p *UserService) Create(u *root.User) error {
	user := newUserModel(u)
	return p.collection.Insert(&user)
}

func (p *UserService) GetByUsername(username string) (*root.User, error) {
	model := userModel{}
	err := p.collection.Find(bson.M{"username": username}).One(&model)
	return model.toRootUser(), err
}
