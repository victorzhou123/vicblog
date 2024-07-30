package repositoryimpl

import (
	"victorzhou123/vicblog/common/infrastructure/mysql"
	"victorzhou123/vicblog/user/domain"
	"victorzhou123/vicblog/user/domain/repository"
)

func NewUserRepo(db mysql.Impl) repository.User {
	tableNameUser = db.TableName()

	if err := mysql.AutoMigrate(&UserDO{}); err != nil {
		return nil
	}

	return &userRepoImpl{db}
}

type userRepoImpl struct {
	mysql.Impl
}

func (impl *userRepoImpl) GetUserInfo(account *repository.Account) (domain.User, error) {
	userDo := &UserDO{}
	userDo.Username = account.Username.Username()
	userDo.Password = account.Password.Password()

	if err := impl.GetRecord(&userDo, &userDo); err != nil {
		return domain.User{}, err
	}

	return userDo.toUser()
}
