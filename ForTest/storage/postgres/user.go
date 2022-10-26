package postgres

import (
	"database/sql"
	"fmt"
	pb "github/Services/ForTest/genproto/user"

	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

//NewUserRepo ...
func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *pb.User) (*pb.User, error) {

	query := `INSERT INTO testuser (
			  id, 
			  name, 
			  username, 
			  city, 
			  created_at )
    VALUES($1, $2, $3, $4, $5) RETURNING id`

	//this is generatinig new uuid
	id := uuid.New().String()
	user.CreatedAt = time.Now().Format("2006-01-02")

	//this is writing for database
	err := r.db.QueryRow(query,
		id,
		user.Name,
		user.Username,
		user.City,
		user.CreatedAt,
	).Scan(&user.Id)

	if err != nil {
		return nil, err
	}

	user1, err := r.Get(&pb.ById{Userid: user.Id})
	if err != nil {
		return nil, err
	}

	return user1, nil
}

func (r *userRepo) Get(in *pb.ById) (*pb.User, error) {
	var (
		user       pb.User
		nullupdate sql.NullString
		nulldelete sql.NullString
	)
	fmt.Println(in.Userid)
	query := `SELECT
				id,
				name,
				username,
				city,
				created_at,
				updated_at,
				deleted_at
			FROM testuser
			WHERE id = $1 AND deleted_at IS NULL`

	err := r.db.QueryRow(query, in.Userid).Scan(
		&user.Id,
		&user.Name,
		&user.Username,
		&user.City,
		&user.CreatedAt,
		&nullupdate,
		&nulldelete,
	)

	fmt.Println(user.Id,"\n")

	if nullupdate.Valid {
		user.UpdatedAt = nullupdate.String
	}
	if nulldelete.Valid {
		user.DeletedAt = nulldelete.String
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) Update(user *pb.User) (*pb.User, error) {

	query := `UPDATE testuser SET
				name = $1,
				username = $2,
				city = $3,
				updated_at = $4
				WHERE id = $5 AND deleted_at IS NULL`
				
	fmt.Println(user.Id)
	err := r.db.QueryRow(query,
		user.Name,
		user.Username,
		user.City,
		time.Now().Format("2006-01-02"),
		user.Id,
	).Err()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return r.Get(&pb.ById{Userid: user.Id})
}

func (r *userRepo) Delete(in *pb.ById) (*pb.Empty, error) {
	query := `UPDATE testuser SET deleted_at = $1 WHERE id = $2`
	res, err := r.db.Exec(query, time.Now().Format("2006-01-02"), in.Userid)

	if err != nil {
		return nil, err
	}

	if a, _ := res.RowsAffected(); a > 0 || err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
