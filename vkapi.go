package govkapi

import (
	"encoding/json"
	"time"
)

type User struct {
	Id                     int    `json: "id"`
	FirstName              string `json: "first_name"`
	LastName               string `json: "last_name"`
	Deactivated            string `json: "deactivated"`
	Hidden                 int    `json: "hidden"`
	About                  string `json: "about"`
	Activities             string `json: "activities"`
	BDate                  string `json: "bdate"`
	Blacklisted            int    `jsoh: "blacklisted"`
	BlacklistedByMe        int    `json: "blacklisted_by_me"`
	Books                  string `json: "books"`
	CanPost                int    `json: "can_post"`
	CanSeeAllPost          int    `json: "can_see_all_posts"`
	CanSeeAudio            int    `gson: "can_see_audio"`
	CanSendFriendRequest   int    `json "can_send_friend_request"`
	CanWritePrivateMessage int    `json "can_write_private_message"`
	Career                 struct {
		GropeId   int    `json: "group_id"`
		Company   string `json: "company"`
		CountryId int    `json: "country_id"`
		CityId    int    `json: "city_id"`
		CityName  string `json: "city_name"`
		From      int    `json: "from"`
		Until     int    `json: "until"`
		Position  string `json: "position"`
	} `json 'career"`
	City struct {
		Id    int    `json: "id"`
		Title string `json: "title"`
	} `json: "city"`
	CommonCount int    `json: "common_count"`
	Connections string `json: "connections"` //need check
	Contacts    struct {
		MobilePhone string `json: "mobile_phone"`
		HomePhone   string `json: "home_phone"`
	} `json: "contacts"`
	Counters struct {
		Albums        int `json: "albums"`
		Videos        int `json: "videos"`
		Audios        int `json: "audios"`
		Photos        int `json: "photos"`
		Notes         int `json: "notes"`
		Friends       int `json: "friends"`
		Groups        int `json: "groups"`
		OnlineFriends int `json: "online_friends"`
		MutualFriends int `json: "mutual_friends"`
		UserVideos    int `json: "user_videos"`
		Followers     int `json: "followers"`
		Pages         int `json: "pages"`
	} `json: "counters"`
	Country struct {
		Id    int    `json: "id"`
		Title string `json: "title"`
	} `json: "country"`
	CropPhoto struct {
		Photo Photo `json: "photo"`
		Crop  struct {
			X  json.Number `json: "x"`
			Y  json.Number `json: "y"`
			X2 json.Number `json: "x2"`
			Y2 json.Number `json: "y2"`
		} `json: "crop_photo"`
	}
	Domain    string `json: "domain"`
	Education struct {
		University     int    `json: "university"`
		UniversityName string `json: "university_name"`
		Faculty        int    `json: "faculty"`
		FacultyName    string `json: "faculty_name"`
		Graduation     int    `json: "graduation"`
	}
	Exports          string `json: "exports"`
	FirstNameNom     string `json: "first_name_nom"`
	FirstNameGen     string `json: "first_name_gen"`
	FirstNameDat     string `json: "first_name_dat"`
	FirstNameAcc     string `json: "first_name_acc"`
	FirstNameIns     string `json: "first_name_ins"`
	FirstNameAbl     string `json: "first_name_abl"`
	FollowersCount   int    `json: "followers_count"`
	FriendStatus     int    `json: "friend_status"`
	Games            string `json: "games"`
	HasMobile        int    `json: "has_mobile"`
	HasPhoto         int    `json: "has_photo"`
	HomeTown         string `json: "home_town"`
	Interest         string `json: "interests"`
	IsFavorite       int    `json: "is_favorite"`
	IsFriend         int    `json: "is_friend"`
	IsHiddenFromFeed int    `json: "is_hidden_from_feed"`
	LastNameNom      string `json: "last_name_nom"`
	LastNameGen      string `json: "last_name_gen"`
	LastNameDat      string `json: "last_name_dat"`
	LastNameAcc      string `json: "last_name_acc"`
	LastNameIns      string `json: "last_name_ins"`
	LastNameAbl      string `json: "last_name_abl"`
	LastSeen         struct {
		Time     int `json: "time"` //время последнего посещения в формате Unixtime.
		Platform int `json: "platform"`
	} `json: "last_seen"`
	Lists      string `json: "lists"` //разделенные запятой идентификаторы списков друзей, в которых состоит пользователь. Поле доступно только для метода friends.get.
	MaidenName string `json: "maiden_name"`
	Military   struct {
		Unit      string `json: "unit"`
		UnitId    int    `json: "unit_id"`
		CountryId int    `json: "country_id"`
		From      int    `json: "from"`
		Until     int    `json: "until"`
	} `json: "military"`
	Movies     string `json: "movies "`
	Music      string `json: "music "`
	Nickname   string `json: "nickname"`
	Occupation struct {
		Type string `json: "type"`
		Id   int    `json: "id"`
		Name string `json: "name"`
	} `json: "occupation"`
	Online   int `json: "online"`
	Personal struct {
		Political  int      `json: "political"`
		Langs      []string `json: "langs"`
		Religion   string   `json: "religion"`
		InspiredBy string   `json: "inspired_by"`
		PeopleMain int      `json: "people_main"`
		LifeMain   int      `json: "life_main"`
		Smoking    int      `json: "smoking"`
		Alcohol    int      `json: "alcohol"`
	} `json: 'personal'`
	Photo50       string `json: "photo_50"`
	Photo100      string `json: "photo_100"`
	Photo200orig  string `json: "photo_200_orig"`
	Photo200      string `json: "photo_200"`
	Photo400Oorig string `json: "photo_400_orig"`
	PhotoId       string `json: "photo_id"`
	PhotoMax      string `json: "photo_max"`
	PhotoMaxOrig  string `json: "photo_max_orig"`
	Quotes        string `json: "quotes"`
	Relatives     []struct {
		Id   int    `json: "id"`
		Name string `json: "name"`
		Type string `json: "type"`
	} `json: "relatives"`
	Relation int `json: "relation"`
	Schools  []struct {
		Id            string `json: "id"`
		Country       int    `json: "country"`
		City          int    `json: "city"`
		Name          string `json: "name"`
		YearFrom      int    `json: "year_from"`
		YearTo        int    `json: "year_to"`
		YearGraduated int    `json: "year_graduated"`
		Class         string `json: "class"`
		Speciality    string `json: "speciality"`
		Type          int    `json: "type"`
		TypeStr       string `json: "type_str"`
	} `json: "schools"`
	ScreenName   string `json: "screen_name"`
	Sex          int    `json: "sex"`
	Site         string `json: "site"`
	Status       string `json: "status"`
	TimeZone     int    `json: "timezone"`
	TV           string `json: "tv"`
	Universities []struct {
		Id              int    `json: "id"`
		Country         int    `json: "country"`
		City            int    `json: "city"`
		Name            string `json: "name"`
		Faculty         int    `json: "faculty"`
		FacultyName     string `json: "faculty_name"`
		Chair           int    `json: "chair"`
		ChairName       string `json: "chair_name "`
		Graduation      int    `json: "graduation"`
		EducationForm   int    `json: "education_form "`
		EducationStatus string `json: "education_status"`
	} `json: "universities"`
	Verified     int `json: "verified"`
	WallComments int `json: "wall_comments"`
	Check        time.Time
}

func GetUsers(client VKClient, ids []int, names []string, nameCase string, fields ...OptionalUserFields) []User {
	query := "do it" // Need to check
	url := "https://api.vk.com/method/users.get?PARAMETERS&" + query + "&access_token=" + client.AccessToken + "&v=" + vlient.APIVersion
	resp, err := client.Get(url)
}
func SearchUsers(client VKClient, q string, sort, offset, count int, parameters map[OptionalSearchParameters]string, fields ...OptionalUserFields) []User {

}

type OptionalUserFields string

const (
	PhotoId                OptionalUserFields = "photo_id"
	Verified               OptionalUserFields = "verified"
	Sex                    OptionalUserFields = "sex"
	BDate                  OptionalUserFields = "bdate"
	City                   OptionalUserFields = "city"
	Country                OptionalUserFields = "country"
	HomeTown               OptionalUserFields = "home_town"
	HasPhoto               OptionalUserFields = "has_photo"
	Photo50                OptionalUserFields = "photo_50"
	Photo100               OptionalUserFields = "photo_100"
	Photo200               OptionalUserFields = "photo_200"
	Photo200orig           OptionalUserFields = "photo_200_orig"
	Photo200               OptionalUserFields = "photo_200"
	Photo400Oorig          OptionalUserFields = "photo_400_orig"
	PhotoMax               OptionalUserFields = "photo_max"
	PhotoMaxOrig           OptionalUserFields = "photo_max_orig"
	Online                 OptionalUserFields = "online"
	Lists                  OptionalUserFields = "lists"
	Domain                 OptionalUserFields = "domain"
	HasMobile              OptionalUserFields = "has_mobile"
	Contacts               OptionalUserFields = "contacts"
	Site                   OptionalUserFields = "site"
	Education              OptionalUserFields = "education"
	Universities           OptionalUserFields = "universities"
	Schools                OptionalUserFields = "schools"
	Status                 OptionalUserFields = "status"
	LastSeen               OptionalUserFields = "last_seen"
	FollowersCount         OptionalUserFields = "followers_count"
	CommonCount            OptionalUserFields = "common_count"
	Occupation             OptionalUserFields = "occupation"
	Nickname               OptionalUserFields = "nickname"
	Relatives              OptionalUserFields = "relatives"
	Relation               OptionalUserFields = "relation"
	Personal               OptionalUserFields = "personal"
	Connections            OptionalUserFields = "connections"
	Exports                OptionalUserFields = "exports"
	WallComments           OptionalUserFields = "wall_comments"
	Interests              OptionalUserFields = "interests"
	Music                  OptionalUserFields = "music"
	Movies                 OptionalUserFields = "movies"
	TV                     OptionalUserFields = "tv"
	Books                  OptionalUserFields = "books"
	Games                  OptionalUserFields = "games"
	About                  OptionalUserFields = "about"
	Quotes                 OptionalUserFields = "quotes"
	CanPost                OptionalUserFields = "can_post,"
	CanSeeAllPosts         OptionalUserFields = "can_see_all_posts"
	CanSeeAudio            OptionalUserFields = "can_see_audio"
	CanWritePrivateMessage OptionalUserFields = "can_write_private_message"
	CanSendFriendRequest   OptionalUserFields = "can_send_friend_request"
	IsFavorite             OptionalUserFields = "is_favorite"
	IsHiddenFromFeed       OptionalUserFields = "is_hidden_from_feed"
	Timezone               OptionalUserFields = "timezone"
	ScreenName             OptionalUserFields = "screen_name"
	MaidenName             OptionalUserFields = "maiden_name"
	CropPhoto              OptionalUserFields = "crop_photo"
	IsFriend               OptionalUserFields = "is_friend"
	FriendStatus           OptionalUserFields = "friend_status"
	Career                 OptionalUserFields = "career"
	Military               OptionalUserFields = "military"
	Blacklisted            OptionalUserFields = "blacklisted"
	BlacklistedByMe        OptionalUserFields = "blacklisted_by_me"
)

type OptionalSearchParameters string

const (
	City              OptionalSearchParameters = "city"               //положительное число
	Country           OptionalSearchParameters = "country"            //положительное число
	HomeTown          OptionalSearchParameters = "hometown"           //строка
	UniversityCountry OptionalSearchParameters = "university_country" //положительное число
	University        OptionalSearchParameters = "university"         //положительное число
	UniversityYear    OptionalSearchParameters = "university_year"    //положительное число
	UniversityFaculty OptionalSearchParameters = "university_faculty" //положительное число
	UniversityChair   OptionalSearchParameters = "university_chair"   //положительное число
	Sex               OptionalSearchParameters = "sex"                /*положительное число пол. Возможные значения:

	  1 — женщина;
	  2 — мужчина;
	  0 — любой (по умолчанию).*/
	Status OptionalSearchParameters = "status" /*положительное число семейное положение. Возможные значения:

	  1 — не женат (не замужем);
	  2 — встречается;
	  3 — помолвлен(-а);
	  4 — женат (замужем);
	  5 — всё сложно;
	  6 — в активном поиске;
	  7 — влюблен(-а). */

	AgeFrom       OptionalSearchParameters = "age_from"       //положительное число
	AgeTo         OptionalSearchParameters = "age_to"         //положительное число
	BirthDay      OptionalSearchParameters = "birth_day"      //положительное число
	BirthMonth    OptionalSearchParameters = "birth_month"    //положительное число
	BirthYear     OptionalSearchParameters = "birth_year"     //положительное число
	Online        OptionalSearchParameters = "online"         //положительное число
	HasPhoto      OptionalSearchParameters = "has_photo"      //положительное число
	SchoolCountry OptionalSearchParameters = "school_country" //положительное число
	SchoolCity    OptionalSearchParameters = "school_city"    //положительное число
	SchoolClass   OptionalSearchParameters = "school_class"   //положительное число
	School        OptionalSearchParameters = "school"         //положительное число
	SchoolYear    OptionalSearchParameters = "school_year"    //положительное число
	Religion      OptionalSearchParameters = "religion"       //строка
	Interests     OptionalSearchParameters = "interests"      //строка
	Company       OptionalSearchParameters = "company"        //строка
	Position      OptionalSearchParameters = "position"       //строка
	GroupId       OptionalSearchParameters = "group_id"       //положительное число
	FromList      OptionalSearchParameters = "from_list"      /*список слов, разделенных через запятую Разделы среди которых нужно осуществить поиск, перечисленные через запятую. Возможные значения:

	  friends — искать среди друзей;
	  subscriptions — искать среди друзей и подписок пользователя.*/

)

type Photo struct {
	Id        int    `json: "id"`
	AlbumId   int    `json: "album_id"`
	OwnerId   int    `json: "owner_id"`
	UserId    int    `json: "user_id"`
	Text      int    `json: "text"`
	Date      int    `json: "date"` //дата добавления в формате unixtime.
	Photo75   string `json: "photo_75"`
	Photo130  string `json: "photo_130"`
	Photo604  string `json: "photo_604"`
	Photo807  string `json: "photo_807"`
	Photo1280 string `json: "photo_1280"`
	Photo2560 string `json: "photo_2560"`
	Width     int    `json: "width*"`
	Height    int    `json: "height*"`
	Check     time.Time
}

type Group struct {
	Id          int    `json: "id"`
	Name        string `json: "name"`
	ScreenName  string `json: "screen_name"`
	IsClosed    int    `json: "is_closed"`
	Deactivated string `json: "deactivated"`
	IsAdmin     int    `json: "is_admin"`
	AdminLevel  int    `json: "admin_level"`
	IsMember    int    `json: "is_member"`
	InvitedBy   int    `json: "invited_by"`

	HasPhoto int    `json: "has_photo"`
	Photo50  string `json: "photo_50"`
	Photo100 string `json: "photo_100"`
	Photo200 string `json: "photo_200"`
	BanInfo  struct {
		EndDate int    `json: "end_date"`
		Comment string `json: "comment"`
	} `json: "ban_info"`
	MemberStatus int `json: "member_status"`
	City         struct {
		Id    int    `json: "id"`
		Title string `json: "title"`
	} `json: "city"`
	Country struct {
		Id    int    `json: "id"`
		Title string `json: "title"`
	} `json: "country"`
	Place struct {
		Id        int         `json: "id"`
		Title     string      `json: "title"`
		Lattitude json.Number `json: "latitude"`
		Longitude json.Number `json: "longitude"`
		Type      string      `json: "type"`
		Country   int         `json: "country"`
		City      int         `json: "city"`
		Address   string      `json: "address"`
	}
	Description  string `json: "description"`
	WikiPage     string `json: "wiki_page"`
	MembersCount int    `json: "members_count"`
	Counters     struct {
		Photos int `json: "photos"`
		Albums int `json: "albums"`
		Audios int `json: "audios"`
		Videos int `json: "videos"`
		Topics int `json: "topics"`
		Docs   int `json: "docs"`
	} `json: "counters"`
	StartDate       int    `json: "start_date"` //для встреч содержат время начала и окончания встречи в формате unixtime. Для публичных страниц содержит только start_date — дата основания в формате YYYYMMDD.
	FinishDate      int    `json: "finish_date"`
	PublicDateLabel string `json: "public_date_label"` //Текст описания для поля start_date.
	CanMessage      int    `json: "can_message"`
	CanPost         int    `json: "can_post"`
	CanSeeAllPost   int    `json: "can_see_all_posts"`
	CanUploadDoc    int    `json: "can_upload_doc"`
	CanUploadVideo  int    `json: "can_upload_video"`
	CanCreateTopic  int    `json: "can_create_topic"`
	Activity        string `json: "activity"`
	Status          string `json: "status"`
	Contacts        []struct {
		UserId int    `json: "user_id"`
		Desc   string `json: "desc"`
		Phone  string `json: "phone"`
		Email  string `json: "email"`
	} `json: "contacts"`
	Links []struct {
		Id       int    `json: "id"`
		URL      string `json: "url"`
		Name     string `json: "name"`
		Desc     string `json: "desc"`
		Photo58  string `json: "photo_50"`
		Photo100 string `json: "photo_100"`
	} `json: "links"`
	FixedPost        int    `json: "fixed_post"`
	Verified         int    `json: "verified"`
	Site             string `json: "site"`
	IsFavorite       int    `json: "is_favorite"`
	IsHiddenFromFeed int    `json: "is_hidden_from_feed"`
	MainSection      int    `json: "main_section"`
	Market           struct {
		Enabled     int `json: "enabled"`
		PriceMin    int `json: "price_min"`
		PriceMax    int `json: "price_max"`
		MainAlbunId int `json: "main_album_id"`
		ContactId   int `json: "contact_id"`
		Currency    struct {
			Id   int    `json: "id"`
			Name string `json: "name"`
		} `json: "currency"`
		CurrencText string `json: "currency_text"`
	}
	AgeLimits int `json: "age_limits "`
	Check     time.Time
}
type Audio struct {
	Id       int    `json: "id"`
	OwnerId  int    `json: "owner_id"`
	Artist   string `json: "artist"`
	Title    string `json: "title"`
	Duration int    `json: "duration"`
	URL      string `json: "url"`
	LyricsId int    `json: "lyrics_id"`
	AlbumId  int    `json: "album_id"`
	GenreId  int    `json: "genre_id"`
	Date     int    `json: "date"`
	NoSearch int    `json: "no_search"`
	Check    time.Time
}

type Video struct {
	Id          int    `json: "id"`
	OwnerId     int    `json: "owner_id"`
	Title       string `json: "title"`
	Description string `json: "description"`
	Duration    int    `json: "duration"`
	Photo130    string `json: "photo_130"`
	Photo320    string `json: "photo_320"`
	Photo640    string `json: "photo_640"`
	Photo800    string `json: "photo_800"`
	Date        int    `json: "date"`        //дата добавления в формате unixtime.
	AddingDate  int    `json: "adding_date"` //дата добавления в формате unixtime.
	Views       int    `json: "views"`
	Comments    int    `json: "comments"`
	Player      string `json: "player"`
	AccessKey   string `json: "access_key"`
	Processing  int    `json: "processing"`
	Live        int    `json: "live"`
	Check       time.Time
}

type Doc struct {
	Id       int    `json: "id"`
	OwnerId  int    `json: "owner_id"`
	Title    string `json: "title"`
	Size     int    `json: "size"`
	Ext      string `json: "ext"`
	URL      string `json: "url"`
	Photo100 string `json: "photo_100"`
	Photo130 string `json: "photo_130"`
	Date     int    `json: "date"` //дата добавления в формате unixtime.
	Type     int    `json: "type"`
	Check    time.Time
}

type Post struct {
	Id           int    `json: "id"`
	OwnerId      int    `json: "owner_id"`
	FromId       int    `json: "from_id"`
	Date         int    `json: "date"` //дата добавления в формате unixtime.
	Text         string `json: "text"`
	ReplyOwnerId int    `json: "reply_owner_id"`
	ReplyPostId  int    `json: "reply_post_id"`
	FriendsOnly  int    `json: "friends_only"`
	Comments     struct {
		Count   int `json: "count"`
		CanPost int `json: "can_post*"`
	} `json: "comments"`
	Likes struct {
		Count      int `json: "count"`
		UserLikes  int `json: "user_likes*"`
		CanLike    int `json: "can_like* "`
		CanPublish int `json: "can_publish* "`
	} `json: "likes"`
	Reposts struct {
		Count        int `json: "count"`
		UserReposted int `json: "user_reposted*"`
	} `json: "reposts"`
	PostType     string       `json: "post_type"`
	PostSource   PostSource   `json: "post_source*"`
	Attachments  []Attachment `json: "attachments"`
	Geo          Geo          `json: "geo"`
	SignedId     int          `json: "signer_id"`
	CopyHistory  []*Post      `json: "copy_history"` //массив, содержащий историю репостов для записи. Возвращается только в том случае, если запись является репостом. Каждый из объектов массива, в свою очередь, является объектом-записью стандартного формата.   Need to check
	CanPin       int          `json: "can_pin"`
	CanDelete    int          `json: "can_delete"`
	CanEdit      int          `json: "can_edit"`
	IisPinnedD   int          `json: "is_pinned"`
	CopyPostDate int          `json: "copy_post_date"` //время публикации записи-оригинала в формате unixtime (если запись является копией записи с чужой стены).
	CopyPostType string       `json: "copy_post_type"`
	CopyOwnerId  int          `json: "copy_owner_id"`
	CopyPostId   int          `json: "copy_post_id"`
	CopyText     string       `json: "copy_text"`
	MarkedAsAds  int          `json: "marked_as_ads"`
	Check        time.Time
}

type Message struct {
	Id           int          `json: "id"`
	UserId       int          `json: "user_id"`
	FromId       int          `json: "from_id"`
	Date         int          `json: "date"` //дата добавления в формате unixtime.
	ReadState    int          `json: "read_state"`
	Out          int          `json: "out"`
	Title        string       `json: "title"`
	Body         string       `json: "body"`
	Geo          Geo          `json: "geo"`
	Attachments  []Attachment `json: "attachments"`  //Need to check
	FwdMessages  []Message    `json: "fwd_messages"` //Need to check
	Emoji        int          `json: "emoji"`
	Important    int          `json: "important"`
	Deleted      int          `json: "deleted"`
	RandomId     string       `json: "random_id"`
	ChatId       int          `json: "chat_id"`
	ChatActive   []string     `json: "chat_active"` //Need to check
	PushSettings string       `json: "push_settings"`
	UsersCount   int          `json: "users_count"`
	AdminId      int          `json: "admin_id"`
	Action       string       `json: "action"`
	ActionMid    int          `json: "action_mid"` //Need to check
	ActionEmail  string       `json: "action_email"`
	ActionText   string       `json: "action_text"`
	Photo50      string       `json: "photo_50"`
	Photo100     string       `json: "photo_100"`
	Photo200     string       `json: "photo_200"`
	Check        time.Time
}

type Geo struct {
	Type        string `json: "type"`
	Coordinates string `json: "coordinates"` //Need to check
	Place       struct {
		Id         int    `json: "id"`
		Title      string `json: "title"`
		Latitude   string `json: "latitude"`  //Need to check
		Longitude  string `json: "longitude"` //Need to check
		Created    int    `json: "created"`   //Need to check
		Icon       string `json: "icon"`
		Country    string `json: "country"`
		City       string `json: "city"`
		Type       string `json: "type"` //тип чекина;  Need to check
		GroupId    int    `json: "group_id"`
		GroupPhoto string `json: "group_photo"` //миниатюра аватара сообщества;   Need to check
		Checkins   int    `json: "checkins"`
		Updated    int    `json: "updated"` //время последнего чекина;  Need to check
		Address    string `json: "address"`
	} `json: "place"`
	Check time.Time
}
type Attachment struct {
	Type  string          `json: "type"`
	Raw   jsom.RawMessage //Need to check
	Check time.Time
}

type Chat struct {
	Id           int    `json: "id"`
	Type         string `json: "type"`
	Title        string `json: "title"`
	AdminId      int    `json: "admin_id"`
	Users        []int  `json: "users"`
	PushSettings struct {
		Sound         int `json: "sound"`
		DisabledUntil int `json: "disabled_until "`
	} `json: "push_settings"`
	Photo50  string `json: "photo_50"`
	Photo100 string `json: "photo_100"`
	Photo200 string `json: "photo_200"`
	Left     int    `json: "left"`
	Kicked   int    `json: "kicked"`
	Check    time.Time
}
type Comment struct {
	Id             int          `json: "id"`
	FromId         int          `json: "from_id"`
	Date           int          `json: "date"` //дата добавления в формате unixtime.
	Text           string       `json: "text"`
	ReplyToUser    int          `json: "reply_to_user"`
	ReplyToComment int          `json: "reply_to_comment	"`
	Attachments    []Attachment `json: "attachments"` //Need to check
	Likes          struct {
		Count     int `json: "count"`
		UserLikes int `json: "user_likes"`
		CanLike   int `json: "can_like"`
	}
	Check time.Time
}

type Note struct {
	Id           int       `json: "id"`
	OwnerId      int       `json: "owner_id"`
	Title        string    `json: "title"`
	Text         string    `json: "text"`
	Date         int       `json: "date"` //дата добавления в формате unixtime.
	Comments     []Comment `json: "comments"`
	ReadComments int       `json: "read_comments"`
	ViewURL      string    `json: "view_url"`
	Check        time.Time
}
type Page struct {
	Id                       int    `json: "id"`
	GroupId                  int    `json: "group_id"`
	CreatorId                int    `json: "creator_id"`
	Title                    string `json: "title"`
	CurrentUserCanEdit       int    `json: "current_user_can_edit"`
	CurrentUserCanEditAccess int    `json: "current_user_can_edit_access"`
	WhoCanView               int    `json: "who_can_view"`
	WhoCanEdit               int    `json: "who_can_edit"`
	Edited                   int    `json: "edited"`
	Created                  int    `json: "created"`
	EditorId                 int    `json: "editor_id"`
	Views                    int    `json: "views"`
	Parent                   string `json: "parent"`
	Parent2                  string `json: "parent2"`
	Source                   string `json: "source"`
	Html                     string `json: "html"`
	ViewURL                  string `json: "view_url"`
	Check                    time.Time
}
type Item struct {
	Id          int    `json: "id"`
	OwnerId     int    `json: "owner_id"`
	Title       string `json: "title"`
	Description string `json: "description"`
	Price       struct {
		Amount   int `json: "amount"`
		Currency struct {
			Id   int    `json: "id"`
			Name string `json: "name"`
		} `json: "currency"`
		Text string `json: "text"`
	} `json: "price"`
	Category struct {
		Id      int    `json: "id"`
		Name    string `json: "name"`
		Section struct {
			Id   int    `json: "id"`
			Name string `json: "name"`
		} `json: "section"`
	} `json: "category"`
	ThumbPhoto   string  `json: "thumb_photo"`
	Date         int     `json: "date"` //дата добавления в формате unixtime.
	Availability int     `json: "availability"`
	Photos       []Photo `json: "photos"`
	CanComment   int     `json: "can_comment"`
	CanRepost    int     `json: "can_repost"`
	Likes        struct {
		UserLikes int `json: "user_likes"`
		Count     int `json: "count"`
	} `json: "likes"`
	Check time.Time
}
type Album struct {
	Id          int    `json: "id"`
	OwnerId     int    `json: "owner_id"`
	Title       string `json: "title"`
	Photo       Photo  `json: "photo"`
	Count       int    `json: "count"`
	UpdatedTime int    `json: "updated_time"`
	Check       time.Time
}
type PostSource struct {
	Type     string `json: "type"`
	Platform string `json: "platform"`
	Data     string `json: "data"`
	URL      string `json: "url"`
}

type PrivacyView []json.RawMessage

type PushSettings struct {
	Msg            []string `json: "msg"`
	Chat           []string `json: "chat"`
	Friend         []string `json: "friend"`
	FriendFound    string   `json: "friend_found"`
	FriendAccepted string   `json: "friend_accepted"`
	Reply          string   `json: "reply"`
	Comment        []string `json: "comment"`
	Mention        []string `json: "mention"`
	Like           []string `json: "like"`
	Repost         []string `json: "repost"`
	WallPost       string   `json: "wall_post"`
	WallPublish    string   `json: "wall_publish"`
	GroupInvite    string   `json: "group_invite"`
	GroupAccepted  string   `json: "group_accepted"`
	EventSoon      string   `json: "event_soon"`
	TagPhoto       []string `json: "tag_photo"`
	AppRequest     string   `json: "app_request "`
	SdkOpen        string   `json: "sdk_open"`
	NewPost        string   `json: "new_post"`
	Birthday       string   `json: "birthday"`
	Check          time.Time
}

type Topic struct {
	Id           int    `json: "id"`
	Title        string `json: "title"`
	Created      int    `json: "created"` //дата создания обсуждения в Unixtime.
	CreatedBy    int    `json: "created_by"`
	Updated      int    `json: "updated"` //дата создания обсуждения в Unixtime.
	UpsatedBy    int    `json: "updated_by"`
	IsClosed     int    `json: "is_closed"`
	IsFixed      int    `json: "is_fixed"`
	Comments     int    `json: "comments"`
	FirstComment string `json: "first_comment"`
	LastComment  string `json: "last_comment"`
	Check        time.Time
}
type Tag struct {
	UserId     int         `json: "user_id"`
	Id         int         `json: "id"`
	PlacerId   int         `json: "placer_id"`
	TaggedName string      `json: "tagged_name"`
	Date       int         `json: "date"` //дата добавления в формате unixtime.
	X          json.Number `json: "x"`
	Y          json.Number `json: "y"` //Need to check
	X2         json.Number `json: "x2"`
	Y2         json.Number `json: "y2"`
	ViewedD    int         `json: "viewed"`
}