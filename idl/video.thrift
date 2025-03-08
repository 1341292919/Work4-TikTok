namespace go video

include"model.thrift"

struct PublishRequest{
    1: required string title,
    2: required string description,
    3: required binary data (api.form="data"),//视频file
}

struct PublishResponse{
     1: model.BaseResp base,
}
//发布列表
struct QueryPublishListRequest{
     1: required i64 user_id,
     2: required i64 page_size,  //每一页的数量
     3: required i64 page_num,   //页码
}

struct QueryPublishListResponse{
    1:model.BaseResp base,
    2:model.VideoList data,
}
struct SearchVideoByKeywordRequest{
     1: required i64 page_size,  //每一页的数量
     2: required i64 page_num,   //页码
     3: required string keyword, //关键词
     4: i64 from_date,
     5: i64 to_date,
     6:string username
}

struct SearchVideoByKeywordResponse{
     1:model.BaseResp base,
     2:model.VideoList data,
}

struct GetPopularListRequest{
      1: required i64 page_size,  //每一页的数量
      2: required i64 page_num,   //页码
}
struct GetPopularListResponse{
     1:model.BaseResp base,
     2:model.VideoList data,
}

service VideoService{
    PublishResponse PublishVideo(1:PublishRequest req) (api.post="/video/publish"),
    QueryPublishListResponse QueryList(1:QueryPublishListRequest req)(api.get="/video/list"),
    SearchVideoByKeywordResponse SearchVideo(1:SearchVideoByKeywordRequest req)(api.post="/video/search"),
    GetPopularListResponse GetPopularVideo(1:GetPopularListRequest req)(api.get="/video/popular"),
}