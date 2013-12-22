package controllers

const (
	_ = iota
	AdminGroup
	MemberGroup
)

var Permissions = map[string]int{
	// Admin
	"Admin.Index":      AdminGroup,
	"Admin.ListUser":   AdminGroup,
	"Admin.DeleteUser": AdminGroup,
	"Admin.ListBook":   AdminGroup,
	"Admin.DeleteBook": AdminGroup,
	"Admin.NewBook":    AdminGroup,
	"Admin.EditBook":   AdminGroup,

	// User
	"User.Edit":       MemberGroup,
	"User.Borrow":     MemberGroup,
	"User.Owned":      MemberGroup,
	"User.BorrowHis":  MemberGroup,
	"User.BookDel":    MemberGroup,
	"User.BookReturn": MemberGroup,

	// Book
	"Book.Collect":   MemberGroup,
	"Book.UnCollect": MemberGroup,
	"Book.Booking":   MemberGroup,
	"Book.UnBooking": MemberGroup,

	// Topic
	"Topic.New":   MemberGroup,
	"Topic.Edit":  MemberGroup,
	"Topic.Reply": MemberGroup,
}
