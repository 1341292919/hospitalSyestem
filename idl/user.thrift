namespace go user

include "model.thrift"

struct LoginRequest{
    1: required string username
    2: required string password,
    3: required i64 type,
    4: required i64 id,
}

struct LoginResponse{
    1: model.BaseResp base,
    2: optional model.User data,
}

struct NewUserRequest{
      1: required string username
      2: required string password,
      3: required i64 type,
      4: required i64 id,
}

struct NewUserResponse{
    1: model.BaseResp base,
    2: optional model.User data,
}

struct UpdateDoctorMessageRequest{
        1: optional string name
        2: required string specialty,
        3: required string title,
        4: required string department,
        5: required string contact_phone,
       6: required i64 id,
}
struct UpdateDoctorMessageResponse{
     1: model.BaseResp base,
}


struct UpdateNurseMessageRequest{
       1: optional string name,
       2: required string contact_phone,
       3: required string position,
       4: required string department,
       5:required i64 id,
}
struct UpdateNurseMessageResponse{
     1: model.BaseResp base,
}

struct QueryUserRequest{
       1: required i64 id,
       2: required i64 type,
}
struct QueryUserResponse{
     1: model.BaseResp base,
     2: model.User data,
}

service UserService {
    LoginResponse Login(1:LoginRequest req)(api.post="/user/login"),
    NewUserResponse NewUser(1:NewUserRequest req)(api.post="/user/add"),
    UpdateDoctorMessageResponse  UpdateDoctor(1:UpdateDoctorMessageRequest req)(api.put="/user/doctor/update"),
    UpdateNurseMessageResponse  UpdateNurse(1:UpdateNurseMessageRequest req)(api.put="/user/nurse/update"),
    QueryUserResponse Query(1:QueryUserRequest req)(api.get="/user/info"),
}
