### 简介
项目层次结构  
router：路由层  
controller：负责页面的渲染，转发相应的请求数据到Service  
service：业务逻辑层，各种校验逻辑在此层进行，确认合法后传入至manager  
manager：数据的增删查改，专用于调用DAO，并提供接口供service层调用  
dao：调用数据库的增删查改方法，并提供DAO接口供manager层调用  
model：模拟数据库  