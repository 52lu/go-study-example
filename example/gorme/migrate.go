/**
 * @Description 测试使用gorm迁移
 **/
package gorme

// 自动迁移schema,(根据结构体创建或者更新schema)
func GormAutoMigrate(host, port, use, pass, database string) error {
	mysqlByDefault, err := ConnectMysqlByDefault(host, port, use, pass, database)
	if err != nil {
		return err
	}
	// 指定索引创建
	err = mysqlByDefault.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
	if err != nil {
		return err
	}
	return nil
}

