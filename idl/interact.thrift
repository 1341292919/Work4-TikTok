namespace go interact

include "model.thrift"

struct LikeVideoRequest{
    1:required i64 video_id,
    2:required i64 action_type,
}
struct LikeVideoResponse{
    1:model.BaseResp base,
}
//点赞列表是用户的点赞列表
struct QueryLikeListRequest{
    1:required i64 user_id,
    2: required i64 page_size,  //每一页的数量
    3: required i64 page_num,   //页码
}

struct QueryLikeListResponse{
    1:model.BaseResp base,
    2:model.VideoList data,
}

struct CommentRequest{
    1:required i64 video_id,
    //可选，video_id 和 comment_id 必须存在其中一个
    //comment_id是什么鬼？？
    2:required string content,
}

struct CommentResponse{
    1:model.BaseResp base,
}

struct QueryCommentListRequest{
    1: required i64 video_id,
        //可选，video_id 和 comment_id 必须存在其中一个
    2: required i64 page_size,  //每一页的数量
    3: required i64 page_num,   //页码
}

struct QueryCommentListResponse{
    1:model.BaseResp base,
    2:model.CommentList data,
}

struct DeleteCommentRequest{
    1: i64 video_id,
    2: i64 comment_id
}

struct DeleteCommentResponse{
    1:model.BaseResp base,
}

service InteractService{
   LikeVideoResponse HitLikeButton(1:LikeVideoRequest req)(api.post="/like/action"),
   QueryLikeListResponse QueryLikeList(1:QueryLikeListRequest req)(api.get="/like/list"),
   CommentResponse CommentVideo(1:CommentRequest req)(api.post="/comment/publish"),
   QueryCommentListResponse QueryCommentList(1:QueryCommentListRequest req)(api.get="/comment/list"),
   DeleteCommentResponse DeleteComment(1:DeleteCommentRequest req)(api.delete="/comment/delete")
}