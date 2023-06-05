package repos

import (
	models "WP/biz/models"
	"WP/pkg/postgres"
	"context"
	"fmt"
)

const (
	FindByIdQuery              = "select id, name, family, age, sex from users where id = $1"
	FindByIdWithInjectionQuery = "select id, name, family, age, sex from users where id = "
	FindByLimitQuery           = "select id, name, family, age, sex from users order by created_at desc limit $1"
)

// TODO: add dialog id to methods when categories finalized

type UsersRepository interface {
	FindById(ctx context.Context, userId int32) (*models.User, error)
	FindByLimit(ctx context.Context, limit int32) ([]*models.User, error)
	FindByIdWithInjection(ctx context.Context, userId string) (*models.User, error)
}

type UsersRepositoryImpl struct {
	masterPg  *postgres.PGXDatabase
	tableName string
}

func NewUsersRepository(masterPg *postgres.PGXDatabase, options ...UsersRepoOption) UsersRepository {

	pr := UsersRepositoryImpl{
		masterPg:  masterPg,
		tableName: "users",
	}

	for _, option := range options {
		pr = option(pr)
	}

	return &pr
}

func (u *UsersRepositoryImpl) FindById(ctx context.Context, userId int32) (*models.User, error) {
	var user models.User
	err := u.masterPg.QueryRow(ctx, FindByIdQuery, userId).Scan(&user.Id, &user.Name, &user.Family, &user.Age, &user.Sex)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UsersRepositoryImpl) FindByIdWithInjection(ctx context.Context, userId string) (*models.User, error) {
	var user models.User
	query := FindByIdWithInjectionQuery + userId
	err := u.masterPg.QueryRow(ctx, query).Scan(&user.Id, &user.Name, &user.Family, &user.Age, &user.Sex)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UsersRepositoryImpl) FindByLimit(ctx context.Context, limit int32) ([]*models.User, error) {
	var res []*models.User

	rows, err := u.masterPg.Query(ctx, FindByLimitQuery, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.Id, &user.Name, &user.Family, &user.Age, &user.Sex)
		if err != nil {
			return nil, err
		}
		res = append(res, &user)
	}
	return res, nil
}

type UsersRepoOption func(UsersRepositoryImpl) UsersRepositoryImpl

func UsersRepoOptionWithTableName(tableName string) UsersRepoOption {
	return func(repo UsersRepositoryImpl) UsersRepositoryImpl {
		if tableName == "" {
			return repo
		}
		repo.tableName = tableName
		return repo
	}
}

func UsersRepoOptionWithAutoCreate() UsersRepoOption {
	return func(repo UsersRepositoryImpl) UsersRepositoryImpl {
		scheme := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s ( 
			name varchar(30) not null,
			family varchar(30) not null,
			id SERIAL PRIMARY KEY,
			age integer not null,
			sex varchar(20) not null,
			created_at timestamp not null default now(),
		);`, repo.tableName)
		_, err := repo.masterPg.Exec(context.Background(), scheme)
		if err != nil {
			panic("create users table failed")
		}
		return repo
	}
}
