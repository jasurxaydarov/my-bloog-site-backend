package postgres

import (
	"context"
	"jasurxaydarov/my-bloog-site-backend/modles"

	"github.com/google/uuid"
	"github.com/saidamir98/udevs_pkg/logger"
)

func (c *ContentRepo)CreateViwer(ctx context.Context,req modles.ViewerReqReg)(*modles.Viewer,error){

	id:=uuid.New()

	query:=`
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
	_,err:=c.db.Exec(
		ctx,
		query,
		id,
		req.FullName,
		req.Username,
		req.Gmail,
		req.Password,
	)
	
	if err!= nil{
		c.log.Error("error on Createing Viwer",logger.Error(err))
		return nil,nil
	}
	
	return nil,nil
}
func (c *ContentRepo)GetViwer(ctx context.Context,id string)(*modles.Viewer,error){
	
	
	return nil,nil
}
func (c *ContentRepo)GetViwers(ctx context.Context,list modles.GetList)(*modles.GetViwersResp,error){
	
	
	return nil,nil
}
func (c *ContentRepo)UpdateViwer(ctx context.Context,req modles.Viewer)(error){
	
	
	return nil
}
func (c *ContentRepo)DeleteViwer(ctx context.Context,id string)(error){
	
	
	return nil
}
