package db

import (
	"errors"
	"fmt"
	"time"

	"github.com/musgit-dev/musgit/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Composer struct {
	gorm.Model
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"primary_key"`
}

type Piece struct {
	gorm.Model
	ID              uint   `gorm:"primary_key"`
	Name            string `gorm:"uniqueIndex:piece_composer"`
	ComposerID      int64  `gorm:"uniqueIndex:piece_composer"`
	Composer        Composer
	PieceComplexity models.PieceComplexity
	State           models.PieceState
	Practices       []Practice
	Users           []*User `gorm:"many2many:user_practices;"`
}

type Lesson struct {
	gorm.Model
	ID        uint `gorm:"primary_key"`
	State     models.LessonState
	StartDate time.Time
	EndDate   time.Time
	Comment   string
	UserID    uint
}

type Practice struct {
	gorm.Model
	ID        uint `gorm:"primary_key"`
	StartDate time.Time
	EndDate   time.Time
	Progress  models.PracticeProgressEvalutation
	PieceID   uint
	Piece     Piece
	LessonID  uint
	Lesson    Lesson
	UserID    uint
}

type Warmup struct {
	gorm.Model
	ID        uint `gorm:"primary_key"`
	StartDate time.Time
	EndDate   time.Time
	State     models.WarmupState
	LessonID  uint
	Lesson    Lesson
	UserID    uint
}

type User struct {
	gorm.Model
	ID        uint       `gorm:"primary_key"`
	Name      string     `gorm:"primary_key"`
	Pieces    []Piece    `gorm:"many2many:user_pieces;"`
	Practices []Practice `gorm:"many2many:user_practices;"`
	Warmups   []Warmup   `gorm:"many2many:user_warmups;"`
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dbUrl string) (*Adapter, error) {
	db, openErr := gorm.Open(sqlite.Open(dbUrl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if openErr != nil {
		return nil, fmt.Errorf("Db connection error: %v", openErr)
	}
	if err := db.AutoMigrate(
		&Composer{}, &Piece{}, &Practice{}, &Lesson{}, &Warmup{}, &User{},
	); err != nil {
		return nil, fmt.Errorf("Db migration error: %v", err)
	}

	return &Adapter{db: db}, nil
}

func (a *Adapter) checkPiece(name string, composerId int64) bool {
	var p Piece
	res := a.db.Where("name = ? AND composer_id = ?", name, composerId).
		First(&p)
	return res.RowsAffected != 0
}

func (a *Adapter) checkComposer(name string) int64 {
	var c Composer
	_ = a.db.Where("name = ?", name).First(&c)
	return int64(c.ID)
}

func (a *Adapter) GetPiece(id int64) (models.Piece, error) {

	var p Piece

	res := a.db.Joins("Composer").First(&p, id)
	var practices []*models.Practice

	for _, v := range p.Practices {
		practices = append(practices, &models.Practice{
			StartDate: v.StartDate,
			EndDate:   v.EndDate,
			Progress:  v.Progress,
		})
	}
	piece := models.Piece{
		ID:         int64(p.ID),
		Composer:   models.Composer{Name: p.Composer.Name},
		Name:       p.Name,
		State:      p.State,
		Complexity: p.PieceComplexity,
		Practices:  practices,
	}
	return piece, res.Error
}

func (a *Adapter) AddLesson(l *models.Lesson) (*models.Lesson, error) {
	lessonModel := Lesson{
		State:     l.State,
		StartDate: l.StartDate,
		EndDate:   l.EndDate,
		Comment:   l.Comment,
	}
	res := a.db.Create(&lessonModel)
	if res.Error == nil {
		l.ID = int64(lessonModel.ID)
	}
	return l, res.Error
}

func (a *Adapter) GetLesson(id int64) (models.Lesson, error) {

	var l Lesson

	res := a.db.First(&l, id)

	lesson := models.Lesson{
		ID:        int64(l.ID),
		State:     l.State,
		StartDate: l.StartDate,
		EndDate:   l.EndDate,
		Comment:   l.Comment,
	}
	return lesson, res.Error
}

func (a *Adapter) GetLastLesson() (models.Lesson, error) {

	var l Lesson

	res := a.db.Last(&l)

	lesson := models.Lesson{
		ID:        int64(l.ID),
		State:     l.State,
		StartDate: l.StartDate,
		EndDate:   l.EndDate,
		Comment:   l.Comment,
	}
	return lesson, res.Error
}

func (a *Adapter) GetLessons() []models.Lesson {
	var lessons []models.Lesson
	a.db.Find(&lessons)
	return lessons
}

func (a *Adapter) UpdateLesson(l *models.Lesson) error {
	res := a.db.Save(l)
	return res.Error
}

func (a *Adapter) AddPiece(piece *models.Piece) (*models.Piece, error) {

	composerId := a.checkComposer(piece.Composer.Name)
	if a.checkPiece(piece.Name, composerId) {
		return &models.Piece{}, errors.New("Already exists")
	}

	var practices []Practice

	for _, v := range piece.Practices {
		practices = append(practices, Practice{
			StartDate: v.StartDate,
			EndDate:   v.EndDate,
			Progress:  v.Progress,
		})
	}
	pieceModel := Piece{
		Name:            piece.Name,
		State:           piece.State,
		PieceComplexity: piece.Complexity,
		Practices:       practices,
	}

	if composerId != 0 {
		pieceModel.ComposerID = composerId
	} else {
		pieceModel.Composer = Composer{Name: piece.Composer.Name}
	}

	res := a.db.Create(&pieceModel)
	if res.Error == nil {
		piece.ID = int64(pieceModel.ID)
	}
	return piece, res.Error
}

func (a *Adapter) GetPieces() []models.Piece {
	var ps []Piece
	var pieces []models.Piece
	a.db.Joins("Composer").Find(&ps)
	for _, p := range ps {
		piece := models.Piece{
			ID:         int64(p.ID),
			Composer:   models.Composer{Name: p.Composer.Name},
			Name:       p.Name,
			State:      p.State,
			Complexity: p.PieceComplexity,
		}
		pieces = append(pieces, piece)
	}
	return pieces
}

func (a *Adapter) UpdatePiece(p *models.Piece) error {
	res := a.db.Save(p)
	return res.Error
}

func (a *Adapter) AddPractice(
	practice *models.Practice,
) (*models.Practice, error) {

	practiceModel := Practice{
		StartDate: practice.StartDate,
		PieceID:   uint(practice.PieceID),
		LessonID:  uint(practice.LessonID),
	}

	res := a.db.Create(&practiceModel)
	if res.Error == nil {
		practice.ID = int64(practiceModel.ID)
	}
	return practice, res.Error
}

func (a *Adapter) AddWarmup(
	w *models.Warmup,
) error {

	warmupModel := Warmup{
		StartDate: w.StartDate,
		LessonID:  uint(w.LessonID),
	}

	res := a.db.Create(&warmupModel)
	if res.Error == nil {
		w.ID = int64(warmupModel.ID)
	}
	return res.Error
}

func (a *Adapter) UpdateWarmup(warmup *models.Warmup) error {
	res := a.db.Save(warmup)
	return res.Error
}

func (a *Adapter) GetActiveWarmup() (*models.Warmup, error) {
	var w Warmup

	res := a.db.Last(&w).Where("state = 0")
	if res.Error != nil {
		return &models.Warmup{}, res.Error
	}
	warmup := models.Warmup{
		ID:        int64(w.ID),
		StartDate: w.StartDate,
	}
	return &warmup, res.Error
}

func (a *Adapter) GetPractice(id int64) (models.Practice, error) {

	var p Practice

	res := a.db.First(&p, id)

	lesson := models.Practice{
		ID:        int64(p.ID),
		StartDate: p.StartDate,
		EndDate:   p.EndDate,
		Progress:  p.Progress,
		LessonID:  int64(p.LessonID),
	}
	return lesson, res.Error
}

func (a *Adapter) GetPractices() []models.Practice {
	var practices []models.Practice
	a.db.Find(&practices)
	return practices
}

func (a *Adapter) UpdatePractice(practice *models.Practice) error {
	res := a.db.Save(practice)
	return res.Error
}

func (a *Adapter) AddUser(user *models.User) (*models.User, error) {
	userModel := User{
		Name: user.Name,
	}
	res := a.db.Create(&userModel)
	if res.Error == nil {
		user.ID = int64(userModel.ID)
	}
	return user, res.Error
}
func (a *Adapter) GetUser(id int64) (*models.User, error) {
	var u User
	res := a.db.First(&u, id)
	user := models.User{
		ID:   int64(u.ID),
		Name: u.Name,
	}
	return &user, res.Error
}

func (a *Adapter) GetUsers() ([]*models.User, error) {
	var urs []User
	var users []*models.User
	a.db.Find(&urs)
	for _, u := range urs {
		user := models.User{
			ID:   int64(u.ID),
			Name: u.Name,
		}
		users = append(users, &user)
	}
	return users, nil
}
