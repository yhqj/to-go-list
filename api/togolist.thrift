    namespace go togolist

    struct AddTaskReq {
        1: string Name (api.query="name");// 添加 api 注解为方便进行参数绑定
        2: string Type (api.query="type");// type就是
    }

    struct HelloResp {
        1: string RespBody;
    }


    service HelloService {
        HelloResp HelloMethod(1: AddTaskReq request) (api.get="/hello");
    }