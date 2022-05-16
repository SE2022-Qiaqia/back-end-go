# 选课系统后端

## API列表

下面API没有说明权限，只是给了功能性列表。

- 系统初始化
  - [x] 检查是否初始化
  - [x] 初始化
- 公开API
  - [x] 登录
  - [x] 注册
- 用户相关
  - [x] 获取当前用户信息
  - [x] 获取其他用户信息
  - [x] 添加用户（与注册不一样）
  - [x] 获取用户列表
  - [x] 删除用户
  - [x] 更新用户信息
  - [x] 更新个人信息（非管理员仅限修改用户名）
  - [x] 修改密码
  - [x] 修改他人密码（管理员）
- 课程相关
  - [x] 获取可选课程列表
  - [x] 添加课程共信息
  - [x] 开设课头
  - [ ] 更新课程信息
  - [ ] 更新课头信息
  - [ ] 获取用户课程安排表
  - [ ] 查询课头内学生及相关信息
  - [ ] 登记课头内学生成绩
  - [ ] 选课
  - [ ] 撤课
  - [ ] 强制选课
  - [ ] 强制撤课
  - [ ] 登记成绩
- 学院相关
  - [x] 获取学院列表
  - [x] 开设学院
- 学期相关
  - [ ] 获取学期列表
  - [ ] 开设学期
