package service

import (
	"Gin_Project/go/dao"
	"Gin_Project/go/entity"
	"fmt"
	"strings"
)

// 新建spot信息
func CreateSpot(spot *entity.Spot) (err error) {
	if err = dao.UserSession.Create(spot).Error; err != nil {
		return err
	}
	return
}

// 获取spot集合
func GetAllSpot() (spotList []*entity.Spot, err error) {
	if err := dao.UserSession.Find(&spotList).Error; err != nil {
		return nil, err
	}
	return
}

// 景区排序集合
func GetSpotByScore() (spotList []*entity.Spot, err error) {
	if err := dao.UserSession.Order("score desc").Find(&spotList).Error; err != nil {
		return nil, err
	}
	return
}

// 依据地区筛选得到景区集合
func FilterByAddress(address string) (spotList []*entity.Spot, err error) {
	if err := dao.UserSession.Where("address=?", address).Order("score desc").Find(&spotList).Error; err != nil {
		return nil, err
	}
	return
}

// 依据类型筛选得到景区集合
func FilterByType(t string) (spotList []*entity.Spot, err error) {
	if err := dao.UserSession.Where("type=?", t).Order("score desc").Find(&spotList).Error; err != nil {
		return nil, err
	}
	return
}

// 依据市级地址筛选得到景区集合
func FilterByAddress1(t string) (spotList []*entity.Spot, err error) {
	if err := dao.UserSession.Where("address=?", t).Order("score desc").Find(&spotList).Error; err != nil {
		return nil, err
	}
	return
}

// 依据名字筛选得到景区集合
func FilterByName(name string) (spotList []*entity.Spot, err error) {
	if err := dao.UserSession.Where("name=?", name).Order("score desc").Find(&spotList).Error; err != nil {
		return nil, err
	}
	return
}

// 依据名字筛选得到景区集合
func Filter(ttype string, address string) (spotList []*entity.Spot, err error) {
	a := []string{""}
	address = strings.Trim(address, "[]")       // 去除方括号
	addressSlice := strings.Split(address, ",") // 拆分字符串为切片
	for index, value := range addressSlice {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
		a = append(a, value[0:len(value)-1])
	}
	if err := dao.UserSession.Where("type = ? and address IN (?)", ttype, a).Order("score desc").Find(&spotList).Error; err != nil {
		return nil, err
	}
	return
}

// 依据名字筛选得到景区集合
func Filter1(ttype string, address string) (spotList []*entity.Spot, err error) {
	var al int
	var tl int

	a := []string{""}
	address = strings.Trim(address, "[]")       // 去除方括号
	addressSlice := strings.Split(address, ",") // 拆分字符串为切片
	for index, value := range addressSlice {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
		a = append(a, value[0:len(value)-1])
		al = index + 1
		if index == 0 && value == "" {
			al = 0
		}
	}
	t := []string{""}
	ttype = strings.Trim(ttype, "[]")      // 去除方括号
	typeSlice := strings.Split(ttype, ",") // 拆分字符串为切片
	for index, value := range typeSlice {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
		t = append(t, value)
		tl = index + 1
		if index == 0 && value == "" {
			tl = 0
		}
	}

	if tl != 0 && al != 0 {
		if err := dao.UserSession.Where("type IN (?) and address IN (?)", t, a).Order("score desc").Find(&spotList).Error; err != nil {
			return nil, err
		}
		return
	}
	if tl == 0 && al != 0 {
		if err := dao.UserSession.Where("address IN (?)", a).Order("score desc").Find(&spotList).Error; err != nil {
			return nil, err
		}
		return
	}
	if tl != 0 && al == 0 {
		if err := dao.UserSession.Where("type IN (?)", t).Order("score desc").Find(&spotList).Error; err != nil {
			return nil, err
		}
		return
	}
	if err := dao.UserSession.Find(&spotList).Error; err != nil {
		return nil, err
	}
	return
}

// 依据名字模糊搜索得到景区集合
func Search(name string) (spotList []*entity.Spot, err error) {
	if err := dao.UserSession.Where("name LIKE ? or address LIKE ? or address1 LIKE ?", "%"+name+"%", "%"+name+"%", "%"+name+"%").Order("score desc").Find(&spotList).Error; err != nil {
		return nil, err
	}
	return
}

// 依据名字模糊搜索得到景区集合
func SearchByName(name string) (spotList []*entity.Spot, err error) {
	if err := dao.UserSession.Where("name LIKE ?", "%"+name+"%").Order("score desc").Find(&spotList).Error; err != nil {
		return nil, err
	}
	return
}

// 根据id查询景点spot
func GetSpotById(id int) (spot entity.Spot, err error) {
	if err := dao.UserSession.Where("id=?", id).First(&spot).Error; err != nil {
		return spot, err
	}
	return
}

// 根据id删除spot
func DeleteSpotById(id string) (err error) {
	err = dao.UserSession.Where("id=?", id).Delete(&entity.Spot{}).Error
	return
}

// 更新景点信息
func UpdateSpot(spot *entity.Spot) (err error) {
	err = dao.UserSession.Save(spot).Error
	return
}
