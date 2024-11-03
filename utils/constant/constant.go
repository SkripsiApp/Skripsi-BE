package constant

// Constanta For Role
const (
	USER       = "user"
	ADMIN      = "admin"
	SUPERADMIN = "super_admin"
)

// Constanta For Success
const (
	SUCCESS_LOGIN       = "berhasil melakukan login"
	SUCCESS_NULL        = "data belum tersedia"
	SUCCESS_CREATE_DATA = "berhasil membuat data"
	SUCCESS_DELETE_DATA = "berhasil menghapus data"
	SUCCESS_GET_DATA    = "berhasil mendapatkan data"
)

// Constanta For Error
const (
	ERROR_TEMPLATE         = "gagal menguraikan template"
	ERROR_DATA_ID          = "id tidak ditemukan"
	ERROR_ID_INVALID       = "id salah"
	ERROR_DATA_EMAIL       = "email tidak ditemukan"
	ERROR_FORMAT_EMAIL     = "format email tidak valid"
	ERROR_EMAIL_EXIST      = "email sudah digunakan"
	ERROR_USERNAME_EXIST   = "username sudah digunakan"
	ERROR_AKSES_ROLE       = "akses ditolak"
	ERROR_PASSWORD         = "password lama tidak sesuai"
	ERROR_CONFIRM_PASSWORD = "konfirmasi password tidak sesuai"
	ERROR_EXTRA_TOKEN      = "gagal ekstrak token"
	ERROR_ID_ROLE          = "id atau role tidak ditemukan"
	ERROR_GET_DATA         = "data tidak ditemukan"
	ERROR_EMPTY            = "harap lengkapi data dengan benar"
	ERROR_EMPTY_FILE       = "tidak ada file yang di upload"
	ERROR_HASH_PASSWORD    = "hash password"
	ERROR_DATA_NOT_FOUND   = "data tidak ditemukan"
	ERROR_DATA_EXIST       = "data sudah ada"
	ERROR_MISSION_LIMIT    = "tahapan misi tidak boleh dari 5"
	ERROR_INVALID_TITLE    = "error: tahapan misi tidak boleh sama"
	ERROR_INVALID_ID       = "error: id tidak boleh sama"
	ERROR_INVALID_UPDATE   = "error: data harus berberbeda dengan data sebelumnya"
	ERROR_INVALID_INPUT    = "data yang diinput tidak sesuai"
	ERROR_NOT_FOUND        = "data tidak ditemukan"
	ERROR_RECORD_NOT_FOUND = "record not found"
	ERROR_INVALID_TYPE     = "berupa angka"
	ERROR_INVALID_STATUS   = "status tidak valid"
	ERROR_LIMIT            = "limit tidak boleh lebih dari 10"
	ERROR_LENGTH_PASSWORD  = "minimal 6 karakter, ulangi kembali"
)
