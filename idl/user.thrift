namespace go user

include "model.thrift"

struct RegisterRequest{
    1: required string username,
    2: required string password,
}
struct RegisterResponse{
    1: model.BaseResp base,
}

struct LoginRequest{
    1: required string username,
    2: required string password,
}

struct LoginResponse{
    1: model.BaseResp base,
    2: model.User data,
}

struct UploadAvatarRequest{
     1: binary data (api.form="data"),
    //body参数avatar文件
}
struct UploadAvatarResponse{
    1: model.BaseResp base,
    2: model.User data,
}
struct GetUserInformationRequest{
    1: required i64  user_id
}
struct GetUserInformationResponse{
    1: model.BaseResp base,
    2:model.User data,
}
service UserService {
    RegisterResponse Register (1: RegisterRequest req) (api.post="/user/register"),
    LoginResponse Login(1: LoginRequest req) (api.post="/user/login"),
    UploadAvatarResponse UploadAvatar(1:UploadAvatarRequest req)(api.put="/user/avatar/upload"),
    GetUserInformationResponse GetInformation(1:GetUserInformationRequest req)(api.get="/user/info")
}
