package models

type Message struct {
	Th string `json:"th"`
	En string `json:"en"`
}

type Constants struct {
	Result  bool    `json:"result"`
	Message Message `json:"message"`
}

func Get_data_error() Constants {
	return Constants{
		false,
		Message{
			"ส่งข้อมูลผิดพลาด",
			"Sent data error!",
		},
	}
}

func Get_data_success() Constants {
	return Constants{
		true,
		Message{
			"ส่งข้อมูลสำเร็จ",
			"Sent data success!",
		},
	}
}

func Create_user_success() Constants {
	return Constants{
		true,
		Message{
			"สร้างผู้ใช้งานสำเร็จ",
			"Create user success!",
		},
	}
}

func Check_ID_fail() Constants {
	return Constants{
		false,
		Message{
			"มี ID นี้อยู่ในระบบแล้ว",
			"Has this ID in sysytem",
		},
	}
}

func Create_user_fail() Constants {
	return Constants{
		false,
		Message{
			"สร้างผู้ใช้งานไม่สำเร็จ",
			"Create user fail!",
		},
	}
}

func Delete_user_success() Constants {
	return Constants{
		true,
		Message{
			"ลบผู้ใช้งานสำเร็จ",
			"Delete user success!",
		},
	}
}

func Delete_user_fail() Constants {
	return Constants{
		false,
		Message{
			"ลบผู้ใช้งานไม่สำเร็จ",
			"Delete user fail!",
		},
	}
}

func Update_user_success() Constants {
	return Constants{
		true,
		Message{
			"อัพเดทผู้ใช้งานสำเร็จ",
			"Update user success!",
		},
	}
}

func Update_user_fail() Constants {
	return Constants{
		false,
		Message{
			"อัพเดทผู้ใช้งานไม่สำเร็จ",
			"Update user fail!",
		},
	}
}
