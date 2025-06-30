package db

import (
	"Hospital/pkg/constants"
	"Hospital/pkg/crypt"
	"Hospital/pkg/errno"
	"context"
	"time"
)

func CreateUser(ctx context.Context, username, password string, t, id int64) (*User, error) {
	if t == 1 {
		var admin *Admin
		err := DB.WithContext(ctx).
			Table(constants.TableAdmin).
			Where("BINARY admin_id = ?", id).
			First(&admin).
			Error
		if err == nil { //找到了
			return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "user already exists")
		}
		admin = &Admin{
			Username:   username,
			Password:   password,
			CreateTime: time.Now(),
			Status:     1,
		}
		err = DB.WithContext(ctx).
			Table(constants.TableAdmin).
			Create(admin).
			Error
		if err != nil {
			return nil, errno.NewErrNo(errno.InternalDatabaseErrorCode, "create admin error"+err.Error())
		}
		user := &User{
			Id:         id,
			Name:       username,
			CreateTime: admin.CreateTime,
			Identity:   "管理员",
		}
		return user, err
	} else if t == 2 {

		var doctor *Doctor
		err := DB.WithContext(ctx).
			Table(constants.TableDoctor).
			Where("BINARY doctor_id = ?", id).
			First(&doctor).
			Error
		if err == nil { //找到了
			return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "doctor already exists")
		}
		p, _ := crypt.PasswordHash(password)
		doctor = &Doctor{
			Name:       username,
			Password:   p,
			DoctorId:   id,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		}
		err = DB.WithContext(ctx).
			Table(constants.TableDoctor).
			Create(doctor).
			Error
		if err != nil {
			return nil, errno.NewErrNo(errno.InternalDatabaseErrorCode, "create doctor error:"+err.Error())
		}
		user := &User{
			Id:         id,
			Name:       username,
			CreateTime: doctor.CreateTime,
			Identity:   "医生",
		}
		return user, err
	} else if t == 3 {

		var nurse *Nurse
		err := DB.WithContext(ctx).
			Table(constants.TableNurse).
			Where("BINARY nurse_id = ?", id).
			First(&nurse).
			Error
		if err == nil { //找到了
			return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "nurse already exists")
		}
		p, _ := crypt.PasswordHash(password)
		nurse = &Nurse{
			Name:       username,
			Password:   p,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
			NurseId:    id,
		}
		err = DB.WithContext(ctx).
			Table(constants.TableNurse).
			Create(nurse).
			Error
		if err != nil {
			return nil, errno.NewErrNo(errno.InternalDatabaseErrorCode, "create nurse error"+err.Error())
		}
		user := &User{
			Id:         id,
			Name:       username,
			CreateTime: nurse.CreateTime,
			Identity:   "护士",
		}
		return user, err
	}
	return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "invalid type")
}

func UpdateNurseMessage(ctx context.Context, req *UpdateNurseRequest) error {
	var nurse *Nurse
	err := DB.WithContext(ctx).
		Table(constants.TableNurse).
		Where("BINARY nurse_id = ?", req.Id).
		First(&nurse).
		Error
	if err != nil { //找不到了
		return errno.NewErrNo(errno.InternalServiceErrorCode, "nurse no exists")
	}
	if req.Name == "" {
		// 使用Select明确指定要更新的字段
		err = DB.WithContext(ctx).
			Table(constants.TableNurse).Model(&Nurse{}).
			Where("nurse_id = ?", req.Id).
			Select("ContactPhone", "Department", "Position", "UpdateTime"). // 明确选择字段
			Updates(Nurse{
				ContactPhone: req.ContactPhone,
				Department:   req.Department,
				Position:     req.Position,
				UpdateTime:   time.Now(),
			}).Error
		if err != nil {
			return errno.NewErrNo(errno.InternalDatabaseErrorCode, "update nurse error")
		}
	} else {
		// 使用Select明确指定要更新的字段
		err = DB.WithContext(ctx).
			Table(constants.TableNurse).Model(&Nurse{}).
			Where("nurse_id = ?", req.Id).
			Select("ContactPhone", "Department", "Position", "Name", "UpdateTime"). // 明确选择字段
			Updates(Nurse{
				ContactPhone: req.ContactPhone,
				Department:   req.Department,
				Position:     req.Position,
				Name:         req.Name,
				UpdateTime:   time.Now(),
			}).Error
		if err != nil {
			return errno.NewErrNo(errno.InternalDatabaseErrorCode, "update nurse error")
		}
	}
	return nil
}

func UpdateDoctorMessage(ctx context.Context, req *UpdateDoctorRequest) error {
	var doctor *Doctor
	err := DB.WithContext(ctx).
		Table(constants.TableDoctor).
		Where("BINARY doctor_id = ?", req.Id).
		First(&doctor).
		Error
	if err != nil { //找不到了
		return errno.NewErrNo(errno.InternalServiceErrorCode, "Doctor no exists")
	}
	if req.Name == "" {
		// 使用Select明确指定要更新的字段
		err = DB.WithContext(ctx).
			Table(constants.TableDoctor).
			Model(&Doctor{}).
			Where("doctor_id = ?", req.Id).
			Select("ContactPhone", "Department", "Specialty", "Title", "UpdateTime"). // 明确选择字段
			Updates(Doctor{
				ContactPhone: req.ContactPhone,
				Department:   req.Department,
				Specialty:    req.Specialty,
				Title:        req.Title,
				UpdateTime:   time.Now(),
			}).Error
		if err != nil {
			return errno.NewErrNo(errno.InternalDatabaseErrorCode, "update doctor error")
		}
	} else {
		err = DB.WithContext(ctx).
			Table(constants.TableDoctor).
			Model(&Doctor{}).
			Where("doctor_id = ?", req.Id).
			Select("ContactPhone", "Department", "Specialty", "Title", "Name", "UpdateTime"). // 明确选择字段
			Updates(Doctor{
				ContactPhone: req.ContactPhone,
				Department:   req.Department,
				Specialty:    req.Specialty,
				Title:        req.Title,
				Name:         req.Name,
				UpdateTime:   time.Now(),
			}).Error
		if err != nil {
			return errno.NewErrNo(errno.InternalDatabaseErrorCode, "update doctor error")
		}
	}
	return nil
}

func LoginCheck(ctx context.Context, username, password string, t, id int64) (*User, error) {
	if t == 1 {
		var admin *Admin
		err := DB.WithContext(ctx).
			Table(constants.TableAdmin).
			Where("BINARY admin_id = ?", id).
			First(&admin).
			Error
		if err != nil { //找到了
			return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "user not exists")
		}
		if admin.Password != password {
			return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "password error")
		} else {
			return &User{
				Id:           admin.AdminId,
				Name:         admin.Username,
				CreateTime:   admin.CreateTime,
				ContactPhone: admin.ContactPhone,
				Identity:     "管理员",
			}, nil
		}
	} else if t == 2 {
		var doctor *Doctor
		err := DB.WithContext(ctx).
			Table(constants.TableDoctor).
			Where("BINARY doctor_id = ?", id).
			First(&doctor).
			Error
		if err != nil { //找到了
			return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "doctor not exists")
		}
		if !crypt.VerifyPassword(password, doctor.Password) {
			return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "password not match")
		} else {
			return &User{
				Id:           doctor.DoctorId,
				Name:         doctor.Name,
				CreateTime:   doctor.CreateTime,
				Title:        doctor.Title,
				ContactPhone: doctor.ContactPhone,
				Department:   doctor.Department,
				Specialty:    doctor.Specialty,
				Identity:     "医生",
			}, nil
		}
	} else if t == 3 {
		var nurse *Nurse
		err := DB.WithContext(ctx).
			Table(constants.TableNurse).
			Where("BINARY nurse_id = ?", id).
			First(&nurse).
			Error
		if err != nil { //找到了
			return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "nurse not exists")

		}
		if !crypt.VerifyPassword(password, nurse.Password) {
			return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "password not match")
		} else {
			return &User{
				Id:           nurse.NurseId,
				Name:         nurse.Name,
				CreateTime:   nurse.CreateTime,
				ContactPhone: nurse.ContactPhone,
				Position:     nurse.Position,
				Department:   nurse.Department,
				Identity:     "护士",
			}, nil
		}
	}
	return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "invalid type")
}

func GetUserMessage(ctx context.Context, t, id int64) (*User, error) {
	if t == 1 {
		var admin *Admin
		err := DB.WithContext(ctx).
			Table(constants.TableAdmin).
			Where("BINARY admin_id = ?", id).
			First(&admin).
			Error
		if err != nil { //找到了
			return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "user not exists")
		}
		return &User{
			Id:           admin.AdminId,
			Name:         admin.Username,
			CreateTime:   admin.CreateTime,
			ContactPhone: admin.ContactPhone,
			Identity:     "管理员",
		}, nil

	} else if t == 2 {
		var doctor *Doctor
		err := DB.WithContext(ctx).
			Table(constants.TableDoctor).
			Where("BINARY doctor_id = ?", id).
			First(&doctor).
			Error
		if err != nil { //找到了
			return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "doctor not exists")
		}
		return &User{
			Id:           doctor.DoctorId,
			Name:         doctor.Name,
			CreateTime:   doctor.CreateTime,
			Title:        doctor.Title,
			ContactPhone: doctor.ContactPhone,
			Department:   doctor.Department,
			Specialty:    doctor.Specialty,
			Identity:     "医生",
		}, nil

	} else if t == 3 {
		var nurse *Nurse
		err := DB.WithContext(ctx).
			Table(constants.TableNurse).
			Where("BINARY nurse_id = ?", id).
			First(&nurse).
			Error
		if err != nil { //找到了
			return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "nurse not exists")

		}
		return &User{
			Id:           nurse.NurseId,
			Name:         nurse.Name,
			CreateTime:   nurse.CreateTime,
			ContactPhone: nurse.ContactPhone,
			Position:     nurse.Position,
			Department:   nurse.Department,
			Identity:     "护士",
		}, nil
	}
	return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "invalid type")
}

func GetAdminList(ctx context.Context) ([]*User, error) {
	var admins []*Admin
	err := DB.WithContext(ctx).
		Table(constants.TableAdmin).
		Find(&admins).
		Error
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalDatabaseErrorCode, "get admin list error"+err.Error())
	}

	result := make([]*User, 0, len(admins))
	for _, admin := range admins {
		result = append(result, &User{
			Id:           admin.AdminId,
			Name:         admin.Username,
			CreateTime:   admin.CreateTime,
			ContactPhone: admin.ContactPhone,
		})

	}

	return result, nil
}

func GetDoctorList(ctx context.Context) ([]*User, error) {
	var doctors []*Doctor
	err := DB.WithContext(ctx).
		Table(constants.TableDoctor).
		Find(&doctors).
		Error
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalDatabaseErrorCode, "get doctor list error"+err.Error())
	}

	result := make([]*User, 0, len(doctors))
	for _, doctor := range doctors {
		result = append(result, &User{
			Id:           doctor.DoctorId,
			Name:         doctor.Name,
			CreateTime:   doctor.CreateTime,
			ContactPhone: doctor.ContactPhone,
			Title:        doctor.Title,
			Department:   doctor.Department,
			Specialty:    doctor.Specialty,
		})

	}

	return result, nil
}

func GetNurseList(ctx context.Context) ([]*User, error) {
	var nurses []*Nurse
	err := DB.WithContext(ctx).
		Table(constants.TableNurse).
		Find(&nurses).
		Error
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalDatabaseErrorCode, "get nurse list error"+err.Error())
	}

	result := make([]*User, 0, len(nurses))
	for _, nurse := range nurses {
		result = append(result, &User{
			Id:           nurse.NurseId,
			Name:         nurse.Name,
			CreateTime:   nurse.CreateTime,
			ContactPhone: nurse.ContactPhone,
			Position:     nurse.Position,
			Department:   nurse.Department,
		})

	}

	return result, nil
}
