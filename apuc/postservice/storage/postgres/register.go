package postgres

import (
	
	pb "github/Services/apuc/postservice/genproto/post_service"

	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type PostRepo struct {
	db *sqlx.DB
}

func NewPostRepo(db *sqlx.DB) *PostRepo {
	return &PostRepo{db: db}
}

func (r *PostRepo) Create(req *pb.Post) (*pb.Post, error) {
	query := `
        INSERT INTO 
			postos (
				id,
				userid,
				title, 
				created_at,
				images )
        VALUES($1,$2,$3,$4,$5) 
		RETURNING id
    `
	req.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	a := []string{}
	for _, i := range req.Image {
		a = append(a, i.Image)
	}
	id := uuid.New().String()

	err := r.db.DB.QueryRow(query,
		id,
		req.UserId,
		req.Title,
		req.CreatedAt,
		pq.Array(a),
	).Scan(&req.Id)

	post, err := r.Get(&pb.ById{Userid: req.Id})
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (r *PostRepo) Delete(id *pb.ById) (*pb.Empty, error) {
	query := `
		UPDATE postos SET 
			deleted_at = $1
		WHERE id = $2
	`
	newTime := time.Now().Format("2006-01-02 15:04:05")

	_, err := r.db.DB.Exec(query, newTime, id.Userid)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (r *PostRepo) DeleteByUser(id *pb.ById) (*pb.Empty, error) {
	query := `
		UPDATE postos SET 
			deleted_at = $1
		WHERE userid = $2
	`
	newTime := time.Now().Format("2006-01-02 15:04:05")

	_, err := r.db.DB.Exec(query, newTime, id.Userid)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (r *PostRepo) Get(in *pb.ById) (*pb.Post, error) {
	var (
		media pb.Post
	)
	query := `SELECT 
			id, 
			userid,
			title, 
			images
	FROM postos 
	WHERE id = $1 AND deleted_at IS NULL`

	a := []string{}

	err := r.db.QueryRow(query, in.Userid).Scan(
		&media.Id,
		&media.UserId,
		&media.Title,
		pq.Array(&a),
	)
	if err != nil {
		return nil, err
	}

	for _, i := range a {
		media.Image = append(media.Image, &pb.Photo{Image: i})
	}

	return &media, nil
}

func (r *PostRepo) Update(n *pb.Post) (*pb.Post, error) {

	query := `UPDATE postos SET 
				title = $1,
				images = $2,
				update_at = $3
				WHERE id = $4 AND deleted_at IS NULL`

	n.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	a := []string{}
	for _, i := range n.Image {
		a = append(a, i.Image)
	}

	err := r.db.QueryRow(query,
		n.Title,
		pq.Array(a),
		n.UpdatedAt,
		n.Id).Err()
	if err != nil {
		return nil, err
	}

	post, err := r.Get(&pb.ById{Userid: n.Id})
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (r *PostRepo) GetById(req *pb.ByUId) (*pb.ListResp, error) {
	offset := (req.Page - 1) * req.Limit
	var resp pb.ListResp

	query := `SELECT 
				id, 
				userid,
				title, 
				images,
				created_at
			FROM postos 
			WHERE userid = $3 AND deleted_at IS NULL
			LIMIT $1
		    OFFSET $2`

	rows, err := r.db.DB.Query(query, req.Limit, offset, req.Userid)

	if err != nil {
		return nil, err
	}

	for rows.Next() {

		var Post pb.Post
		a := []string{}

		err = rows.Scan(
			&Post.Id,
			&Post.UserId,
			&Post.Title,
			pq.Array(&a),
			&Post.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		for _, i := range a {
			Post.Image = append(Post.Image, &pb.Photo{Image: i})
		}

		resp.Posts = append(resp.Posts, &Post)
		resp.Count += 1
	}

	return &resp, nil
}

func (r *PostRepo) List(req *pb.ListReq) (*pb.ListResp, error) {

	offset := (req.Page - 1) * req.Limit
	var resp pb.ListResp
	query := `
		SELECT 
			id, 
			userid,
			title, 
			images,
			created_at
        FROM postos
		WHERE deleted_at IS NULL
		LIMIT $1
		OFFSET $2
	`
	rows, err := r.db.DB.Query(query, req.Limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		var Post pb.Post
		a := []string{}

		err = rows.Scan(
			&Post.Id,
			&Post.UserId,
			&Post.Title,
			pq.Array(&a),
			&Post.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		for _, i := range a {
			Post.Image = append(Post.Image, &pb.Photo{Image: i})
		}

		resp.Posts = append(resp.Posts, &Post)
	}

	query = `
		SELECT count(*) FROM postos
		WHERE deleted_at IS NULL
	`
	err = r.db.DB.QueryRow(query).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
