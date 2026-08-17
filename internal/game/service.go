package game

import (
	"context"
	"fmt"
	"os"

	"github.com/minio/minio-go/v7"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	Repository Repository
	minio      *minio.Client
}

func NewService(repo Repository, minio *minio.Client) *Service {
	return &Service{Repository: repo, minio: minio}
}

func (s Service) SeedAll(ctx context.Context) error {
	games := retrieveGames()
	for _, game := range games {
		exists, err := s.Repository.Exists(ctx, game.ID)
		if err != nil {
			return err
		}
		if !exists {
			bucket := os.Getenv("MINIO_BUCKET")
			objectName := fmt.Sprintf("games/icons/%s", game.ID)
			_, err := s.minio.FPutObject(ctx, bucket, objectName, game.IconUrl, minio.PutObjectOptions{ContentType: "application/png"})
			if err != nil {
				return err
			}
			err = s.Repository.Create(ctx, &game)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func mustObjectID(hex string) primitive.ObjectID {
	obj, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		panic(err)
	}
	return obj
}

func retrieveGames() []Game {
	return []Game{
		{ID: mustObjectID("66b8a1010000000000000001"), Name: "Counter-Strike 2", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Counter-Strike_2_logo.svg?width=128"},
		{ID: mustObjectID("66b8a1010000000000000002"), Name: "Minecraft", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Minecraft-creeper-face.svg?width=128"},
		{ID: mustObjectID("66b8a1010000000000000003"), Name: "Roblox", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Roblox_Logo_2022.svg?width=128"},
		{ID: mustObjectID("66b8a1010000000000000004"), Name: "Fortnite", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/FortniteLogo.svg?width=128"},
		{ID: mustObjectID("66b8a1010000000000000005"), Name: "League of Legends", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/League_of_Legends_2019_vector.svg?width=128"},
		{ID: mustObjectID("66b8a1010000000000000006"), Name: "VALORANT", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Valorant_logo_-_pink_color_version.svg?width=128"},
		{ID: mustObjectID("66b8a1010000000000000007"), Name: "The Sims 4", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/The_Sims_4_Logo.svg?width=128"},
		{ID: mustObjectID("66b8a1010000000000000008"), Name: "Overwatch 2", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Overwatch_2_full_logo.svg?width=128"},
		{ID: mustObjectID("66b8a1010000000000000009"), Name: "Rocket League", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Rocket_League_logo.svg?width=128"},
		{ID: mustObjectID("66b8a101000000000000000a"), Name: "Dota 2", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Dota_2.svg?width=128"},
		{ID: mustObjectID("66b8a101000000000000000b"), Name: "PUBG: BATTLEGROUNDS", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/PUBG_logo.svg?width=128"},
		{ID: mustObjectID("66b8a101000000000000000c"), Name: "Grand Theft Auto V", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Grand_Theft_Auto_V_logo.svg?width=128"},
		{ID: mustObjectID("66b8a101000000000000000d"), Name: "Dead by Daylight", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Dead_by_Daylight_Logo.svg?width=128"},
		{ID: mustObjectID("66b8a101000000000000000e"), Name: "Path of Exile 2", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Path_of_Exile_2_logo.svg?width=128"},
		{ID: mustObjectID("66b8a101000000000000000f"), Name: "Apex Legends", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Apex_legends_logo.svg?width=128"},
		{ID: mustObjectID("66b8a1010000000000000010"), Name: "Diablo IV", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Diablo_IV_Logo.svg?width=128"},
		{ID: mustObjectID("66b8a1010000000000000011"), Name: "Call of Duty: Warzone", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Call_of_Duty_Warzone_Logo.svg?width=128"},
		{ID: mustObjectID("66b8a1010000000000000012"), Name: "R.E.P.O.", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/R.E.P.O._logo.svg?width=128"},
		{ID: mustObjectID("66b8a1010000000000000013"), Name: "Rust", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Rust_logo.svg?width=128"},
		{ID: mustObjectID("66b8a1010000000000000014"), Name: "Warframe", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Warframe_logo.svg?width=128"},
		{ID: mustObjectID("66b8a1010000000000000015"), Name: "Tom Clancy's Rainbow Six Siege", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Rainbow_Six_Siege_logo.svg?width=128"},
		{ID: mustObjectID("66b8a1010000000000000016"), Name: "EA Sports FC 26", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/EA_Sports_FC_26_logo.svg?width=128"},
		{ID: mustObjectID("66b8a1010000000000000017"), Name: "Stardew Valley", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Stardew_Valley_Logo.svg?width=128"},
		{ID: mustObjectID("66b8a1010000000000000018"), Name: "Marvel Rivals", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Marvel_Rivals_logo.svg?width=128"},
		{ID: mustObjectID("66b8a1010000000000000019"), Name: "Baldur's Gate 3", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Baldur%27s_Gate_3_logo.svg?width=128"},
		{ID: mustObjectID("66b8a101000000000000001a"), Name: "Elden Ring", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Elden_Ring_logo.svg?width=128"},
		{ID: mustObjectID("66b8a101000000000000001b"), Name: "Palworld", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Palworld_logo.svg?width=128"},
		{ID: mustObjectID("66b8a101000000000000001c"), Name: "Helldivers 2", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Helldivers_2_logo.svg?width=128"},
		{ID: mustObjectID("66b8a101000000000000001d"), Name: "Destiny 2", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Destiny_2_logo.svg?width=128"},
		{ID: mustObjectID("66b8a101000000000000001e"), Name: "Final Fantasy XIV", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Final_Fantasy_XIV_logo.svg?width=128"},
		{ID: mustObjectID("66b8a101000000000000001f"), Name: "World of Warcraft", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/World_of_Warcraft_Logo.svg?width=128"},
		{ID: mustObjectID("66b8a1010000000000000020"), Name: "Genshin Impact", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Genshin_Impact_logo.svg?width=128"},
		{ID: mustObjectID("66b8a1010000000000000021"), Name: "Honkai: Star Rail", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Honkai_Star_Rail_logo.svg?width=128"},
		{ID: mustObjectID("66b8a1010000000000000022"), Name: "Lost Ark", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Lost_Ark_logo.svg?width=128"},
		{ID: mustObjectID("66b8a1010000000000000023"), Name: "Black Myth: Wukong", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Black_Myth_Wukong_logo.svg?width=128"},
		{ID: mustObjectID("66b8a1010000000000000024"), Name: "Monster Hunter Wilds", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Monster_Hunter_Wilds_logo.svg?width=128"},
		{ID: mustObjectID("66b8a1010000000000000025"), Name: "Cyberpunk 2077", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Cyberpunk_2077_logo.svg?width=128"},
		{ID: mustObjectID("66b8a1010000000000000026"), Name: "Red Dead Redemption 2", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Red_Dead_Redemption_2_Logo.svg?width=128"},
		{ID: mustObjectID("66b8a1010000000000000027"), Name: "The Witcher 3: Wild Hunt", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/The_Witcher_3_Wild_Hunt_logo.svg?width=128"},
		{ID: mustObjectID("66b8a1010000000000000028"), Name: "Team Fortress 2", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Team_Fortress_2_logo.svg?width=128"},
		{ID: mustObjectID("66b8a1010000000000000029"), Name: "War Thunder", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/War_Thunder_logo.svg?width=128"},
		{ID: mustObjectID("66b8a101000000000000002a"), Name: "Terraria", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Terraria_Logo.svg?width=128"},
		{ID: mustObjectID("66b8a101000000000000002b"), Name: "ARK: Survival Ascended", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Ark_Survival_Ascended_logo.svg?width=128"},
		{ID: mustObjectID("66b8a101000000000000002c"), Name: "Lethal Company", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Lethal_Company_logo.svg?width=128"},
		{ID: mustObjectID("66b8a101000000000000002d"), Name: "Phasmophobia", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Phasmophobia_logo.svg?width=128"},
		{ID: mustObjectID("66b8a101000000000000002e"), Name: "The Finals", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/The_Finals_logo.svg?width=128"},
		{ID: mustObjectID("66b8a101000000000000002f"), Name: "Hogwarts Legacy", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Hogwarts_Legacy_logo.svg?width=128"},
		{ID: mustObjectID("66b8a1010000000000000030"), Name: "Forza Horizon 5", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Forza_Horizon_5_logo.svg?width=128"},
		{ID: mustObjectID("66b8a1010000000000000031"), Name: "Fall Guys", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Fall_Guys_logo.svg?width=128"},
		{ID: mustObjectID("66b8a1010000000000000032"), Name: "Among Us", IconUrl: "https://commons.wikimedia.org/wiki/Special:Redirect/file/Among_Us_logo.svg?width=128"},
	}
}
