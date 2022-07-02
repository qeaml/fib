package gallery

import "time"

type UserFlag uint8

const (
	UserFlagNone  UserFlag = 0
	UserFlagAdmin UserFlag = 1 << (iota - 1)
	UserFlagModerator
	UserFlagBanned
	UserFlagLocked
)

type User struct {
	ID         string
	Flags      UserFlag
	Name       string
	Bio        string
	Avatar     uint32
	Registered time.Time
	LastLogin  time.Time
}

func (u *User) Rank() string {
	if u.Flags&UserFlagAdmin != 0 {
		return "Administrator"
	}
	if u.Flags&UserFlagModerator != 0 {
		return "Moderator"
	}
	if u.Flags&UserFlagBanned != 0 {
		return "Banned"
	}
	return "User"
}

type ImageFlag uint8

const (
	ImageFlagNone   ImageFlag = 0
	ImageFlagHidden ImageFlag = 1 << (iota - 1)
	ImageFlagPrivate
	ImageFlagNSFW
)

type Image struct {
	ID         uint32
	Flags      ImageFlag
	Title      string
	Desc       string
	Tags       []string
	Album      string
	Uploader   string
	UploadedAt time.Time
	UpdatedAt  time.Time
}

func (i *Image) Visibility() string {
	if i.Flags&ImageFlagHidden != 0 {
		return "Hidden"
	}
	if i.Flags&ImageFlagPrivate != 0 {
		return "Private"
	}
	if i.Flags&ImageFlagNSFW != 0 {
		return "NSFW"
	}
	return "Public"
}

// Public returns true if the image is not private or hidden
func (i *Image) Public() bool {
	return i.Flags&(ImageFlagPrivate|ImageFlagHidden) == 0
}

func (i *Image) NSFW() bool {
	return i.Flags&ImageFlagNSFW != 0
}

type Album struct {
	ID         string
	Flags      ImageFlag
	Title      string
	Desc       string
	Tags       []string
	Uploader   string
	UploadedAt time.Time
	UpdatedAt  time.Time
}

func (a *Album) Visibility() string {
	if a.Flags&ImageFlagHidden != 0 {
		return "Hidden"
	}
	if a.Flags&ImageFlagPrivate != 0 {
		return "Private"
	}
	if a.Flags&ImageFlagNSFW != 0 {
		return "NSFW"
	}
	return "Public"
}
