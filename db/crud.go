package db

import (
	"errors"
	"strconv"

	"farmers/datatypes"
	"farmers/files"
)

func InsertMember(member datatypes.Member, farmers []datatypes.Farmer) {
	gid := InsertGroup(member)
	InsertFarmers(farmers, gid)
}

func InsertGroup(member datatypes.Member) string {
	gid := GetGid(member.Location)
	driver, err := OpenDatabase()
	files.CheckError(err)
	defer driver.Close()
	sttmt, err := driver.Prepare("INSERT INTO members (id,membership,location,produce,gid) VALUES (?,?,?,?,?)")
	files.CheckError(err)
	sttmt.Exec(nil, member.Group, member.Location, member.Produce, gid)
	defer sttmt.Close()

	return gid
}

func InsertFarmers(farmers []datatypes.Farmer, gid string) {
	driver, err := OpenDatabase()
	files.CheckError(err)
	defer driver.Close()

	sql_farmer := `INSERT INTO farmers(id, first, second, gid) VALUES(?, ?, ?, ?);`
	sttmt, err := driver.Prepare(sql_farmer)
	files.CheckError(err)

	for _, farmer := range farmers {
		sttmt.Exec(nil, farmer.First, farmer.Second, gid)
	}
}

func GetUsers() ([]datatypes.Farmer, error) {
	db, err := OpenDatabase()
	if err != nil {
		return nil, errors.New("Failed to open db")
	}
	sql_all := `SELECT * FROM farmers`
	rows, err := db.Query(sql_all)
	if err != nil {
		return nil, errors.New("Failed to find users")
	}
	defer rows.Close()

	var farmers []datatypes.Farmer

	for rows.Next() {
		var farmer datatypes.Farmer
		if err := rows.Scan(&farmer.Id, &farmer.First, &farmer.Second, &farmer.Gid); err != nil {
			return farmers, err
		}
		farmers = append(farmers, farmer)

		if err = rows.Err(); err != nil {
			return farmers, err
		}
	}
	return farmers, nil
}

func UpdateCharges(charges datatypes.Charges) {
	driver, err := OpenDatabase()
	files.CheckError(err)

	sql := "INSERT INTO charges(means, load, cost) VALUES(?, ?, ?);"
	sttmt, err := driver.Prepare(sql)
	files.CheckError(err)
	sttmt.Exec(charges.Means, charges.Load, charges.Cost)
}

func GetCharges() ([]datatypes.Charges, error) {
	db, err := OpenDatabase()
	if err != nil {
		return nil, errors.New("Failed to open db")
	}
	sql_all := `SELECT * FROM charges`
	rows, err := db.Query(sql_all)
	if err != nil {
		return nil, errors.New("Failed to find charges")
	}
	defer rows.Close()

	var charges []datatypes.Charges

	for rows.Next() {
		var charge datatypes.Charges
		if err := rows.Scan(&charge.Means, &charge.Load, &charge.Cost); err != nil {
			return charges, err
		}
		charges = append(charges, charge)

		if err = rows.Err(); err != nil {
			return charges, err
		}
	}
	return charges, nil
}

func GetGid(location string) string {
	data := files.GetConfig()
	gid := location
	gid += strconv.Itoa(data.Gid)

	data.Gid = data.Gid + 1
	files.SetConfig(data)

	return gid
}
