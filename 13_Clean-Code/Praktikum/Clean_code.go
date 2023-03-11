package main

type user struct {
	id       int
	username int
	password int
}

type userservice struct {
	t []user
}

func (u userservice) getallusers() []user {
	return u.t
}

func (u userservice) getuserbyid(id int) user {
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}

	return user{}
}

//Kekurangan dari Code diatas adalah :
//Kekurangan komentar: Tidak ada komentar di dalam kode, sehingga sulit bagi pengembang lain untuk memahami apa yang dilakukan oleh kode tersebut.
//Konvensi penamaan tidak konsisten: Konvensi penamaan untuk struct user dan struct userservice tidak konsisten. user memiliki nama field dengan huruf kecil sementara userservice memiliki nama field dengan huruf besar.
//Spasi tidak konsisten: Terdapat ketidak konsistenan dalam penggunaan spasi di seluruh kode. Misalnya, tidak ada spasi di sekitar operator = pada deklarasi struct userservice, tetapi ada spasi di sekitar operator = pada deklarasi fungsi getallusers.
//Penamaan variabel buruk: Nama variabel pada struct user tidak cukup deskriptif. id, username, dan password tidak memberikan informasi yang cukup tentang apa yang direpresentasikan.
