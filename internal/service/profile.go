package service

import (
	"online_lib_api"
	"online_lib_api/internal/storage"
)

type ProfileService struct {
	storage storage.Profile
}

func NewProfileService(storage storage.Profile) *ProfileService {
	return &ProfileService{storage: storage}
}

func (s *ProfileService) GetProfileData(user_id int) (online_lib_api.Profile, error) {
	profile_data, err := s.storage.GetProfileData(user_id)
	return profile_data, err
}

func (s *ProfileService) AddBook(user_id int, book_id int) error {
	err := s.storage.AddBook(user_id, book_id)
	return err
}

func (s *ProfileService) DeleteBook(user_id int, book_id int) error {
	err := s.storage.DeleteBook(user_id, book_id)
	return err
}