package postgres

import (
	"context"
	"errors"
	"jasurxaydarov/my-bloog-site-backend/modles"
	"jasurxaydarov/my-bloog-site-backend/pgx/helpers"

	"github.com/jackc/pgx/v5"
	"github.com/saidamir98/udevs_pkg/logger"
)

type viwerRepo struct {
	db  *pgx.Conn
	log logger.LoggerI
}

func NewViwerRepo(db *pgx.Conn, log logger.LoggerI) ViwerRepoI {

	return &viwerRepo{db, log}
}
func (v *viwerRepo) CreateViwer(ctx context.Context, req modles.Viewer) (*modles.Clamis, error) {

	query := `
		INSERT INTO viwers(
			viwer_id,
			fullname,
			username,
			gmail,
			password 
		)VALUES(
			$1,$2,$3,$4,$5
		)
	`
	_, err := v.db.Exec(
		ctx,
		query,
		req.ViewerID.String(),
		req.FullName,
		req.Username,
		req.Gmail,
		req.Password,
	)

	if err != nil {
		v.log.Error("error on Createing Viwer", logger.Error(err))
		return nil, err
	}

	resp,err:=v.GetViwerClaims(ctx,req.ViewerID.String())
	
	return resp, nil
}

func (v *viwerRepo) LogInViwer(ctx context.Context, req modles.LoginViwer) (*modles.Clamis, error) {

	var viwerId,gmail,hashPassword string

	query:=`
		SELECT 
			viwer_id,
			gmail,
			password
		FROM
			viwers
		WHERE	
			username =$1
	`

	err:=v.db.QueryRow(ctx,query,req.Username).Scan(&viwerId,&gmail,&hashPassword)

	if err != nil{
		return nil,err
	}

	if !helpers.CompareHashPassword(hashPassword,req.Password){
		return nil,errors.New("password is incorrect")
	}


	return &modles.Clamis{UserId: viwerId,UserRole:"viwer"}, nil
}

func (v *viwerRepo) GetViwer(ctx context.Context, id string) (*modles.Viewer, error) {
	var resp modles.Viewer
	query:=`
		SELECT 
			* 
		FROM 
			viwers
		WHERE
			viwer_id = $1
	`
	//  viwer_id               |    fullname    | username |            gmail            | password 
	err:=v.db.QueryRow(ctx,query,id).Scan(
			&resp.ViewerID,
			&resp.FullName,
			&resp.Username,
			&resp.Gmail,
			&resp.Password,
		)

		if err!=nil{
			v.log.Error("err on GetViwer",logger.Error(err))
			return nil,nil
		}
	return &resp, nil
}
func (c *viwerRepo) GetViwers(ctx context.Context, list modles.GetList) (*modles.GetViwersResp, error) {

	return nil, nil
}
func (c *viwerRepo) UpdateViwer(ctx context.Context, req modles.Viewer) error {

	return nil
}
func (c *viwerRepo) DeleteViwer(ctx context.Context, id string) error {

	return nil
}
func(v *viwerRepo)GetViwerClaims(ctx context.Context, id string) (*modles.Clamis, error){

	var resp string
	query:=`
		SELECT 
			viwer_id
		FROM 
			viwers
		WHERE
			viwer_id = $1
	`
	//  viwer_id               |    fullname    | username |            gmail            | password 
	err:=v.db.QueryRow(ctx,query,id).Scan(&resp)

		if err!=nil{
			v.log.Error("err on GetViwer",logger.Error(err))
			return nil,nil
		}
	return &modles.Clamis{UserId: resp,UserRole: "viwer"}, nil
}


