package postgres

import (
	pb "github/Services/apuc/userservice/genproto/user_service"
	"database/sql"
	"fmt"

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
	query := `
    INSERT INTO user2 (
			id, 
			name, 
			username, 
			city, 
			created_at )
    VALUES($1, $2, $3, $4, $5) RETURNING id`

	id := uuid.New().String()

	err := r.db.QueryRow(query,
		id,
		user.Name,
		user.Username,
		user.City,
		time.Now().Format("2006-01-02"),
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
	query := `SELECT 
				id, 
				name, 
				username,  
				city, 
				created_at, 
				updated_at, 
				deleted_at 
			FROM user2 
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

func (r *userRepo) Delete(in *pb.ById) (*pb.Empty, error) {
	query := `UPDATE user2 SET deleted_at = $1 WHERE id = $2`

	res, err := r.db.Exec(query, time.Now().Format("2006-01-02"), in.Userid)

	if err != nil {
		return nil, err
	}

	if a, _ := res.RowsAffected(); a > 0 || err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (r *userRepo) Update(user *pb.User) (*pb.User, error) {

	query := `UPDATE user2 SET 
				name = $1, 
				username = $2, 
				city = $3, 
				updated_at = $4 
				WHERE id = $5 AND deleted_at IS NULL`
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

func (r *userRepo) ListUser(req *pb.ListReq) (*pb.ListResp, error) {

	offset := (req.Page - 1) * req.Limit
	var (
		nullupdate sql.NullString
		nulldelete sql.NullString
		resp pb.ListResp
	)
	query := `
		SELECT 
			id, 
			name, 
			username, 
			city, 
			created_at,
			deleted_at,
			updated_at
        FROM user2
		WHERE deleted_at IS NULL
		LIMIT $1
		OFFSET $2
	`
	rows, err := r.db.DB.Query(query, req.Limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var User pb.User
		err = rows.Scan(
			&User.Id,
			&User.Name,
			&User.Username,
			&User.City,
			&User.CreatedAt,
			&nulldelete,
			&nullupdate,
		)
		if err != nil {
			return nil, err
		}
		if nullupdate.Valid {
			User.UpdatedAt = nullupdate.String
		}
		if nulldelete.Valid {
			User.DeletedAt = nulldelete.String
		}
	

		resp.Users = append(resp.Users, &User)
	}

	query = `
		SELECT count(*) FROM user2
		WHERE deleted_at IS NULL
	`
	err = r.db.DB.QueryRow(query).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}
	
	return &resp, nil
}


