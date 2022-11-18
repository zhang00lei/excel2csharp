### unity中Excel转C#相关配置工具

相关步骤：将Excel相关数据转为json字符串，并保存在对应位置。
应用启动后在指定位置加载json，并使用litjson插件序列化为对应的类对象。

1. 需要注意的是，除了CommonConfig配置表，其他的配置表**首个字段名称需为Id，类型为int**，
   类型名称为#号的列不再解析。除了CommonConfig配置表，其他配置表中最多只能有一个enum字段，且对应注释为第一个#号列的内容。
2. 对于某些不能单独提取出作为一张表的字段配置，可配置在CommonConfig中
3. litjson略做修改，参考这里https://www.cnblogs.com/msxh/p/12541159.html
4. 支持相关数据类型如下

| 类型                  | 说明         | 默认值     |
|---------------------|------------|---------|
| int                 | 整形         | 0       |
| List&lt;int&gt;     | 整形集合       | null    |
| float               | 浮点型        | 0       |
| List&lt;float&gt;   | 浮点型集合      | null    |
| bool                | bool型      | false   |
| List&lt;bool&gt;    | bool型集合    | null    |
| string              | 字符串        | ""      |
| List&lt;string&gt;  | 字符串集合      | null    |
| vector2             | 二元组        | (0,0)   |
| List&lt;vector2&gt; | 二元组集合      | null    |
| vector3             | 三元组        | (0,0,0) |
| List&lt;vector3&gt; | 三元组集合      | (0,0,0) |
| Dict&lt;x,y&gt;     | 字典(x为基准类型) | null    |
| enum                | 枚举         |         |  