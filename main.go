package main

import (
	"belajar-golang-web/model"
	"fmt"
	"os"

	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	var user1 = &model.User{
		Id:       "u001",
		Name:     "Achmad Hadzami Setiawan",
		Password: "zem10",
		Gender:   model.UserGender_MALE,
	}

	var userList = &model.UserList{
		ListUser: []*model.User{
			user1,
		},
	}

	var garage1 = &model.Garage{
		Id:   "g001",
		Name: "Kalimdor",
		Coordinate: &model.GarageCoordinate{
			Latitude:  23.2212847,
			Longitude: 53.22033123,
		},
	}

	var garageList = &model.GarageList{
		List: []*model.Garage{
			garage1,
		},
	}

	var garageListByUser = &model.GarageListByUser{
		List: map[string]*model.GarageList{
			user1.Id: garageList,
		},
	}

	fmt.Printf("# ==== Original\n       %#v \n", user1)
	fmt.Printf("# ==== As String\n       %s \n", user1.String())
	fmt.Printf("# ==== Original\n       %#v \n", garageListByUser)
	fmt.Printf("# ==== Original\n       %#v \n", userList)

	jsonb, err := protojson.Marshal(garageList)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	fmt.Printf("# ==== As JSON String\n       %s \n", string(jsonb))

	protoObject := &model.GarageList{}
	err = protojson.Unmarshal(jsonb, protoObject)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	fmt.Printf("# ==== As String\n       %s \n", protoObject.String())
}
