package planning

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"strings"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Id          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	StartUnix   int64
	EndUnix     int64
}

type Schedule struct {
	gorm.Model
	Id        string
	Scheduled []*Item
}

type Sys struct {
	Filename string
	Db       gorm.DB
}

func NewSys(filename string) Sys {
	s := Sys{Filename: filename}
	db, err := gorm.Open(sqlite.Open(filename), &gorm.Config{})
	if err != nil {
		panic("DB was not made")
	}
	s.Db = db
	return s
}

func (s *Sys) AddItem(title string) {
	i := NewItem()
	i.Title = title

}

func (i Item) GoString() string {
	var str strings.Builder
	str.WriteString(fmt.Sprintf("ID: %s\n", i.Id))
	if i.Title != "" {
		str.WriteString(fmt.Sprintf("Title: %s\n", i.Title))
	}
	if i.Description != "" {
		str.WriteString(fmt.Sprintf("Description: %s\n", i.Description))
	}

	if i.StartUnix > 0 {
		t := time.Unix(i.StartUnix, 0)
		str.WriteString(fmt.Sprintf("Start DateTime: %s\n", t))
	}

	if i.EndUnix > 0 {
		t := time.Unix(i.EndUnix, 0)
		str.WriteString(fmt.Sprintf("End DateTime: %s\n", t))
	}

	if len(i.Tags) > 0 {
		str.WriteString("Tags: ")
		for ii, t := range i.Tags {
			if ii == 0 {
				str.WriteString(fmt.Sprintf("%s", t))
			} else {
				str.WriteString(fmt.Sprintf(", %s", t))
			}
		}
		str.WriteString("\n")
	}
	return str.String()
}

func (s Schedule) GoString() string {
	var str strings.Builder
	str.WriteString(fmt.Sprintf("ID: %s\n", s.Id))
	if len(s.Scheduled) > 0 {
		str.WriteString("Items:\n")
		for _, i := range s.Scheduled {
			str.WriteString(fmt.Sprintf("%#v\n", i))
		}

	}

	return str.String()
}

func (s *Schedule) Schedule(i *Item) {
	nowUnix := time.Now().Unix()
	startFuture := nowUnix - i.StartUnix
	endFuture := nowUnix - i.EndUnix
	if startFuture < 0 && endFuture < 0 {
		s.Scheduled = append(s.Scheduled, i)
	} else {
		fmt.Printf("%d>%d and %d>%d\n", i.StartUnix, startFuture, i.EndUnix, endFuture)
	}
}

func NewItem() Item {
	i := Item{Id: genId("item")}
	return i
}

func NewSchedule() Schedule {
	s := Schedule{Id: genId("schd")}
	return s
}

func genId(prefix string) string {
	bytes := make([]byte, 15)
	_, _ = rand.Read(bytes)
	iD := base32.StdEncoding.EncodeToString(bytes)
	prefixID := fmt.Sprintf("%s_%s", prefix, iD)
	return prefixID
}

func main() {
	mainSched := NewSchedule()
	testItem := NewItem()
	testItem.Title = "I am a test"
	testItem.StartUnix = 1700000000
	testItem.EndUnix = 1700000002
	testItem.Tags = []string{"timmy", "workky"}
	mainSched.Schedule(&testItem)
	testItem2 := NewItem()
	testItem2.Title = "I am a test2"
	testItem2.Description = "Extra work"
	testItem2.StartUnix = 1700000000
	testItem2.EndUnix = 1750000002
	testItem.Description = "Pointers!"
	mainSched.Schedule(&testItem2)
	fmt.Printf("%#v", &mainSched)
}
