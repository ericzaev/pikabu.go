package structs

type PikabuResponseError struct {
	Message string
}

type PikabuResponseStory struct {
	Error    PikabuResponseError
	Response PikabuStoryComposite
}

type PikabuResponseFeed struct {
	Error    PikabuResponseError
	Response PikabuFeedComposite
}

type PikabuResponseProfile struct {
	Error    PikabuResponseError
	Response struct {
		User PikabuProfile
	}
}

type PikabuResponseComments struct {
	Error    PikabuResponseError
	Response PikabuCommentsComposite
}

type PikabuFeedComposite struct {
	TotalStories       int   `json:"total_stories"`
	HideVisitedStories bool  `json:"hide_visited_stories"`
	HiddenStoriesCount int   `json:"hidden_stories_count"`
	QueryTime          int64 `json:"query_time"`
	Data               []PikabuStory
}

type PikabuCommentsComposite struct {
	Next bool            `json:"has_next_page_comments"`
	List []PikabuComment `json:"comments"`
}

type PikabuStoryComposite struct {
	Story                     PikabuStory
	Comments                  []PikabuComment
	HideLargeCommentsBranches bool   `json:"hide_large_comments_branches"`
	MaxCommentsBranchDepth    int    `json:"max_comments_branch_depth"`
	CommentsSortBy            string `json:"comments_sort_by"`
	HasNextPageComments       bool   `json:"has_next_page_comments"`
	QueryTime                 int64  `json:"query_time"`
}

type PikabuStory struct {
	StoryPluses          int           `json:"story_pluses"`
	StoryMinuses         int           `json:"story_minuses"`
	StoryDigs            int           `json:"story_digs"`
	CurrUserVote         int           `json:"curr_user_vote"`
	CanVote              bool          `json:"can_vote"`
	CommentsCount        int           `json:"comments_count"`
	CommunityAvatar      string        `json:"community_avatar"`
	CommunityId          int           `json:"community_id"`
	CommunityLink        string        `json:"community_link"`
	CommunityName        string        `json:"community_name"`
	CommunityIsNsfw      bool          `json:"community_is_nsfw"`
	IsCommunityModerator bool          `json:"is_community_moderator"`
	CanMoveToCommunity   bool          `json:"can_move_to_community"`
	CanMoveToCommonFeed  bool          `json:"can_move_to_common_feed"`
	IsAdult              bool          `json:"is_adult"`
	IsAuthors            bool          `json:"is_authors"`
	IsLongpost           bool          `json:"is_longpost"`
	IsRatingHidden       bool          `json:"is_rating_hidden"`
	IsSavedStory         bool          `json:"is_saved_story"`
	IsVisited            bool          `json:"is_visited"`
	HiddenRatingIcon     string        `json:"hidden_rating_icon"`
	StoryId              int           `json:"story_id"`
	StoryTime            int64         `json:"story_time"`
	StoryTitle           string        `json:"story_title"`
	Tags                 []string      `json:"tags"`
	UserApprove          string        `json:"user_approve"`
	UserAvatar           string        `json:"user_avatar"`
	UserAvatarOrDefault  string        `json:"user_avatar_or_default"`
	UserId               int           `json:"user_id"`
	UserName             string        `json:"user_name"`
	UserNameLabelCodes   []interface{} `json:"user_name_label_codes"`
	UserProfileUrl       string        `json:"user_profile_url"`
	IsUserSubscribed     bool          `json:"is_user_subscribed"`
	CanEdit              bool          `json:"can_edit"`
	CanEditCommunity     bool          `json:"can_edit_community"`
	OgCaption            string        `json:"og_caption"`
	OgDesc               string        `json:"og_desc"`
	OgImage              string        `json:"og_image"`
	OgTitle              string        `json:"og_title"`
	OgUrl                string        `json:"og_url"`
	IsHotComments        bool          `json:"is_hot_comments"`
	ParentStoryId        interface{}   `json:"parent_story_id"`
	StoryData            []struct {
		Type string      `json:"type"`
		Data interface{} `json:"data"`
	} `json:"story_data"`
	TopComment PikabuComment `json:"top_comment"`
}

type PikabuStoryData struct {
	Small     string `json:"small"`
	Large     string `json:"large"`
	Animation struct {
		Preview string `json:"preview"`
		RawUrl  string `json:"raw_url"`
		Formats struct {
		} `json:"formats"`
	} `json:"animation"`
	ImgSize []int `json:"img_size"`
}

type PikabuProfile struct {
	CurrentUserId   int    `json:"current_user_id"`
	UserId          string `json:"user_id"`
	UserName        string `json:"user_name"`
	Rating          string `json:"rating"`
	Gender          string `json:"gender"`
	CommentsCount   int    `json:"comments_count"`
	StoriesCount    int    `json:"stories_count"`
	StoriesHotCount int    `json:"stories_hot_count"`
	PlusesCount     int    `json:"pluses_count"`
	MinusesCount    int    `json:"minuses_count"`
	SignupDate      string `json:"signup_date"`
	IsRatingBan     bool   `json:"is_rating_ban"`
	Avatar          string `json:"avatar"`
	Awards          []struct {
		Id         string  `json:"id"`
		UserId     string  `json:"user_id"`
		AwardId    string  `json:"award_id"`
		AwardTitle string  `json:"award_title"`
		AwardImage string  `json:"award_image"`
		StoryId    string  `json:"story_id"`
		StoryTitle string  `json:"story_title"`
		Date       *string `json:"date"`
		IsHidden   string  `json:"is_hidden"`
		CommentId  *string `json:"comment_id"`
		Link       string  `json:"link"`
	} `json:"awards"`
	IsSubscribed        bool   `json:"is_subscribed"`
	IsIgnored           bool   `json:"is_ignored"`
	IsIgnoredInStories  bool   `json:"is_ignored_in_stories"`
	IsIgnoredInComments bool   `json:"is_ignored_in_comments"`
	Note                string `json:"note"`
	Approved            string `json:"approved"`
	Communities         []struct {
		Id        int    `json:"id"`
		Name      string `json:"name"`
		Link      string `json:"link"`
		Avatar    string `json:"avatar"`
		AvatarUrl string `json:"avatar_url"`
	} `json:"communities"`
	SubscribersCount     int           `json:"subscribers_count"`
	IsUserBanned         bool          `json:"is_user_banned"`
	IsUserFullyBanned    bool          `json:"is_user_fully_banned"`
	PublicBanHistory     []interface{} `json:"public_ban_history"`
	UserBanTime          int64         `json:"user_ban_time"`
	About                string        `json:"about"`
	IsSlowModeEnabled    bool          `json:"is_slow_mode_enabled"`
	SlowModeText         interface{}   `json:"slow_mode_text"`
	RegistrationDuration string        `json:"registration_duration"`
	CanBlockAuthor       bool          `json:"can_block_author"`
	IsAdvertBlogsUser    bool          `json:"is_advert_blogs_user"`
}

type PikabuComment struct {
	CommentId   int `json:"comment_id"`
	ParentId    int `json:"parent_id"`
	CommentTime int `json:"comment_time"`
	CommentDesc struct {
		Text   string `json:"text"`
		Images []struct {
			Small   string `json:"small"`
			Large   string `json:"large,omitempty"`
			ImgSize []int  `json:"img_size"`
		} `json:"images"`
		Videos []struct {
			Url      string  `json:"url"`
			Thumb    string  `json:"thumb"`
			Width    int     `json:"width"`
			Height   int     `json:"height"`
			Duration int     `json:"duration"`
			Ratio    float64 `json:"ratio"`
		} `json:"videos"`
	} `json:"comment_desc"`
	CommentRating                *int        `json:"comment_rating"`
	CommentPluses                *int        `json:"comment_pluses"`
	CommentMinuses               *int        `json:"comment_minuses"`
	CommentPlatform              *int        `json:"comment_platform"`
	StoryId                      int         `json:"story_id"`
	StoryUrl                     string      `json:"story_url"`
	StoryTitle                   string      `json:"story_title"`
	CanEdit                      bool        `json:"can_edit"`
	CanVote                      bool        `json:"can_vote"`
	CanReplay                    bool        `json:"can_replay"`
	CurrUserVote                 int         `json:"curr_user_vote"`
	UserId                       int         `json:"user_id"`
	UserName                     string      `json:"user_name"`
	UserGender                   string      `json:"user_gender"`
	UserAvatarUrl                string      `json:"user_avatar_url"`
	UserProfileUrl               string      `json:"user_profile_url"`
	UserProfileDeleted           bool        `json:"user_profile_deleted"`
	IsModeratorUsername          bool        `json:"is_moderator_username"`
	IsIgnoredUser                bool        `json:"is_ignored_user"`
	IsUserIgnored                bool        `json:"is_user_ignored"`
	IgnoreCode                   int         `json:"ignore_code"`
	IgnoredBy                    []string    `json:"ignored_by"`
	IsHiddenComment              bool        `json:"is_hidden_comment"`
	IsUnreaded                   bool        `json:"is_unreaded"`
	IsCommentSaved               bool        `json:"is_comment_saved"`
	IsDeleted                    bool        `json:"is_deleted"`
	IsStoryComment               bool        `json:"is_story_comment"`
	IsCurrUserCommunityModerator bool        `json:"is_curr_user_community_moderator"`
	IsCurrUserPikabuTeam         bool        `json:"is_curr_user_pikabu_team"`
	IsUserCommunityModerator     bool        `json:"is_user_community_moderator"`
	IsUserPikabuTeam             bool        `json:"is_user_pikabu_team"`
	HasUserNote                  bool        `json:"has_user_note"`
	IsOfficial                   bool        `json:"is_official"`
	UserApprove                  string      `json:"user_approve"`
	StorySubsCode                int         `json:"story_subs_code"`
	IsOwnStory                   bool        `json:"is_own_story"`
	IsComstory                   bool        `json:"is_comstory"`
	ComstoryData                 interface{} `json:"comstory_data"`
	IsHidden                     bool        `json:"is_hidden"`
	CanHide                      bool        `json:"can_hide"`
	CanReveal                    bool        `json:"can_reveal"`
	CanIgnore                    bool        `json:"can_ignore"`
	CanBlockAuthor               bool        `json:"can_block_author"`
}
