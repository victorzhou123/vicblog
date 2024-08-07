package repositoryimpl

import (
	"victorzhou123/vicblog/common/infrastructure/mysql"
	"victorzhou123/vicblog/user/domain"
	"victorzhou123/vicblog/user/domain/repository"
)

func NewUserRepo(db mysql.Impl) repository.User {
	tableNameUser = db.TableName()

	if err := mysql.AutoMigrate(&UserPO{}); err != nil {
		return nil
	}

	return &userRepoImpl{db}
}

type userRepoImpl struct {
	mysql.Impl
}

func (impl *userRepoImpl) GetUserInfo(account *repository.Account) (domain.User, error) {
	po := &UserPO{}
	po.Username = account.Username.Username()
	po.Password = account.Password.Password()

	if err := impl.GetRecord(&po, &po); err != nil {
		return domain.User{}, err
	}

	return po.toUser()
}
