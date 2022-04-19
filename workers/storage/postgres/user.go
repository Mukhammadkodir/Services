package postgres

import (
	"database/sql"
	pb "github/Services/workers/genproto/user_service"

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
			f_name, 
			l_name, 
			password,
			monthly,
			position, 
			created_at )
			VALUES($1, $2, $3, $4, $5, $6, $7)`

	id := uuid.New().String()

	err := r.db.QueryRow(query,
		id,
		user.FName,
		user.LName,
		user.Password,
		user.Monthly,
		user.Position,
		time.Now().Format("2006-01-02"),
	).Err()

	if err != nil {
		return nil, err
	}

	user1, err := r.Login(&pb.PasswordReq{Password: user.Password})
	if err != nil {
		return nil, err
	}

	return user1, nil
}

func (r *userRepo) Get(in *pb.PasswordReq) (*pb.GetUser, error) {
	var (
		user       pb.GetUser
		nullupdate sql.NullString
		nulldelete sql.NullString
	)
	query := `SELECT 
				id, 
				f_name, 
				l_name, 
				password,
				position,
				created_at, 
				updated_at, 
				deleted_at 
			FROM user2 
			WHERE password = $1 AND deleted_at IS NULL`
	err := r.db.QueryRow(query, in.Password).Scan(
		&user.Id,
		&user.FName,
		&user.LName,
		&user.Password,
		&user.Position,
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

func (r *userRepo) Login(in *pb.PasswordReq) (*pb.User, error) {
	var (
		user       pb.User
		nullupdate sql.NullString
		nulldelete sql.NullString
	)
	query := `SELECT 
				id, 
				f_name, 
				l_name, 
				password,
				position,
				monthly,
				created_at, 
				updated_at, 
				deleted_at 
			FROM user2 
			WHERE password = $1 AND deleted_at IS NULL`
	err := r.db.QueryRow(query, in.Password).Scan(
		&user.Id,
		&user.FName,
		&user.LName,
		&user.Password,
		&user.Position,
		&user.Monthly,
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

func (r *userRepo) Delete(in *pb.PasswordReq) (*pb.EmptyResp, error) {

	query := `UPDATE user2 SET deleted_at = $1 WHERE password = $2`

	res, err := r.db.Exec(query, time.Now().Format("2006-01-02"), in.Password)

	if err != nil {
		return nil, err
	}

	if a, _ := res.RowsAffected(); a > 0 || err != nil {
		return nil, err
	}

	return &pb.EmptyResp{}, nil
}

func (r *userRepo) Update(user *pb.UpReq) (*pb.User, error) {

	query := `UPDATE users SET 
				f_name 		= $1, 
				l_name 		= $2, 
				password	= $3,
				position 	= $4
				updated_at 	= $5 
				WHERE password 	= $6 AND deleted_at IS NULL`
	err := r.db.QueryRow(query,
		user.FName,
		user.LName,
		user.NewPassword,
		user.Position,
		time.Now().Format("2006-01-02"),
		user.OldPassword,
	).Err()

	if err != nil {
		return nil, err
	}

	return r.Login(&pb.PasswordReq{Password: user.NewPassword})
}

func (r *userRepo) ListUser(req *pb.ListReq) (*pb.ListResp, error) {

	offset := (req.Page - 1) * req.Limit
	var (
		nullupdate sql.NullString
		nulldelete sql.NullString
		resp       pb.ListResp
	)
	query := `
		SELECT 
			id, 
			f_name, 
			l_name, 
			password,
			monthly,
			position,
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
			&User.FName,
			&User.LName,
			&User.Password,
			&User.Monthly,
			&User.Position,
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

func (r *userRepo) CheckField(req *pb.PasswordReq) (*pb.Status, error) {
	var existsClient pb.Status

	row := r.db.QueryRow(`
		SELECT position FROM user1 WHERE  password = $1 AND deleted_at IS NULL`, req.Password,
	)
	if err := row.Scan(&existsClient); err != nil {
		return &pb.Status{Position: "false"}, err
	}

	return &existsClient, nil
}

func (r *userRepo) OpenDay(req *pb.ById) (*pb.Hours, error) {
	query := `
    INSERT INTO time (
			id, 
			user_id, 
			opened,
			date )
	VALUES($1, $2, $3, $4 )`

	id := uuid.New().String()

	err := r.db.QueryRow(query,
		id,
		req.Userid,
		time.Now().Format("2006-01-02 15:04:00"),
		time.Now().Format("2006-01-02"),
	).Err()

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *userRepo) CloseDay(req *pb.Hours) (*pb.Hours, error) {
	query := `
		UPDATE time set (
			daily 		= $1, 
			closed 		= $2, 
			monthly 	= $3
		WHERE id 	= $4 `

	query1 := `
			SELECT date_part('hour',age('$1', '$2'));
			`
	query2 := `SELECT *
	FROM time 
	WHERE date > CURRENT_DATE - INTERVAL '1 MONTH' 
	  AND date < CURRENT_DATE + INTERVAL '1 DAY' AND user_id = $1 ;`

	req.Closed = time.Now().Format("2006-01-02 15:04:00")

	err := r.db.QueryRow(query1,
		req.Opened,
		req.Closed,
	).Scan(&req.Daily)

	if err != nil {
		return nil, err
	}

	var day int
	rows, err := r.db.Query(query2, req.UserId)
	for rows.Next() {
		err = rows.Scan(
			&day,
		)

		day += day
	}
	req.Monthly = string(day)

	err = r.db.QueryRow(query,
		req.Daily,
		req.Closed,
		req.Monthly,
		req.Id,
	).Err()

	if err != nil {
		return nil, err
	}

	return nil, nil
}

// func (r *userRepo) ListTime(req *pb.ById) (*pb.Status, error) {

// 	var (
// 		resp int64
// 	)

// 	query := `
// 		SELECT 
// 			daily
// 		FROM time
// 		WHERE user_id = $1 AND closed NOT NULL
// 	`
// 	rows, err := r.db.DB.Query(query, req.Userid)

// 	if err != nil {
// 		return nil, err
// 	}

// 	for rows.Next() {
// 		var time pb.Status
// 		err = rows.Scan(
// 			&time.Position,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}

// 		resp += int64(time.Position)
// 	}

// 	query = `
// 		SELECT count(*) FROM user2
// 		WHERE deleted_at IS NULL
// 	`
// 	err = r.db.DB.QueryRow(query).Scan(&resp.Count)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &resp, nil
// }

//select date_part('hour',age('2022-04-19 10:00:00', '2022-04-30 12:00:00'));
