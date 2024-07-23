package Carservice

func (s *CarService) IsExist(id uint64) (bool, error) {

	var isExist bool

	sql := `
			select  
    		count(*) > 0  
			from cars 
			where id = ?;`

	err := s.Db.Raw(sql, id).Scan(&isExist).Error

	if err != nil {
		return false, err
	}

	return isExist, nil

}
