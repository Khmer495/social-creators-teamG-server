package infrastracture

import "github.com/Khmer495/social-creators-teamG-server/go/entity/model"

var Prefectures = model.Prefectures{
	model.Prefecture{
		ID:    1,
		Name:  "北海道",
		Order: 1,
	},
	model.Prefecture{
		ID:    26,
		Name:  "京都府",
		Order: 26,
	},
}

var Cities = model.Cities{
	model.City{
		ID:           1,
		PrefectureID: 1,
		Name:         "札幌",
		Order:        1,
	},
	model.City{
		ID:           2,
		PrefectureID: 26,
		Name:         "京都市役所",
		Order:        1,
	},
	model.City{
		ID:           3,
		PrefectureID: 26,
		Name:         "河原町",
		Order:        2,
	},
	model.City{
		ID:           4,
		PrefectureID: 26,
		Name:         "四条烏丸",
		Order:        1,
	},
}

var Users = model.Users{
	model.User{
		ID:         1,
		FamilyName: "久米",
		FirstName:  "祐貴",
	},
}

var Restaurants = model.Restaurants{
	model.Restaurant{
		ID:                  1,
		Name:                "レストラン&カフェ サラオ",
		Address:             "京都府京都市中京区 御幸町通御池上る亀屋町379-1 コンフォール御池フォルテ1F",
		Lat:                 35.011424,
		Lng:                 135.76663,
		PrefectureID:        26,
		CityID:              2,
		OpenTime:            "11:30:00",
		CloseTime:           "21:30:00",
		FromKidsWelcomeTime: "15:00:00",
		ToKidsWelcomeTime:   "17:30:00",
		Menu:                "ランチ：1500円、ディナー：2500円",
		IsOkBaby:            true,
		IsOkInfant:          true,
		IsNoSmoking:         true,
		IsOkBabyCar:         true,
		IsZashiki:           true,
		IsPrivateRoom:       true,
		IsParkingArea:       true,
		IsLuggageStorage:    true,
		IsNapSpace:          true,
		IsBabyBed:           true,
		IsKidsMenu:          true,
		IsKidsChair:         false,
		IsKidsDish:          true,
		IsKidsToilet:        false,
		IsKidsToy:           true,
		IsOkBabyFood:        true,
		IsNursingRoom:       true,
		IsKidsRoom:          false,
		IsAllergenLabel:     true,
	},
	model.Restaurant{
		Name:                "mumokuteki　cafe&foods",
		Address:             "京都府京都市中京区式部町２６１ ヒューマンフォーラム本社ビル 2F",
		Lat:                 35.0065569,
		Lng:                 135.766561,
		PrefectureID:        26,
		CityID:              3,
		OpenTime:            "11:30:00",
		CloseTime:           "22:00:00",
		FromKidsWelcomeTime: "15:00:00",
		ToKidsWelcomeTime:   "18:00:00",
		Menu:                "ランチ：1150円、ディナー：1400~1700円",
		IsOkBabyCar:         true,
		IsZashiki:           false,
		IsPrivateRoom:       false,
		IsParkingArea:       false,
		IsLuggageStorage:    true,
		IsNapSpace:          true,
		IsBabyBed:           true,
		IsKidsMenu:          true,
		IsKidsChair:         true,
		IsKidsDish:          true,
		IsKidsToilet:        false,
		IsKidsToy:           true,
		IsOkBabyFood:        true,
		IsNursingRoom:       true,
		IsKidsRoom:          true,
	},
	model.Restaurant{
		Name:                "Daniel’s　Sole　ダニエルズソーレ",
		Address:             "京都府京都市中京区高倉通錦小路上ル貝屋町567",
		Lat:                 35.0054629,
		Lng:                 135.7626743,
		PrefectureID:        26,
		CityID:              4,
		OpenTime:            "11:30:00",
		CloseTime:           "23:00:00",
		FromKidsWelcomeTime: "14:00:00",
		ToKidsWelcomeTime:   "17:30:00",
		Menu:                "ランチ：1300円、ディナー：3500円",
		IsOkBabyCar:         true,
		IsZashiki:           false,
		IsPrivateRoom:       false,
		IsParkingArea:       false,
		IsLuggageStorage:    false,
		IsNapSpace:          false,
		IsBabyBed:           false,
		IsKidsMenu:          true,
		IsKidsChair:         true,
		IsKidsDish:          true,
		IsKidsToilet:        true,
		IsKidsToy:           false,
		IsOkBabyFood:        true,
		IsNursingRoom:       false,
		IsKidsRoom:          false,
	},
	model.Restaurant{
		Name:                "ロティー チキンアンド ジャッキー タコス 烏丸高辻",
		Address:             "京都府京都市下京区高辻通烏丸東入ル匂天神町642-3",
		Lat:                 35.000432,
		Lng:                 135.760186,
		PrefectureID:        26,
		CityID:              4,
		OpenTime:            "11:00:00",
		CloseTime:           "23:00:00",
		FromKidsWelcomeTime: "15:00:00",
		ToKidsWelcomeTime:   "17:00:00",
		Menu:                "ランチ：1000円、ディナー：2000円、宴会：3000円",
		IsOkBabyCar:         true,
		IsZashiki:           false,
		IsPrivateRoom:       false,
		IsParkingArea:       false,
		IsLuggageStorage:    false,
		IsNapSpace:          false,
		IsBabyBed:           false,
		IsKidsMenu:          false,
		IsKidsChair:         false,
		IsKidsDish:          true,
		IsKidsToilet:        true,
		IsKidsToy:           false,
		IsOkBabyFood:        true,
		IsNursingRoom:       false,
		IsKidsRoom:          false,
	},
}

var Comments = model.Comments{
	model.Comment{
		UserID:       1,
		RestaurantID: 1,
		Text:         "おすすめです",
	},
	model.Comment{
		UserID:       1,
		RestaurantID: 2,
		Text:         "よかったです",
	},
	model.Comment{
		UserID:       1,
		RestaurantID: 3,
		Text:         "また来たいです",
	},
}

var Goods = model.Goods{
	model.Good{
		UserID:       1,
		RestaurantID: 1,
	},
	model.Good{
		UserID:       1,
		RestaurantID: 2,
	},
}

var Favorites = model.Favorites{
	model.Favorite{
		UserID:       1,
		RestaurantID: 1,
	},
	model.Favorite{
		UserID:       1,
		RestaurantID: 3,
	},
}
