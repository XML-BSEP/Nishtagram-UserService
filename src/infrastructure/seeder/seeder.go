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

	if cnt,_ := mongoCli.Database(db).Collection("persons").EstimatedDocumentCount(*ctx, nil); cnt == 0 {
		personCollection := mongoCli.Database(db).Collection("persons")
		seedPerson(personCollection, ctx)
	}

	if cnt,_ := mongoCli.Database(db).Collection("profiles").EstimatedDocumentCount(*ctx, nil); cnt == 0 {
		personCollection := mongoCli.Database(db).Collection("profiles")
		seedProfile(personCollection, ctx)
	}

	if cnt,_ := mongoCli.Database(db).Collection("profiles_info").EstimatedDocumentCount(*ctx, nil); cnt == 0 {
		personCollection := mongoCli.Database(db).Collection("profiles_info")
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

}



func seedPerson(tags *mongo.Collection, ctx *context.Context) {
	_, err := tags.InsertMany(*ctx, []interface{} {
		bson.D{{"_id",  "123456789"},
		{"name", "Pera"},
		{"surname", "Peric"},
		{"gender", enum.Gender(0)},
		{"date_of_birth", time.Date( 1998, 06, 8, 20, 20, 20, 651387237, time.UTC)},
		{"address","Bulevar Patrijarha Pavla 10, Novi Sad, Srbija"},
		{"phone", "011/2112-2111"},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}

func seedProfile(tags *mongo.Collection, ctx *context.Context) {
	_, err := tags.InsertMany(*ctx, []interface{} {
		bson.D{{"username", "pera1234"},
			{"privacy_permission", enum.PrivacyPermission(0)},
			{"allow_tagging", true},
			{"allow_notification", true},
			{"type", enum.ProfileType(0)},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}

func seedProfileInfo(tags *mongo.Collection, ctx *context.Context) {
	person1 := domain.Person{ID: "123456789",
	Name: "Pera",
	Surname: "Peric",
	Gender: enum.Gender(0),
	DateOfBirth: time.Date( 1998, 06, 8, 20, 20, 20, 651387237, time.UTC),
	Address: "Bulevar Patrijarha Pavla 10, Novi Sad, Srbija",
	Phone: "011/2112-2111"}

	profile1 := domain.Profile{
		Username : "pera1234",
		PrivacyPermission : enum.PrivacyPermission(1),
		AllowTagging : true,
		AllowNotification: true,
		Type : enum.ProfileType(0),
	}

	_, err := tags.InsertMany(*ctx, []interface{} {
		bson.D{{"email", "para@pera.com"},
			{"password", "$2y$12$haErHk8CNLaodc4rfEFwNeX5MRkXlfpeziRCimK025BJhFVlQmpIa"},
			{"biography", "Ja sam kul osoba"},
			{"web_page", "pera.com"},
			{"category", enum.Category(6)},
			{"profile_image", ""},
			{"person", person1},
			{"profile", profile1},
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
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
