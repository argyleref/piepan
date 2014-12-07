package piepan

import (
	"github.com/layeh/gumble/gumble"
)

type disconnectEventWrapper struct {
	Client *gumble.Client
	Type   int

	String string

	IsError bool
	IsUser  bool

	IsOther             bool
	IsVersion           bool
	IsUserName          bool
	IsUserCredentials   bool
	IsServerPassword    bool
	IsUsernameInUse     bool
	IsServerFull        bool
	IsNoCertificate     bool
	IsAuthenticatorFail bool
}

type userChangeEventWrapper struct {
	Client *gumble.Client
	Type   int
	User   *gumble.User
	Actor  *gumble.User

	String string

	IsConnected     bool
	IsDisconnected  bool
	IsKicked        bool
	IsBanned        bool
	IsChangeName    bool
	IsChangeChannel bool
	IsChangeComment bool
}

type channelChangeEventWrapper struct {
	Client  *gumble.Client
	Type    int
	Channel *gumble.Channel

	IsCreated           bool
	IsRemoved           bool
	IsMoved             bool
	IsChangeName        bool
	IsChangeDescription bool
}

type permissionDeniedEventWrapper struct {
	Client  *gumble.Client
	Type    int
	Channel *gumble.Channel
	User    *gumble.User

	Permission int
	String     string

	IsOther              bool
	IsPermission         bool
	IsSuperUser          bool
	IsInvalidChannelName bool
	IsTextTooLong        bool
	IsTemporaryChannel   bool
	IsMissingCertificate bool
	IsInvalidUserName    bool
	IsChannelFull        bool
	IsNestingLimit       bool
}
