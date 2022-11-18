### unity中Excel转C#相关配置工具
相关步骤：将Excel相关数据转为json字符串，并保存在对应位置。
应用启动后在指定位置加载json，并使用litjson插件序列化为对应的类对象。
支持相关数据类型如下：

| 类型          | 说明    | 默认值     |
|-------------|-------|---------|
| int         | 整形    | 0       |
| List<int>   | 整形集合  | null    |
| float       | 浮点型   | 0       |
| List<float> | 浮点型集合 | null    |
| bool        | bool型 | false   |
| vector2     | 二元组   | (0,0)   |
| vector3     | 三元组   | (0,0,0) |
