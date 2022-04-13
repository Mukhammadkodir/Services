package postgres

import (
	pb "github/Services/apuc/commentservice/genproto/comment_service"
	"database/sql"
	"fmt"

	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type commentRepo struct {
	db *sqlx.DB
}

//NewcommentRepo ...
func NewCommentRepo(db *sqlx.DB) *commentRepo {
	return &commentRepo{db: db}
}

func (r *commentRepo) Create(comment *pb.Comment) (*pb.Comment, error) {
	query := `
    INSERT INTO comment (
			id, 
			user_id, 
			post_id, 
			text, 
			created_at )
    VALUES($1, $2, $3, $4, $5) RETURNING id`

	id := uuid.New().String()

	err := r.db.QueryRow(query,
		id,
		comment.PostId,
		comment.Text,
		comment.UserId,
		time.Now().Format("2006-01-02"),
	).Scan(&comment.Id)

	if err != nil {
		return nil, err
	}

	comment1, err := r.Get(&pb.ById{Id: comment.Id})
	if err != nil {
		return nil, err
	}

	return comment1, nil
}

func (r *commentRepo) Get(in *pb.ById) (*pb.Comment, error) {
	var (
		comment       pb.Comment
		nullupdate sql.NullString
		nulldelete sql.NullString
	)
	query := `SELECT 
				id, 
				user_id, 
				post_id, 
				text, 
				created_at, 
				updated_at, 
				deleted_at 
			FROM comment 
			WHERE id = $1 AND deleted_at IS NULL`

	err := r.db.QueryRow(query, in.Id).Scan(
		&comment.Id,
		&comment.UserId,
		&comment.PostId,
		&comment.Text,
		&comment.CreatedAt,
		&nullupdate,
		&nulldelete,
	)

	if nullupdate.Valid {
		comment.UpdatedAt = nullupdate.String
	}
	if nulldelete.Valid {
		comment.DeletedAt = nulldelete.String
	}

	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func (r *commentRepo) Delete(in *pb.ById) (*pb.Empty, error) {
	query := `UPDATE comment SET deleted_at = $1 WHERE id = $2`

	res, err := r.db.Exec(query, time.Now().Format("2006-01-02"), in.Id)

	if err != nil {
		return nil, err
	}

	if a, _ := res.RowsAffected(); a > 0 || err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (r *commentRepo) Update(comment *pb.Comment) (*pb.Comment, error) {

	query := `UPDATE comment SET 
				text = $1,  
				updated_at = $2 
				WHERE id = $3 AND deleted_at IS NULL`
	err := r.db.QueryRow(query,
		comment.Text,
		time.Now().Format("2006-01-02"),
		comment.Id,
	).Err()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return r.Get(&pb.ById{Id: comment.Id})
}

func (r *commentRepo) ListComment(req *pb.ListReq) (*pb.ListResp, error) {

	offset := (req.Page - 1) * req.Limit
	var (
		nullupdate sql.NullString
		nulldelete sql.NullString
		resp pb.ListResp
	)
	query := `
		SELECT 
			id, 
			user_id, 
			post_id, 
			text, 
			created_at,
			deleted_at,
			updated_at
        FROM comment
		WHERE deleted_at IS NULL
		LIMIT $1
		OFFSET $2
	`
	rows, err := r.db.DB.Query(query, req.Limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var comment pb.Comment
		err = rows.Scan(
			&comment.Id,
			&comment.UserId,
			&comment.PostId,
			&comment.Text,
			&comment.CreatedAt,
			&nulldelete,
			&nullupdate,
		)
		if err != nil {
			return nil, err
		}
		if nullupdate.Valid {
			comment.UpdatedAt = nullupdate.String
		}
		if nulldelete.Valid {
			comment.DeletedAt = nulldelete.String
		}
	

		resp.Comments = append(resp.Comments, &comment)
	}

	query = `
		SELECT count(*) FROM comment
		WHERE deleted_at IS NULL
	`
	err = r.db.DB.QueryRow(query).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}
	
	return &resp, nil
}


