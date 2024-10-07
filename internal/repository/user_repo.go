package repository

import (
	"database/sql"
	"go-login-crud/model"
	"log"
)

type UserRepository interface {
	GetAllUsers(limit, page int) (data []model.User, err error)
	FindById(id string) (data model.User, err error)
	FindByUsername(username string) (data model.User, err error)
	InsertUser(data model.User) (result string, err error)
	UpdateUser(data model.User) (result string, err error)
	DeleteUser(id string) (result string, err error)
}

type userRepo struct {
	db *sql.DB
}

func CreateUserRepository(db *sql.DB) UserRepository {
	return &userRepo{db: db}
}

func (u *userRepo) GetAllUsers(limit, page int) (data []model.User, err error) {
	query := `SELECT user_id, username, "password", email, phone FROM public.tb_user ORDER BY user_id DESC LIMIT $1 OFFSET ($2-1)*$1;`
	rows, err := u.db.Query(query, limit, page)
	if err != nil {
		log.Fatalf("error query: %s", err)
		return data, err
	}

	for rows.Next() {
		var result model.RespUser
		errScan := rows.Scan(&result.UserId, &result.Username, &result.Password, &result.Email, &result.Phone)
		if errScan != nil {
			log.Fatalf("error query: %s", errScan)
			return data, errScan
		}
		
		respData := result.CreateUserData()
		data = append(data, respData)
	}


	return data, nil
}

func (u *userRepo) FindById(id string) (data model.User, err error) {
	query := `SELECT user_id, username, "password", email, phone FROM public.tb_user WHERE user_id = $1;`

	err = u.db.QueryRow(query, id).Scan(&data)
	if err != nil {
		log.Fatalf("error query: %s", err)
		return data, err
	}

	return data, nil
}

func (u *userRepo) FindByUsername(username string) (data model.User, err error) {
	query := `SELECT user_id, username, "password", email, phone FROM public.tb_user WHERE username = $1;`
	rows, err := u.db.Query(query, username)
	if err != nil {
		log.Fatalf("error query: %s", err)
		return data, err
	}

	for rows.Next() {
		var result model.RespUser
		errScan := rows.Scan(&result.UserId, &result.Username, &result.Password, &result.Email, &result.Phone)
		if errScan != nil {
			log.Fatalf("error query: %s", errScan)
			return data, errScan
		}

		respData := result.CreateUserData()
		data = respData
	}


	return data, nil
}

func (u *userRepo) InsertUser(data model.User) (result string, err error) {
	query := `INSERT INTO public.tb_user (user_id, username, "password", email, phone, created_at, updated_at) VALUES(gen_random_uuid(), $1, $2, $3, $4, now(), now()) RETURNING user_id;`

	err = u.db.QueryRow(query, data.Username, data.Password, data.Email, data.Phone).Scan(&result)
	if err != nil {
		log.Fatalf("error query: %s", err)
		return "", err
	}

	return result, nil
}

func (u *userRepo) UpdateUser(data model.User) (result string, err error) {
	query := `UPDATE public.tb_user SET username=$2, "password"=$3, email=$4, phone=$5, updated_at=now() WHERE user_id::text=$1 RETURNING user_id;`
	
	errdb := u.db.QueryRow(query, data.UserId, data.Username, data.Password, data.Email, data.Phone).Scan(&result)
	if errdb != nil {
		log.Fatalf("error query: %s", errdb)
		return "", errdb
	}

	return result, nil
}

func (u *userRepo) DeleteUser(id string) (result string, err error) {
	query := `DELETE FROM public.tb_user WHERE user_id::text=$1 RETURNING user_id;`

	errdb := u.db.QueryRow(query, id).Scan(&result)
	if errdb != nil {
		log.Fatalf("error query: %s", errdb)
		return "", errdb
	}

	return result, nil
}