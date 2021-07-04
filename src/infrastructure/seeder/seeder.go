package seeder

import (
	"context"
	uuid2 "github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
	"user-service/domain"
	"user-service/domain/enum"
)

func DropDatabase(db string, mongoCli *mongo.Client, ctx *context.Context){
	err := mongoCli.Database(db).Drop(*ctx)
	if err != nil {
		return
	}
}

func SeedData(db string, mongoCli *mongo.Client, ctx *context.Context) {
	DropDatabase(db, mongoCli, ctx)

	if cnt,_ := mongoCli.Database(db).Collection("profiles").EstimatedDocumentCount(*ctx, nil); cnt == 0 {
		personCollection := mongoCli.Database(db).Collection("profiles")
		seedProfileInfo(personCollection, ctx)
	}

	if cnt,_ := mongoCli.Database(db).Collection("report_users").EstimatedDocumentCount(*ctx, nil); cnt == 0 {
		personCollection := mongoCli.Database(db).Collection("report_users")
		seedReportUser(personCollection, ctx)
	}

	if cnt,_ := mongoCli.Database(db).Collection("request_verifications").EstimatedDocumentCount(*ctx, nil); cnt == 0 {
		personCollection := mongoCli.Database(db).Collection("request_verifications")
		seedRequestVerification(personCollection, ctx)
	}

	if cnt,_ := mongoCli.Database(db).Collection("admins").EstimatedDocumentCount(*ctx, nil); cnt == 0 {
		personCollection := mongoCli.Database(db).Collection("admins")
		seedAdmins(personCollection, ctx)
	}

}

func seedAdmins(tags *mongo.Collection, ctx *context.Context) {
	person1 := domain.Person{
		Name: "Mika",
		Surname: "Mikic",
		Gender: enum.Gender(0),
		DateOfBirth: time.Date( 1998, 06, 8, 20, 20, 20, 651387237, time.UTC),
		Address: "Novi Sad, Srbija",
		Phone: "011/2112-21111",
	}

	_, err := tags.InsertMany(*ctx, []interface{} {
		bson.D{{"_id", "111"},
			{"username", "admin1"},
			{"person", person1},
		},
	})

	if err != nil {
		log.Fatal(err)
	}

}

func seedProfileInfo(tags *mongo.Collection, ctx *context.Context) {
	person1 := domain.Person{
				Name: "Pera",
				Surname: "Peric",
				Gender: enum.Gender(0),
				DateOfBirth: time.Date( 1998, 06, 8, 20, 20, 20, 651387237, time.UTC),
				Address: "Novi Sad, Srbija",
				Phone: "011/2112-2111"}

	profile1 := domain.Profile{
				Username : "user1",
				PrivacyPermission : enum.PrivacyPermission(0),
				AllowTagging : true,
				AllowNotification: true,
				Type : enum.ProfileType(0)}

	person2 := domain.Person{
		Name: "Jovo",
		Surname: "Jovic",
		Gender: enum.Gender(0),
		DateOfBirth: time.Date( 1997, 06, 8, 20, 20, 20, 651387237, time.UTC),
		Address: "Sombor, Srbija",
		Phone: "013/2112-2111"}

	profile2 := domain.Profile{
		Username : "user2",
		PrivacyPermission : enum.PrivacyPermission(0),
		AllowTagging : true,
		AllowNotification: true,
		Type : enum.ProfileType(0)}

	person3 := domain.Person{
		Name: "Stevan",
		Surname: "Stevanovic",
		Gender: enum.Gender(0),
		DateOfBirth: time.Date( 1992, 06, 8, 20, 20, 20, 651387237, time.UTC),
		Address: "Velka plana, Srbija",
		Phone: "013/2112-2111"}

	profile3 := domain.Profile{
		Username : "user3",
		PrivacyPermission : enum.PrivacyPermission(0),
		AllowTagging : true,
		AllowNotification: true,
		Type : enum.ProfileType(0)}

	person4 := domain.Person{
		Name: "Zoka",
		Surname: "Bosanac",
		Gender: enum.Gender(0),
		DateOfBirth: time.Date( 1992, 06, 8, 20, 20, 20, 651387237, time.UTC),
		Address: "Novi Sad, Srbija",
		Phone: "021/281-221"}

	profile4 := domain.Profile{
		Username : "user4",
		PrivacyPermission : enum.PrivacyPermission(1),
		AllowTagging : true,
		AllowNotification: true,
		Type : enum.ProfileType(0)}

	person5 := domain.Person{
		Name: "Milica",
		Surname: "Milanovic",
		Gender: enum.Gender(1),
		DateOfBirth: time.Date( 1996, 06, 8, 20, 20, 20, 651387237, time.UTC),
		Address: "Novi Sad, Srbija",
		Phone: "021/281-221"}

	profile5 := domain.Profile{
		Username : "user5",
		PrivacyPermission : enum.PrivacyPermission(1),
		AllowTagging : true,
		AllowNotification: true,
		Type : enum.ProfileType(0)}

	person6 := domain.Person{
		Name: "Jovanka",
		Surname: "Stefanovic",
		Gender: enum.Gender(1),
		DateOfBirth: time.Date( 1992, 06, 8, 20, 20, 20, 651387237, time.UTC),
		Address: "Subotica, Srbija",
		Phone: "013/281-221"}

	profile6 := domain.Profile{
		Username : "user6",
		PrivacyPermission : enum.PrivacyPermission(1),
		AllowTagging : true,
		AllowNotification: true,
		Type : enum.ProfileType(0)}

	person7 := domain.Person{
		Name: "Admin",
		Surname: "Adminovic",
		Gender: enum.Gender(1),
		DateOfBirth: time.Date( 1992, 06, 8, 20, 20, 20, 651387237, time.UTC),
		Address: "Administratorovici, Srbija",
		Phone: "013/281-221"}

	profile7 := domain.Profile{
		Username : "admin1",
		PrivacyPermission : enum.PrivacyPermission(0),
		AllowTagging : true,
		AllowNotification: true,
		Type : enum.ProfileType(0)}

	person8 := domain.Person{
		Name: "Agent",
		Surname: "Agentovic",
		Gender: enum.Gender(0),
		DateOfBirth: time.Date( 1992, 06, 8, 20, 20, 20, 651387237, time.UTC),
		Address: "Agentoviicii, Srbija",
		Phone: "021/281-221"}

	profile8 := domain.Profile{
		Username : "agent1",
		PrivacyPermission : enum.PrivacyPermission(1),
		AllowTagging : true,
		AllowNotification: true,
		Type : enum.ProfileType(1)}

	_, err := tags.InsertMany(*ctx, []interface{} {
		bson.D{{"_id", "e2b5f92e-c31b-11eb-8529-0242ac130003"},
			{"email", "user1@gmail.com"},
			{"biography", "Ja sam kul osoba"},
			{"web_page", "pera.com"},
			{"category", enum.Category(0)},
			{"profile_image", "e2b5f92e-c31b-11eb-8529-0242ac130003/1b1a4d40-eb31-4acd-a947-ad398f47c692.jpeg"},
			{"person", person1},
			{"profile", profile1},
		},
		bson.D{{"_id", "424935b1-766c-4f99-b306-9263731518bc"},
			{"email", "user2@gmail.com"},
			{"biography", "Ja sam veoma kul kul kul idegasnamax osoba"},
			{"web_page", "sofascore.com"},
			{"category", enum.Category(6)},
			{"profile_image", "424935b1-766c-4f99-b306-9263731518bc/1babdbb1-dcd8-4325-b979-344900dff180.jpeg"},
			{"person", person2},
			{"profile", profile2},
		},
		bson.D{{"_id", "a2c2f993-dc32-4a82-82ed-a5f6866f7d03"},
			{"email", "user3@gmail.com"},
			{"biography", "IDEMO NIIIIIIIIIIIIIIIIIIIIIIIIIIS"},
			{"web_page", ""},
			{"category", enum.Category(6)},
			{"profile_image", "a2c2f993-dc32-4a82-82ed-a5f6866f7d03/pexels-photo-2078265.jpeg"},
			{"person", person3},
			{"profile", profile3},
		},
		bson.D{{"_id", "43420055-3174-4c2a-9823-a8f060d644c3"},
			{"email", "user4@gmail.com"},
			{"biography", "Biografija cetvrtog user jer mi je ponestalo ideja za biografiju"},
			{"web_page", ""},
			{"category", enum.Category(6)},
			{"profile_image", "43420055-3174-4c2a-9823-a8f060d644c3/ab67616d0000b2737f501335b4b23ce81214c753.jpg"},
			{"person", person4},
			{"profile", profile4},
		},
		bson.D{{"_id", "ead67925-e71c-43f4-8739-c3b823fe21bb"},
			{"email", "user5@gmail.com"},
			{"biography", "Idemo user 5 idemo gassssssss"},
			{"web_page", ""},
			{"category", enum.Category(6)},
			{"profile_image", "ead67925-e71c-43f4-8739-c3b823fe21bb/cool-profile-pictures-retouching-1.jpg"},
			{"person", person5},
			{"profile", profile5},
		},
		bson.D{{"_id", "23ddb1dd-4303-428b-b506-ff313071d5d7"},
			{"email", "user6@gmail.com"},
			{"biography", "Get in there Lewis!  NEK HAMILTON JEDE PRASINUUUU"},
			{"web_page", "www.formula1.com/"},
			{"category", enum.Category(6)},
			{"profile_image", "23ddb1dd-4303-428b-b506-ff313071d5d7/Zoom-How-to-Set-Profile-Picture.jpg"},
			{"person", person6},
			{"profile", profile6},
		},
		bson.D{{"_id", "bdb7d7c5-2c9a-4b4c-ab64-4e4828d93926"},
			{"email", "admin1@gmail.com"},
			{"biography", "Admin 1 bios"},
			{"web_page", "www.formula1.com/"},
			{"category", enum.Category(6)},
			{"profile_image", ""},
			{"person", person7},
			{"profile", profile7},
		},
		bson.D{{"_id", "1d09bb0a-d9fc-11eb-b8bc-0242ac130003"},
			{"email", "agent1@gmail.com"},
			{"biography", "Agent 1 bios"},
			{"web_page", "www.google.com/"},
			{"category", enum.Category(6)},
			{"profile_image", ""},
			{"person", person8},
			{"profile", profile8},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}

func seedReportUser(tags *mongo.Collection, ctx *context.Context) {
	_, err := tags.InsertMany(*ctx, []interface{} {
		bson.D{{"_id", uuid2.New().String()},
			{"reported", "pera123"},
			{"report_type", enum.UserReportType(0)},
			{"timestamp", time.Now()},
			{"report_status", enum.ReportStatus(1)},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}

func seedRequestVerification(tags *mongo.Collection, ctx *context.Context) {
	_, err := tags.InsertMany(*ctx, []interface{} {
		bson.D{{"_id", uuid2.New().String()},
			{"name", "Pera"},
			{"surname", "Peric"},
			{"category", enum.Category(0)},
			{"image", ""},
			{"state", enum.VerificationState(0)},
			{"profile_id", "e2b5f92e-c31b-11eb-8529-0242ac130003"},

		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
