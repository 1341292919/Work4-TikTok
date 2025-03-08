namespace go socialize

include "model.thrift"

struct FollowRequest{
    1:required i64 to_user_id,
    2:required i64 action_type,//0关注 1取关
}
struct FollowResponse{
    1:model.BaseResp base,
}

struct QueryFollowListRequest{  //查看对应id的关注
     1: required i64 page_size,  //每一页的数量
     2: required i64 page_num,   //页码
     3: required i64 user_id,
}
struct QueryFollowListResponse{
      1:model.BaseResp base,
      2:model.SimpleUserList data,
}
struct QueryFollowerListRequest{ //查看指定id的粉丝
     1: required i64 page_size,  //每一页的数量
     2: required i64 page_num,   //页码
     3: required i64 user_id,
}
struct QueryFollowerListResponse{
      1:model.BaseResp base,
      2:model.SimpleUserList data,
}
struct QueryFriendListRequest{
      1: required i64 page_size,  //每一页的数量
      2: required i64 page_num,   //页码
}
struct QueryFriendListResponse{
      1:model.BaseResp base,
      2:model.SimpleUserList data,
}

service SocializeService{
    FollowResponse Follow(1:FollowRequest req)(api.post="/relation/action"),
    QueryFollowListResponse QueryFollowList(1:QueryFollowListRequest req)(api.get="/following/list"),
    QueryFollowerListResponse QueryFollowerList(1:QueryFollowerListRequest req)(api.get="/follower/list"),
    QueryFriendListResponse QueryFriendList(1:QueryFriendListRequest req)(api.get="/friends/list"),
}
