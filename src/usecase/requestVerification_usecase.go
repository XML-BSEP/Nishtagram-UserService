package usecase

import (
	"bufio"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"os"
	"strings"
	"user-service/dto"
	"user-service/infrastructure/mapper"
	"user-service/repository"
)

type requestVerificationUseCase struct {
	RequestVerificationRepository repository.RequestVerificationRepository
	ProfileInfoRepository repository.ProfileInfoRepository
}


type RequestVerificationUseCase interface {
	SaveNewRequestVerification(verification dto.RequestVerificationDTO, ctx context.Context) (string, error)
	GetAllRequestVerificationForWaiting(ctx context.Context) *[]dto.RequestVerificationDTO
	GetAllRequestVerification(ctx context.Context) *[]dto.RequestVerificationDTO
	ApproveRequestVerification(verificationId string, profileId string, ctx context.Context) (bool, error)
	RejectRequestVerification(verificationId string, ctx context.Context) (bool, error)
	EncodeBase64(media string, verificationId string, ctx context.Context) (string, error)
	DecodeBase64(media string, verificationId string, ctx context.Context) (string, error)
	ExistsRequestForProfile(profileId string, ctx context.Context) (bool,error)

}

func (r *requestVerificationUseCase) ExistsRequestForProfile(profileId string, ctx context.Context) (bool, error) {
	exists, err := r.RequestVerificationRepository.ExistsRequestForProfile(profileId, ctx)
	return exists, err

}


func (r *requestVerificationUseCase) SaveNewRequestVerification(verification dto.RequestVerificationDTO, ctx context.Context) (string, error) {

	newVerificationId := uuid.New().String()

	if verification.Image != "" {
		encodedImage, err := r.EncodeBase64(verification.Image, newVerificationId, ctx)
		if err != nil {
			return "Encoded image error", err
		}
		verification.Image = encodedImage
	}

	newVerificationRequest := mapper.NewVerificationRequestDTOtoEntity(verification)
	newVerificationRequest.ID = newVerificationId

	done, err := r.RequestVerificationRepository.SaveNewRequestVerification(newVerificationRequest, ctx)
	if err != nil {
		return "Failed to save", err
	}

	if done {
		return "Inserted new verification", nil
	}

	return "Failed to save", err

}

func (r *requestVerificationUseCase) GetAllRequestVerificationForWaiting(ctx context.Context) *[]dto.RequestVerificationDTO {
	verifications := r.RequestVerificationRepository.GetAllRequestVerificationForWaiting(ctx)

	var verificationsDTO []dto.RequestVerificationDTO
	for _, verification := range *verifications {
		if verification.Image != "" {
			decodedImage, err := r.DecodeBase64(verification.Image, verification.ID, ctx)
			if err != nil {
				return nil
			}
			verification.Image = decodedImage
		}
		verificationsDTO = append(verificationsDTO, mapper.RequestVerificationToDTO(verification))
	}

	return &verificationsDTO
}

func (r *requestVerificationUseCase) GetAllRequestVerification(ctx context.Context) *[]dto.RequestVerificationDTO {

	verifications := r.RequestVerificationRepository.GetAllRequestVerification(ctx)

	var verificationsDTO []dto.RequestVerificationDTO
	for _, verification := range *verifications {
		if verification.Image != "" {
			decodedImage, err := r.DecodeBase64(verification.Image, verification.ID, ctx)
			if err != nil {
				return nil
			}
			verification.Image = decodedImage
		}

		verificationsDTO = append(verificationsDTO, mapper.RequestVerificationToDTO(verification))

	}

	return &verificationsDTO

}

func (r *requestVerificationUseCase) ApproveRequestVerification(verificationId string, profileId string, ctx context.Context) (bool, error) {

	approved, err := r.RequestVerificationRepository.ApproveRequestVerification(verificationId, ctx)

	if approved {
		req, err := r.RequestVerificationRepository.GetByProfileId(profileId, ctx)
		if err != nil {
			return false, err
		}

		error := r.ProfileInfoRepository.ChangeProfileCategory(profileId, req.Category, ctx)

		if error != nil {
			return false, err
		}
	}


	return approved, err
}

func (r *requestVerificationUseCase) RejectRequestVerification(verificationId string,ctx context.Context) (bool, error) {
	rejected, err := r.RequestVerificationRepository.RejectRequestVerification(verificationId, ctx)

	return rejected, err
}


func (r *requestVerificationUseCase) EncodeBase64(media string, verificationId string, ctx context.Context) (string, error) {
	workingDirectory, _ := os.Getwd()
	if !strings.HasSuffix(workingDirectory, "src") {
		firstPart := strings.Split(workingDirectory, "src")
		value := firstPart[0] + "src"
		workingDirectory = value
		os.Chdir(workingDirectory)
	}

	path1 := "./assets/verifications"
	err := os.Chdir(path1)
	if err != nil {
		return "", err
	}
	err = os.Mkdir(verificationId, 0755)
	fmt.Println(err)

	err = os.Chdir(verificationId)
	fmt.Println(err)

	s := strings.Split(media, ",")
	a := strings.Split(s[0], "/")
	format := strings.Split(a[1], ";")
	dec, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		return "", err
	}
	uuid := uuid.NewString()
	f, err := os.Create(uuid + "." + format[0])
	if err != nil {
		return "", err
	}

	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		panic(err)
	}
	if err := f.Sync(); err != nil {
		panic(err)
	}

	os.Chdir(workingDirectory)
	return verificationId + "/" + uuid + "." + format[0], nil
}

func (r *requestVerificationUseCase) DecodeBase64(media string, verificationId string, ctx context.Context) (string, error) {


	workingDirectory, _ := os.Getwd()
	if !strings.HasSuffix(workingDirectory, "src") {
		firstPart := strings.Split(workingDirectory, "src")
		value := firstPart[0] + "src"
		workingDirectory = value
		os.Chdir(workingDirectory)
	}

	path1 := "./assets/verifications"
	err := os.Chdir(path1)
	fmt.Println(err)

	err = os.Chdir(verificationId)
	fmt.Println(err)

	spliced := strings.Split(media, "/")
	var f *os.File
	if len(spliced) > 1 {
		f, _ = os.Open(spliced[1])
	} else {
		f, _ = os.Open(spliced[0])
	}


	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)


	encoded := base64.StdEncoding.EncodeToString(content)

	os.Chdir(workingDirectory)

	return "data:image/jpg;base64," + encoded, nil
}


func NewRequestVerificationUseCase(repo repository.RequestVerificationRepository, infoRepository repository.ProfileInfoRepository) RequestVerificationUseCase {
	return &requestVerificationUseCase{RequestVerificationRepository: repo, ProfileInfoRepository: infoRepository}
}