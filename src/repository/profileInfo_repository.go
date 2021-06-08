package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
	"time"
	"user-service/domain"
	"user-service/domain/enum"
)

type profileInfoRepository struct {
	collection *mongo.Collection
	db *mongo.Client
}

type ProfileInfoRepository interface {
	GetByUsername(username string, ctx context.Context) (domain.ProfileInfo, error)
	GetAllProfiles(ctx context.Context) ([]domain.ProfileInfo, error)
	GetAllUserProfiles(ctx context.Context) ([]domain.ProfileInfo, error)
	GetById(id string, ctx context.Context) (*domain.ProfileInfo, error)
	GetUserById(id string, ctx context.Context) (domain.ProfileInfo, error)
	//GetUserProfileById(id string, ctx context.Context) (dto.UserProfileDTO, error)
	SaveNewUser(user domain.ProfileInfo, ctx context.Context) error
	IsProfilePrivate(username string, ctx context.Context) (bool, error)
	Exists(username string, email string, ctx context.Context) (bool, error)
	GetAllPublicProfiles(ctx context.Context) ([]domain.ProfileInfo, error)
	EditUser(user domain.ProfileInfo, ctx context.Context) error
	SearchUser(search string, ctx context.Context) ([]*domain.ProfileInfo, error)
	IsPrivateById(id string, ctx context.Context) (bool, error)

}

func (p *profileInfoRepository) IsPrivateById(id string, ctx context.Context) (bool, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()


	var profile domain.ProfileInfo
	err := p.collection.FindOne(ctx, bson.M{"_id" : id}).Decode(&profile)
	if err != nil {
		return false, err
	}

	if profile.Profile.PrivacyPermission.String() == "Private" {
		return true, nil
	}

	return false, nil

}


func (p *profileInfoRepository) Exists(username string, email string, ctx context.Context) (bool, error){
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()


	var profile domain.ProfileInfo
	err := p.collection.FindOne(ctx, bson.M{"profile.username" : username, "email" : email}).Decode(&profile)
	if err != nil {
		return false, err
	}
	return true, nil

}

func (p *profileInfoRepository) IsProfilePrivate(username string, ctx context.Context) (bool, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()


	var profile domain.ProfileInfo
	err := p.collection.FindOne(ctx, bson.M{"profile.username" : username}).Decode(&profile)
	if err != nil {
		return false, err
	}

	if profile.Profile.PrivacyPermission.String() == "Private" {
		return true, nil
	}

	return false, nil

}

func (p *profileInfoRepository) GetByUsername(username string, ctx context.Context) (domain.ProfileInfo, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var profile domain.ProfileInfo
	err := p.collection.FindOne(ctx, bson.M{"profile.username" : username}).Decode(&profile)
	if err != nil {
		return profile, err
	}
	return profile, nil
}

func (p *profileInfoRepository) GetAllProfiles(ctx context.Context) ([]domain.ProfileInfo, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	profiles , err := p.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var allProfiles []domain.ProfileInfo
	if err = profiles.All(ctx, &allProfiles); err != nil {
		return nil, err
	}

	return allProfiles, nil
}

func (p *profileInfoRepository) GetAllUserProfiles(ctx context.Context) ([]domain.ProfileInfo, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	profiles , err := p.collection.Find(ctx, bson.M{"profile.type" : enum.ProfileType(0).String()})
	if err != nil {
		return nil, err
	}

	var allProfiles []domain.ProfileInfo
	if err = profiles.All(ctx, &allProfiles); err != nil {
		return nil, err
	}

	return allProfiles, nil
}

func (p *profileInfoRepository) GetById(id string, ctx context.Context) (*domain.ProfileInfo, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var profile *domain.ProfileInfo
	err := p.collection.FindOne(ctx, bson.M{"_id" : id}).Decode(&profile)
	if err != nil {
		return profile, err
	}

	return profile, nil
}

func (p *profileInfoRepository) GetUserById(id string, ctx context.Context) (domain.ProfileInfo, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var profile domain.ProfileInfo

	err := p.collection.FindOne(ctx, bson.M{"_id" : id}).Decode(&profile)
	if err != nil {
		return profile, err
	}
	/*
	userDTO = dto.NewSimplyUserDTO(profile.Person.Name, profile.Person.Surname, profile.Email, profile.Person.Address,
		profile.Person.Phone, profile.Person.DateOfBirth.Format("02-Jan-2006"), profile.Person.Gender, profile.WebPage, profile.Biography,
		profile.Profile.Username, profile.ProfileImage)*/


	return profile, nil

}

/*func (p *profileInfoRepository) GetUserProfileById(id string, ctx context.Context) (dto.UserProfileDTO, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var profile domain.ProfileInfo
	err := p.collection.FindOne(ctx, bson.M{"_id" : id}).Decode(&profile)
	if err != nil {
		return dto.UserProfileDTO{}, err
	}

	userDTO, err1 := p.GetUserById(id, ctx)
	if err1 != nil {
		return dto.UserProfileDTO{}, err
	}

	var isPrivate bool
	if profile.Profile.PrivacyPermission == 0 {
		isPrivate = true
	}

	if profile.Profile.PrivacyPermission == 1 {
		isPrivate = false
	}

	userProfileDTO := dto.NewUserProfileDTO(dto.NewUserDTOfromEntity(userDTO), &isPrivate)

	return userProfileDTO, nil

}*/

func (p *profileInfoRepository) SaveNewUser(user domain.ProfileInfo, ctx context.Context) error {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()


	_, err := p.collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (p *profileInfoRepository) GetAllPublicProfiles(ctx context.Context) ([]domain.ProfileInfo, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	profiles , err := p.collection.Find(ctx, bson.M{"profile.privacy_permission" : 1})
	if err != nil {
		return nil, err
	}

	var allProfiles []domain.ProfileInfo
	if err = profiles.All(ctx, &allProfiles); err != nil {
		return nil, err
	}

	return allProfiles, nil
}


func (p *profileInfoRepository) EditUser(user domain.ProfileInfo, ctx context.Context) error {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	/*Email string `bson:"email" json:"email"`
	Biography string `bson:"biography" json:"biography"`
	WebPage string `bson:"web_page" json:"web_page"`
	Category enum.Category `json:"category" bson:"category"`
	ProfileImage string `bson:"profile_image" json:"profile_image"`
	Person Person `bson:"person" json:"person"`
	Profile Profile `bson:"profile" json:"profile"`*/


	userToUpdate := bson.M{"_id" : user.ID}
	updatedUser := bson.M{"$set": bson.M{
		"email":      user.Email,
		"biography":    user.Biography,
		"web_page":       user.WebPage,
		"category": user.Category,
		"profile_image" : user.ProfileImage,
		"person" : user.Person,
		"profile" : user.Profile,

	}}


	_, err := p.collection.UpdateOne(ctx, userToUpdate, updatedUser)
	if err != nil {
		return  err
	}

	return nil

}

func (p *profileInfoRepository) SearchUser(search string, ctx context.Context) ([]*domain.ProfileInfo, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var foundUsers []*domain.ProfileInfo

	search = strings.TrimSpace(search)
	splitSearch := strings.Split(search, " ")
	for _, splitSearchpart := range splitSearch {

		//username
		filtereds, err := p.collection.Find(ctx, bson.M{"profile.username" : primitive.Regex{Pattern: splitSearchpart, Options: "i"} })
		if err != nil {
			return nil, err
		}
		var usersUsername []*domain.ProfileInfo
		if err = filtereds.All(ctx, &usersUsername); err != nil {
			return nil, err
		}
		for _, userOneSlice := range usersUsername {
			foundUsers = AppendIfMissing(foundUsers, userOneSlice)
		}

		//name
		filtereds, err = p.collection.Find(ctx, bson.M{"person.name" : primitive.Regex{Pattern: splitSearchpart, Options: "i"} })
		if err != nil {
			return nil, err
		}
		var usersName []*domain.ProfileInfo
		if err = filtereds.All(ctx, &usersName); err != nil {
			return nil, err
		}
		for _, userOneSlice := range usersName {
			foundUsers = AppendIfMissing(foundUsers, userOneSlice)
		}

		//surname
		filtereds, err = p.collection.Find(ctx, bson.M{"person.surname" : primitive.Regex{Pattern: splitSearchpart, Options: "i"} })
		if err != nil {
			return nil, err
		}
		var usersSurname []*domain.ProfileInfo
		if err = filtereds.All(ctx, &usersSurname); err != nil {
			return nil, err
		}
		for _, userOneSlice := range usersSurname {
			foundUsers = AppendIfMissing(foundUsers, userOneSlice)
		}


	}



	return foundUsers, nil

}


func NewProfileInfoRepository(db *mongo.Client) ProfileInfoRepository {
	return &profileInfoRepository {
		db : db,
		collection: db.Database("user_db").Collection("profiles"),
	}
}


func AppendIfMissing(slice []*domain.ProfileInfo, i *domain.ProfileInfo) []*domain.ProfileInfo {
	for _, ele := range slice {
		if ele.ID == i.ID {
			return slice
		}
	}
	return append(slice, i)
}
