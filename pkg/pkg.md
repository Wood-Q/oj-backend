# ./pkg

**Folder with project specific functionality**. 存放项目的通用功能模块，适用于业务无关的代码逻辑。

- `./pkg/configs` 项目的配置文件或配置相关逻辑，方便集中管理和修改。
- `./pkg/middleware` 存放 Fiber 框架的中间件，比如认证、日志、错误处理等。
- `./pkg/routes` 定义了所有 API 的路由配置，将请求与具体控制器绑定。
- `./pkg/enums` 集中定义和管理常量
- `./pkg/utils` 存储工具函数或通用方法，例如日期处理、加密解密等逻辑。
