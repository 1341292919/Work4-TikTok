// Code generated by hertz generator.

package interact

import (
	"TikTok/biz/pack"
	"TikTok/biz/service"
	"context"

	interact "TikTok/biz/model/interact"
	"github.com/cloudwego/hertz/pkg/app"
)

// HitLikeButton .
// @router /interact/like [PUT]
func HitLikeButton(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interact.LikeVideoRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp := new(interact.LikeVideoResponse)
	err = service.NewInteractService(ctx, c).LikeVideo(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}
	resp.Base = pack.BuildBaseResp(nil)
	pack.SendResponse(c, resp)

}

// QueryLikeList .
// @router /interact/query/like [GET]
func QueryLikeList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interact.QueryLikeListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp := new(interact.QueryLikeListResponse)
	videoList, count, err := service.NewInteractService(ctx, c).QueryLikeList(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}
	resp.Base = pack.BuildBaseResp(nil)
	resp.Data = pack.VideoList(videoList, count)
	pack.SendResponse(c, resp)
}

// QueryCommentList .
// @router /interact/query/comment [GET]
func QueryCommentList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interact.QueryCommentListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp := new(interact.QueryCommentListResponse)

	commentList, count, err := service.NewInteractService(ctx, c).QueryCommentList(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}
	resp.Base = pack.BuildBaseResp(nil)
	resp.Data = pack.CommentList(commentList, count)
	pack.SendResponse(c, resp)
}

// DeleteComment .
// @router /interact/delete [DELETE]
func DeleteComment(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interact.DeleteCommentRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp := new(interact.DeleteCommentResponse)

	err = service.NewInteractService(ctx, c).DeleteComment(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}
	resp.Base = pack.BuildBaseResp(nil)
	pack.SendResponse(c, resp)
}

// CommentVideo .
// @router /comment/list [GET]
func CommentVideo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interact.CommentRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}
	resp := new(interact.CommentResponse)
	err = service.NewInteractService(ctx, c).CommentVideo(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}
	resp.Base = pack.BuildBaseResp(nil)
	pack.SendResponse(c, resp)
}
