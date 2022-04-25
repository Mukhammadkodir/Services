package postgres

import (
	"database/sql"

	"github/Services/workers/storage/models"

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
func (r *userRepo) Create(m models.User) (*models.User, error) {
	query := `
    INSERT INTO users ( id, 
			f_name, 
			l_name, 
			password,
			position, 
			created_at )
			VALUES($1, $2, $3, $4, $5, $6)`

	id := uuid.New().String()

	err := r.db.QueryRow(query,
		id,
		m.F_name,
		m.L_name,
		m.Password,
		m.Position,
		time.Now().Format("2006-01-02"),
	).Err()

	if err != nil {
		return nil, err
	}

	user1, err := r.Login(models.PasswordReq{Password: m.Password})
	if err != nil {
		return nil, err
	}

	return user1, nil
}

func (r *userRepo) Get(id models.ById) (*models.Get, error) {
	var (
		user       models.Get
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
			FROM users 
			WHERE id = $1 AND deleted_at IS NULL`

	err := r.db.QueryRow(query, id.Userid).Scan(
		&user.Id,
		&user.F_name,
		&user.L_name,
		&user.Password,
		&user.Position,
		&user.Created_at,
		&nullupdate,
		&nulldelete,
	)

	if nullupdate.Valid {
		user.Updated_at = nullupdate.String
	}
	if nulldelete.Valid {
		user.Deleted_at = nulldelete.String
	}

	if err != nil {
		return nil, err
	}

	return r.ListUserDay(user)
}

func (r *userRepo) Login(pass models.PasswordReq) (*models.User, error) {
	var (
		user       models.User
		nullupdate sql.NullString
	)
	query := `SELECT 
				id, 
				f_name, 
				l_name, 
				password,
				position,
				monthly,
				created_at, 
				updated_at
			FROM users 
			WHERE password = $1 AND deleted_at IS NULL`

	err := r.db.QueryRow(query, pass.Password).Scan(
		&user.Id,
		&user.F_name,
		&user.L_name,
		&user.Password,
		&user.Position,
		&user.Monthly,
		&user.Created_at,
		&nullupdate,
	)

	if nullupdate.Valid {
		user.Updated_at = nullupdate.String
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) Delete(in models.PasswordReq) (*models.EmptyResp, error) {

	query := `UPDATE users SET deleted_at = $1 WHERE password = $2`

	res, err := r.db.Exec(query, time.Now().Format("2006-01-02"), in.Password)

	if err != nil {
		return nil, err
	}

	if a, _ := res.RowsAffected(); a > 0 || err != nil {
		return nil, err
	}

	return &models.EmptyResp{Message: "Deleted Ok"}, nil
}

func (r *userRepo) Update(user models.UpReq) (*models.User, error) {

	query := `UPDATE users SET 
				f_name 		= $1, 
				l_name 		= $2, 
				password	= $3,
				position 	= $4
				updated_at 	= $5 
				WHERE password 	= $6 AND deleted_at IS NULL`

	err := r.db.QueryRow(query,
		user.F_name,
		user.L_name,
		user.New_password,
		user.Position,
		time.Now().Format("2006-01-02"),
		user.Old_password,
	).Err()

	if err != nil {
		return nil, err
	}

	return r.Login(models.PasswordReq{Password: user.New_password})
}

func (r *userRepo) ListUser(req models.ListReq) (*models.ListResp, error) {

	offset := (req.Page - 1) * req.Limit
	var (
		nullupdate sql.NullString
		nulldelete sql.NullString
		resp       models.ListResp
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
        FROM users
		WHERE deleted_at IS NULL
		LIMIT $1
		OFFSET $2
	`
	rows, err := r.db.DB.Query(query, req.Limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var User models.User
		err = rows.Scan(
			&User.Id,
			&User.F_name,
			&User.L_name,
			&User.Password,
			&User.Monthly,
			&User.Position,
			&User.Created_at,
			&nulldelete,
			&nullupdate,
		)
		if err != nil {
			return nil, err
		}
		if nullupdate.Valid {
			User.Updated_at = nullupdate.String
		}
		if nulldelete.Valid {
			User.Deleted_at = nulldelete.String
		}

		resp.Users = append(resp.Users, User)
	}

	query = `
		SELECT count(*) FROM users
		WHERE deleted_at IS NULL
	`
	err = r.db.DB.QueryRow(query).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (r *userRepo) CheckField(pass models.PasswordReq) (*models.Status, error) {
	var existsClient models.Status

	row := r.db.QueryRow(`
		SELECT position FROM users WHERE  password = $1 AND deleted_at IS NULL`, pass.Password,
	)
	if err := row.Scan(&existsClient); err != nil {
		return &models.Status{Position: "false"}, err
	}

	return &existsClient, nil
}

func (r *userRepo) OpenDay(req models.ById) (*models.Hour, error) {
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

	return r.GetTime(models.ById{Userid: id})
}

func (r *userRepo) CloseDay(req models.Hour) (*models.Hour, error) {
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

	req.Klozed = time.Now().Format("2006-01-02 15:04:00")

	err := r.db.QueryRow(query1,
		req.Opened,
		req.Klozed,
	).Scan(&req.Daily)

	if err != nil {
		return nil, err
	}

	var day int
	rows, err := r.db.Query(query2, req.User_id)
	for rows.Next() {
		err = rows.Scan(
			&day,
		)

		day += day
	}
	req.Monthly = string(day)

	err = r.db.QueryRow(query,
		req.Daily,
		req.Klozed,
		req.Monthly,
		req.ID,
	).Err()

	if err != nil {
		return nil, err
	}

	return r.GetTime(models.ById{Userid: req.ID})
}

func (r *userRepo) ListUserDay(req models.Get) (*models.Get, error) {

	query := `SELECT *
	FROM time 
	WHERE date > CURRENT_DATE - INTERVAL '1 MONTH' 
	  AND date < CURRENT_DATE + INTERVAL '1 DAY' AND user_id = $1 ;`

	rows, err := r.db.Query(query, req.Id)
	for rows.Next() {
		var day models.Hour
		err = rows.Scan(
			&day.ID,
			&day.User_id,
			&day.Opened,
			&day.Klozed,
			&day.Daily,
			&day.Monthly,
			&day.Date,
		)
		req.Hours = append(req.Hours, day)

	}

	if err != nil {
		return nil, err
	}

	return &req, nil
}

func (r *userRepo) GetTime(id models.ById) (*models.Hour, error) {
	var (
		Time models.Hour
	)
	query := `SELECT 
				id, 
				user_id, 
				opened,
				closed, 
				daily, 
				date,
				monthly,
			FROM time 
			WHERE id = $1`

	err := r.db.QueryRow(query, id.Userid).Scan(
		&Time.ID,
		&Time.User_id,
		&Time.Opened,
		&Time.Klozed,
		&Time.Daily,
		&Time.Date,
		&Time.Monthly,
	)

	if err != nil {
		return nil, err
	}

	return &Time, nil
}
