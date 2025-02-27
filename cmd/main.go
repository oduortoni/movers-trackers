package main

import (
	"farmers/db"
	"farmers/files"
	"farmers/datatypes"
	"farmers/server"
)

func init() {
	app_data := datatypes.Application{
		Gid: 1,
		Root: "storage",
		Database: "storage/vms.db",
	}
	files.SetConfig(&app_data)
	files.CreateFile("storage/logfile")

	db.CreateDB(files.GetConfig().Database)
}

func main() {
	server.StartServer(9000)
}
/**
	member := datatypes.Member{
		Group: "grouping",
		Location: "Nambale",
		Produce: "Sugar",
	}
	farmer := datatypes.Farmer{
		First: "Papa",
		Second: "Hulwa",
	}
	farmer2 := datatypes.Farmer{
		First: "Ololwle",
		Second: "Majorie",
	}
	var slice []datatypes.Farmer
	slice = append(slice, farmer)
	slice = append(slice, farmer2)
	
	db.InsertMember(member, slice)

	users, _ := db.GetUsers()
	log.Println(users)
}
/**
	data := files.GetConfig()
	log.Println(data)
	data.Gid = 21
	files.SetConfig(data)
	data = files.GetConfig()
	log.Println(data)
}*/
/**
	data := files.GetConfig()
	log.Printf("{\n%d, %s, %s\n}\n", data.Gid, data.Root, data.Database)
	_, err := os.OpenFile("hello", os.O_APPEND, 0777)
	files.CheckError(err)
	files.CheckError(errors.New("useless"))
	files.CheckError(errors.New("Am an error"))
}*/

