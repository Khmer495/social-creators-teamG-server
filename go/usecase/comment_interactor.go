package usecase

import "github.com/Khmer495/social-creators-teamG-server/go/entity/model"

type CommentInteractor struct {
	CommentRepository CommentRepository
}

func (interactor *CommentInteractor) CommentById(id int) (comment model.Comment, err error) {
	comment, err = interactor.CommentRepository.FindById(id)
	return
}

func (interactor *CommentInteractor) Comments() (comments model.Comments, err error) {
	comments, err = interactor.CommentRepository.FindAll()
	return
}

func (interactor *CommentInteractor) Add(c model.Comment) (comment model.Comment, err error) {
	comment, err = interactor.CommentRepository.Store(c)
	return
}

func (interactor *CommentInteractor) Update(c model.Comment) (comment model.Comment, err error) {
	comment, err = interactor.CommentRepository.Update(c)
	return
}

func (interactor *CommentInteractor) DeleteById(c model.Comment) (err error) {
	err = interactor.CommentRepository.DeleteById(c)
	return
}
