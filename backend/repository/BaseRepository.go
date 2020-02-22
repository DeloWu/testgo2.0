package repository

import (
    "gopkg.in/mgo.v2/bson"
    "testgo/common/datasource"
    "testgo/common/logger"
)

//依赖注入
type BaseRepository struct {
    Source datasource.IDb `inject:""`
    Log logger.ILogger `inject:""`
}

//获取第一条数据,默认按照createTime升序排序
func (b *BaseRepository) First(colName string, maps interface{}, result interface{}, orders ...string) (error){
    //Sort 默认升序, -field 为相反，即为降序
    err := b.Source.DB().Session.Ping()
    if err != nil{
        b.Log.Errorf("数据库连接失败: %v", err)
        return err
    }
    if len(orders) == 0 {
        err = b.Source.DB().C(colName).Find(maps).Sort("createTime").One(&result)
    }else {
        tempQuery := b.Source.DB().C(colName).Find(maps)
        for _, field := range orders{
            tempQuery = tempQuery.Sort(field)
        }
        err = tempQuery.One(&result)
    }

    if err != nil{
        b.Log.Error(err)
        return err
    }
    return nil
}

//获取最后一条数据,按照createTime升序排序
func (b *BaseRepository) Last(colName string, maps interface{}, result interface{}, orders ...string) (error){
    //Sort 默认升序, -field 为相反，即为降序
    err := b.Source.DB().Session.Ping()
    if err != nil{
        b.Log.Errorf("数据库连接失败: %v", err)
        return err
    }
    if len(orders) == 0 {
        err = b.Source.DB().C(colName).Find(maps).Sort("-createTime").One(&result)
    }else {
        tempQuery := b.Source.DB().C(colName).Find(maps)
        for _, field := range orders{
            tempQuery = tempQuery.Sort(field)
        }
        err = tempQuery.One(&result)
    }
    if err != nil{
        b.Log.Error(err)
        return err
    }
    return nil
}

//默认顺序获取一条数据
func (b *BaseRepository) FindOne(colName string, maps interface{}, result interface{}) (error){
    err := b.Source.DB().Session.Ping()
    if err != nil{
        b.Log.Errorf("数据库连接失败: %v", err)
        return err
    }
    err = b.Source.DB().C(colName).Find(maps).One(&result)
    if err != nil{
        b.Log.Error(err)
        return err
    }
    return nil
}

//根据id获取一条数据
func (b *BaseRepository) FindById(colName string, id string, result interface{}) (error){
    err := b.Source.DB().Session.Ping()
    if err != nil{
        b.Log.Errorf("数据库连接失败: %v", err)
        return err
    }
    err = b.Source.DB().C(colName).FindId(bson.ObjectIdHex(id)).One(result)
    if err != nil{
        b.Log.Error(err)
        return err
    }
    return nil
}

//获取所有数据
func (b *BaseRepository) FindAll(colName string, maps interface{}, result interface{}, orders ...string) (error){
    err := b.Source.DB().Session.Ping()
    if err != nil{
        b.Log.Errorf("数据库连接失败: %v", err)
        return err
    }
    if len(orders) == 0{
        err = b.Source.DB().C(colName).Find(maps).All(&result)
    }else {
        tempQuery := b.Source.DB().C(colName).Find(maps)
        for _, field := range orders{
            tempQuery = tempQuery.Sort(field)
        }
        err = tempQuery.All(&result)
    }
    if err != nil{
        b.Log.Error(err)
        return err
    }
    return nil
}

//获取总数
func (b *BaseRepository) Count(colName string, maps interface{}) (int, error){
    err := b.Source.DB().Session.Ping()
    if err != nil{
        b.Log.Errorf("数据库连接失败: %v", err)
        return 0, err
    }
    num, err := b.Source.DB().C(colName).Find(maps).Count()
    return num, err
}

//分页查询
func (b *BaseRepository) FindByPagination(colName string, maps interface{}, result interface{}, pageIndex int, pageSize int, orders ...string) (error){
    err := b.Source.DB().Session.Ping()
    if err != nil{
        b.Log.Errorf("数据库连接失败: %v", err)
        return err
    }
    tempQuery := b.Source.DB().C(colName).Find(maps)
    if len(orders) == 0 {
        tempQuery = tempQuery.Sort("-createTime")
    }else {
        for _, field := range orders{
            tempQuery = tempQuery.Sort(field)
        }
    }
    err = tempQuery.Skip((pageIndex -1) * pageSize).Limit(pageSize).All(result)
    if err != nil{
        b.Log.Error(err)
        return err
    }
    return nil
}

//插入数据
func (b *BaseRepository) Insert(colName string, docs ...interface{}) error{
    err := b.Source.DB().Session.Ping()
    if err != nil{
        b.Log.Errorf("数据库连接失败: %v", err)
        return err
    }
    for _, doc := range docs{
       //doc[createTime] = time.Now().Unix()
       //doc["updateTime"] = time.Now().Unix()
       err = b.Source.DB().C(colName).Insert(doc)
    }
    if err != nil{
        b.Log.Error(err)
        return err
    }
    return nil
}

//更新一条数据
func (b *BaseRepository) Update(colName string, selector interface{}, update interface{}) error{
    err := b.Source.DB().Session.Ping()
    if err != nil{
        b.Log.Errorf("数据库连接失败: %v", err)
        return err
    }
    err = b.Source.DB().C(colName).Update(selector, update)
    if err != nil{
        b.Log.Error(err)
        return err
    }
    return nil
}

//更新多条数据
func (b *BaseRepository) UpdateAll(colName string, selector interface{}, update interface{}) error{
    err := b.Source.DB().Session.Ping()
    if err != nil{
        b.Log.Errorf("数据库连接失败: %v", err)
        return err
    }
    _, err = b.Source.DB().C(colName).UpdateAll(selector, update)
    if err != nil{
        b.Log.Error(err)
        return err
    }
    return nil
}

//删除一条数据
func (b *BaseRepository) Delete(colName string, selector interface{}) error{
    err := b.Source.DB().Session.Ping()
    if err != nil{
        b.Log.Errorf("数据库连接失败: %v", err)
        return err
    }
    err = b.Source.DB().C(colName).Remove(selector)
    if err != nil{
        b.Log.Error(err)
        return err
    }
    return nil
}

//根据id获取一条数据
func (b *BaseRepository) DeleteById(colName string, id string) error{
    err := b.Source.DB().Session.Ping()
    if err != nil{
        b.Log.Errorf("数据库连接失败: %v", err)
        return err
    }
    err = b.Source.DB().C(colName).RemoveId(bson.ObjectIdHex(id))
    if err != nil{
        b.Log.Error(err)
        return err
    }
    return nil
}

//删除多条数据
func (b *BaseRepository) DeleteAll(colName string, selector interface{}) error{
    err := b.Source.DB().Session.Ping()
    if err != nil{
        b.Log.Errorf("数据库连接失败: %v", err)
        return err
    }
    _, err = b.Source.DB().C(colName).RemoveAll(selector)
    if err != nil{
        b.Log.Error(err)
        return err
    }
    return nil
}