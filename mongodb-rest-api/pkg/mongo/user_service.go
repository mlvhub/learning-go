package mongo

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/mlvhub/learning-go/mongodb-rest-api/pkg"
)

type UserService struct {
	collection *mgo.Collection
	crypto     root.Crypto
}

func NewUserService(session *mgo.Session, config *root.MongoConfig, crypto root.Crypto) *UserService {
	collection := session.DB(config.DBName).C(config.UserTable)
	collection.EnsureIndex(userModelIndex())
	return &UserService{collection, crypto}
}

func (p *UserService) Create(u *root.User) error {
	user := newUserModel(u)
	hashedPassword, err := p.crypto.Generate(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return p.collection.Insert(&user)
}

func (p *UserService) GetByUsername(username string) (*root.User, error) {
	model := userModel{}
	err := p.collection.Find(bson.M{"username": username}).One(&model)
	return model.toRootUser(), err
}

func (p *UserService) Login(c root.Credentials) (error, root.User) {
	model := userModel{}
	err := p.collection.Find(bson.M{"username": c.Username}).One(&model)

	err = p.crypto.Compare(model.Password, c.Password)
	if err != nil {
		return err, root.User{}
	}

	return err, root.User{
		ID:       model.ID.Hex(),
		Username: model.Username,
		Password: "-"}
}
