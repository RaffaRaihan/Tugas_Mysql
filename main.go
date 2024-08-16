package main

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

type User struct {
    ID       int
    Name     string
    Email    string
    Password string
    Gender   string
    Photo    string 
    Address  string 
    Role     string
}

type Role struct {
    ID   int
    Name string
}

type AC struct {
    ID    int    
    Name  string
    Brand string
    PK    float64
    Price float64
}

type Service struct {
    ID            int      
    TechnicianID  int      
    ClientID      int      
    ACID          int      
    Date          string
    Status        string    
}

func main() {
    // Ganti dengan data koneksi MySQL Anda
    dsn := "root:@tcp(127.0.0.1:3306)/tugas_1?charset=utf8mb4&parseTime=True&loc=Local"

    // Membuat koneksi ke database
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Gagal konek ke database")
    }

    fmt.Println("Berhasil konek ke database")

    // Gunakan `db` untuk berinteraksi dengan database
    // Data dasar yang akan diulang
    baseUsers := []User{
        {Name: "Fulan", Email: "fulan@gmail.com", Password: "******", Gender: "L", Photo: "https://lorem.ipsum/dolor.png", Address: "Jl. Cisitu Indah VI no 6", Role: "1"},
        {Name: "Fulanah", Email: "fulanah@gmail.com", Password: "******", Gender: "P", Photo: "https://lorem.ipsum/dolor.png", Address: "Jl. Cisitu Indah VI no 6", Role: "2"},
        {Name: "Ardi", Email: "ardi@gmail.com", Password: "******", Gender: "L", Photo: "https://lorem.ipsum/dolor.png", Address: "Jl. Cisitu Indah VI no 6", Role: "2"},
        {Name: "Samsudin", Email: "samsudin@gmail.com", Password: "******", Gender: "L", Photo: "https://lorem.ipsum/dolor.png", Address: "Jl. Cisitu Indah VI no 6", Role: "2"},
        {Name: "Eko", Email: "eko@gmail.com", Password: "******", Gender: "L", Photo: "https://lorem.ipsum/dolor.png", Address: "Jl. Cisitu Indah VI no 6", Role: "2"},
        {Name: "Sugeng", Email: "sugeng@gmail.com", Password: "******", Gender: "L", Photo: "https://lorem.ipsum/dolor.png", Address: "Jl. Cisitu Indah VI no 6", Role: "3"},
        {Name: "Alif", Email: "alif@gmail.com", Password: "******", Gender: "L", Photo: "https://lorem.ipsum/dolor.png", Address: "Jl. Cisitu Indah VI no 6", Role: "3"},
        {Name: "Siti", Email: "siti@gmail.com", Password: "******", Gender: "P", Photo: "https://lorem.ipsum/dolor.png", Address: "Jl. Cisitu Indah VI no 6", Role: "3"},
        {Name: "Juminten", Email: "juminten@gmail.com", Password: "******", Gender: "P", Photo: "https://lorem.ipsum/dolor.png", Address: "Jl. Cisitu Indah VI no 6", Role: "3"},
        {Name: "Paijo", Email: "paijo@gmail.com", Password: "******", Gender: "L", Photo: "https://lorem.ipsum/dolor.png", Address: "Jl. Cisitu Indah VI no 6", Role: "3"},
        {Name: "Saifuddin", Email: "saifuddin@gmail.com", Password: "******", Gender: "L", Photo: "https://lorem.ipsum/dolor.png", Address: "Jl. Cisitu Indah VI no 6", Role: "3"},
        {Name: "Daffa", Email: "daffa@gmail.com", Password: "******", Gender: "L", Photo: "https://lorem.ipsum/dolor.png", Address: "Jl. Cisitu Indah VI no 6", Role: "3"},
        {Name: "Akbar", Email: "akbar@gmail.com", Password: "******", Gender: "L", Photo: "https://lorem.ipsum/dolor.png", Address: "Jl. Cisitu Indah VI no 6", Role: "3"},
        {Name: "Rafli", Email: "rafli@gmail.com", Password: "******", Gender: "L", Photo: "https://lorem.ipsum/dolor.png", Address: "Jl. Cisitu Indah VI no 6", Role: "3"},
        {Name: "Rini", Email: "rini@gmail.com", Password: "******", Gender: "P", Photo: "https://lorem.ipsum/dolor.png", Address: "Jl. Cisitu Indah VI no 6", Role: "3"},
    }

    // Insert data menggunakan loop
    for i := 0; i < len(baseUsers); i++ {
        user := baseUsers[i]
        if err := db.Create(&user).Error; err != nil {
            fmt.Printf("Gagal menyimpan user %s: %v\n", user.Name, err)
        } else {
            fmt.Printf("Berhasil menyimpan user %s\n", user.Name)
        }
    }
	// Data dasar yang akan diulang
    roles := []Role{
        {Name: "Admin"},
        {Name: "Techmician"},
        {Name: "Client"},
    }

    // Insert data menggunakan loop
    for i := 0; i < len(roles); i++ {
        role := roles[i]
        if err := db.Create(&role).Error; err != nil {
            fmt.Printf("Gagal menyimpan role %s: %v\n", role.Name, err)
        } else {
            fmt.Printf("Berhasil menyimpan role %s\n", role.Name)
        }
    }
	// Data dasar yang akan diinsert
    ac := []AC{
        {ID: 1, Name: "LG-1", Brand: "LG", PK: 0.5, Price: 50000},
        {ID: 2, Name: "Sharp-2", Brand: "Sharp", PK: 1, Price: 60000},
        {ID: 3, Name: "Panasonic-3", Brand: "Panasonic", PK: 2, Price: 70000},
        {ID: 4, Name: "Samsung-4", Brand: "Samsung", PK: 0.5, Price: 80000},
        {ID: 5, Name: "Daikin-5", Brand: "Daikin", PK: 1, Price: 90000},
        {ID: 6, Name: "Gree-6", Brand: "Gree", PK: 2, Price: 100000},
        {ID: 7, Name: "Polytron-7", Brand: "Polytron", PK: 0.5, Price: 110000},
        {ID: 8, Name: "Electrolux-8", Brand: "Electrolux", PK: 1, Price: 120000},
        {ID: 9, Name: "Aqua-9", Brand: "Aqua", PK: 2, Price: 130000},
        {ID: 10, Name: "Midea-10", Brand: "Midea", PK: 0.5, Price: 140000},
        {ID: 11, Name: "LG-11", Brand: "LG", PK: 1, Price: 200000},
        {ID: 12, Name: "Sharp-12", Brand: "Sharp", PK: 2, Price: 210000},
        {ID: 13, Name: "Panasonic-13", Brand: "Panasonic", PK: 0.5, Price: 220000},
        {ID: 14, Name: "Samsung-14", Brand: "Samsung", PK: 1, Price: 230000},
        {ID: 15, Name: "Daikin-15", Brand: "Daikin", PK: 2, Price: 240000},
        {ID: 16, Name: "Gree-16", Brand: "Gree", PK: 0.5, Price: 250000},
        {ID: 17, Name: "Polytron-17", Brand: "Polytron", PK: 1, Price: 260000},
        {ID: 18, Name: "Electrolux-18", Brand: "Electrolux", PK: 2, Price: 270000},
        {ID: 19, Name: "Aqua-19", Brand: "Aqua", PK: 0.5, Price: 280000},
        {ID: 20, Name: "Midea-20", Brand: "Midea", PK: 1, Price: 290000},
    }

    // Insert data menggunakan loop
    for _, ac := range ac {
        if err := db.Create(&ac).Error; err != nil {
            fmt.Printf("Gagal menyimpan AC %s: %v\n", ac.Name, err)
        } else {
            fmt.Printf("Berhasil menyimpan AC %s\n", ac.Name)
        }
    }
    // Data dasar yang akan diinsert
    services := []Service{
        {ID: 1, TechnicianID: 2, ClientID: 6, ACID: 1, Date:"1 Jun 2020", Status: "finish"},
        {ID: 2, TechnicianID: 3, ClientID: 7, ACID: 2, Date:"1 May 2020", Status: "finish"},
        {ID: 3, TechnicianID: 4, ClientID: 8, ACID: 3, Date:"2 Jun 2020", Status: "finish"},
        {ID: 4, TechnicianID: 5, ClientID: 9, ACID: 4, Date:"3 March 2021", Status: "finish"},
        {ID: 5, TechnicianID: 2, ClientID: 6, ACID: 5, Date:"5 Dec 2021", Status: "finish"},
        {ID: 6, TechnicianID: 3, ClientID: 7, ACID: 6, Date:"25 Dec 2021", Status: "finish"},
        {ID: 7, TechnicianID: 4, ClientID: 10, ACID: 7, Date:"1 Jan 2022", Status: "finish"},
        {ID: 8, TechnicianID: 5, ClientID: 11, ACID: 8, Date:"2 Feb 2022", Status: "finish"},
        {ID: 9, TechnicianID: 2, ClientID: 6, ACID: 9, Date:"4 Apr 2022", Status: "finish"},
        {ID: 10, TechnicianID: 3, ClientID: 7, ACID: 10, Date:"5 May 2023", Status: "on repair"},
        {ID: 11, TechnicianID: 4, ClientID: 12, ACID: 11, Date:"6 Jun 2023", Status: "on repair"},
        {ID: 12, TechnicianID: 5, ClientID: 13, ACID: 12, Date:"7 Jul 2023", Status: "on repair"},
        {ID: 13, TechnicianID: 2, ClientID: 6, ACID: 13, Date:"8 Aug 2023", Status: "paid"},
        {ID: 14, TechnicianID: 3, ClientID: 7, ACID: 14, Date:"9 Sep 2023", Status: "paid"},
        {ID: 15, TechnicianID: 4, ClientID: 14, ACID: 15, Date:"10 Oct 2023", Status: "unpaid"},
    }

    // Insert data menggunakan loop
    for _, service := range services {
        if err := db.Create(&service).Error; err != nil {
            fmt.Printf("Gagal menyimpan service ID %d: %v\n", service.ID, err)
        } else {
            fmt.Printf("Berhasil menyimpan service ID %d\n", service.ID)
        }
    }
    // Menghitung jumlah perbaikan yang sudah dibayar
    var paidRepairCount int64
    result := db.Model(&Service{}).Where("status = ?", "paid").Count(&paidRepairCount)
    if result.Error != nil {
        fmt.Printf("Gagal menghitung jumlah perbaikan yang dibayar: %v\n", result.Error)
    } else {
        fmt.Printf("Jumlah perbaikan yang sudah dibayar: %d\n", paidRepairCount)
    }
    // Update status pembayaran menjadi "paid" untuk client_id = 7
    result = db.Model(&Service{}).
        Where("client_id = ?", 7).
        Update("status", "paid")

    if result.Error != nil {
        fmt.Printf("Gagal memperbarui status: %v\n", result.Error)
    } else {
        fmt.Printf("Berhasil memperbarui status untuk client_id = 7\n")
    }
    // Menghapus semua data dari tabel services
    result = db.Delete(&Service{}, "1 = 1")

    if result.Error != nil {
        fmt.Printf("Gagal menghapus data: %v\n", result.Error)
    } else {
        fmt.Printf("Berhasil menghapus semua data dari tabel services\n")
    }
}
