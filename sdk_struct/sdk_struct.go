package sdk_struct

import "open_im_sdk/pkg/server_api_params"

////////////////////////// message/////////////////////////

// 接收消息
type MessageReceipt struct {
	GroupID     string   `json:"groupID"`
	UserID      string   `json:"userID"`
	MsgIDList   []string `json:"msgIDList"`
	ReadTime    int64    `json:"readTime"`
	MsgFrom     int32    `json:"msgFrom"`
	ContentType int32    `json:"contentType"`
	SessionType int32    `json:"sessionType"`
}

// 撤回消息
type MessageRevoked struct {
	RevokerID                   string `json:"revokerID"`
	RevokerRole                 int32  `json:"revokerRole"`
	ClientMsgID                 string `json:"clientMsgID"`
	RevokerNickname             string `json:"revokerNickname"`
	RevokeTime                  int64  `json:"revokeTime"`
	SourceMessageSendTime       int64  `json:"sourceMessageSendTime"`
	SourceMessageSendID         string `json:"sourceMessageSendID"`
	SourceMessageSenderNickname string `json:"sourceMessageSenderNickname"`
	SessionType                 int32  `json:"sessionType"`
}

// 图片信息
type ImageInfo struct {
	Width  int32  `json:"x"`
	Height int32  `json:"y"`
	Type   string `json:"type,omitempty"`
	Size   int64  `json:"size"`
}
type PictureBaseInfo struct {
	UUID   string `json:"uuid,omitempty"`
	Type   string `json:"type,omitempty"`
	Size   int64  `json:"size"`
	Width  int32  `json:"width"`
	Height int32  `json:"height"`
	Url    string `json:"url,omitempty"`
}

// 音频信息
type SoundBaseInfo struct {
	UUID      string `json:"uuid,omitempty"`
	SoundPath string `json:"soundPath,omitempty"`
	SourceURL string `json:"sourceUrl,omitempty"`
	DataSize  int64  `json:"dataSize"`
	Duration  int64  `json:"duration"`
}

// 视频信息
type VideoBaseInfo struct {
	VideoPath      string `json:"videoPath,omitempty"`
	VideoUUID      string `json:"videoUUID,omitempty"`
	VideoURL       string `json:"videoUrl,omitempty"`
	VideoType      string `json:"videoType,omitempty"`
	VideoSize      int64  `json:"videoSize"`
	Duration       int64  `json:"duration"`
	SnapshotPath   string `json:"snapshotPath,omitempty"`
	SnapshotUUID   string `json:"snapshotUUID,omitempty"`
	SnapshotSize   int64  `json:"snapshotSize"`
	SnapshotURL    string `json:"snapshotUrl,omitempty"`
	SnapshotWidth  int32  `json:"snapshotWidth"`
	SnapshotHeight int32  `json:"snapshotHeight"`
}

// 文件信息
type FileBaseInfo struct {
	FilePath  string `json:"filePath,omitempty"`
	UUID      string `json:"uuid,omitempty"`
	SourceURL string `json:"sourceUrl,omitempty"`
	FileName  string `json:"fileName,omitempty"`
	FileSize  int64  `json:"fileSize"`
}

// 消息结构
type MsgStruct struct {
	// 客户端消息ID
	ClientMsgID string `json:"clientMsgID,omitempty"`
	// 服务端消息ID
	ServerMsgID string `json:"serverMsgID,omitempty"`
	CreateTime  int64  `json:"createTime"`
	SendTime    int64  `json:"sendTime"`
	// 私聊/ 群聊
	SessionType int32  `json:"sessionType"`
	SendID      string `json:"sendID,omitempty"`
	RecvID      string `json:"recvID,omitempty"`
	MsgFrom     int32  `json:"msgFrom"`
	// 消息类型: 图片 / 视频 / ...
	ContentType      int32  `json:"contentType"`
	SenderPlatformID int32  `json:"platformID"`
	SenderNickname   string `json:"senderNickname,omitempty"`
	SenderFaceURL    string `json:"senderFaceUrl,omitempty"`
	GroupID          string `json:"groupID,omitempty"`
	// 消息内容
	Content string `json:"content,omitempty"`
	// 消息seq
	Seq          uint32                            `json:"seq"`
	IsRead       bool                              `json:"isRead"`
	Status       int32                             `json:"status"`
	OfflinePush  server_api_params.OfflinePushInfo `json:"offlinePush,omitempty"`
	AttachedInfo string                            `json:"attachedInfo,omitempty"`
	Ex           string                            `json:"ex,omitempty"`
	PictureElem  struct {
		SourcePath      string          `json:"sourcePath,omitempty"`
		SourcePicture   PictureBaseInfo `json:"sourcePicture,omitempty"`
		BigPicture      PictureBaseInfo `json:"bigPicture,omitempty"`
		SnapshotPicture PictureBaseInfo `json:"snapshotPicture,omitempty"`
	} `json:"pictureElem,omitempty"`
	SoundElem struct {
		UUID      string `json:"uuid,omitempty"`
		SoundPath string `json:"soundPath,omitempty"`
		SourceURL string `json:"sourceUrl,omitempty"`
		DataSize  int64  `json:"dataSize"`
		Duration  int64  `json:"duration"`
	} `json:"soundElem,omitempty"`
	VideoElem struct {
		VideoPath      string `json:"videoPath,omitempty"`
		VideoUUID      string `json:"videoUUID,omitempty"`
		VideoURL       string `json:"videoUrl,omitempty"`
		VideoType      string `json:"videoType,omitempty"`
		VideoSize      int64  `json:"videoSize"`
		Duration       int64  `json:"duration"`
		SnapshotPath   string `json:"snapshotPath,omitempty"`
		SnapshotUUID   string `json:"snapshotUUID,omitempty"`
		SnapshotSize   int64  `json:"snapshotSize"`
		SnapshotURL    string `json:"snapshotUrl,omitempty"`
		SnapshotWidth  int32  `json:"snapshotWidth"`
		SnapshotHeight int32  `json:"snapshotHeight"`
	} `json:"videoElem,omitempty"`
	FileElem struct {
		FilePath  string `json:"filePath,omitempty"`
		UUID      string `json:"uuid,omitempty"`
		SourceURL string `json:"sourceUrl,omitempty"`
		FileName  string `json:"fileName,omitempty"`
		FileSize  int64  `json:"fileSize"`
	} `json:"fileElem,omitempty"`
	MergeElem struct {
		Title        string       `json:"title,omitempty"`
		AbstractList []string     `json:"abstractList,omitempty"`
		MultiMessage []*MsgStruct `json:"multiMessage,omitempty"`
	} `json:"mergeElem,omitempty"`
	AtElem struct {
		Text         string     `json:"text,omitempty"`
		AtUserList   []string   `json:"atUserList,omitempty"`
		AtUsersInfo  []*AtInfo  `json:"atUsersInfo,omitempty"`
		QuoteMessage *MsgStruct `json:"quoteMessage,omitempty"`
		IsAtSelf     bool       `json:"isAtSelf"`
	} `json:"atElem,omitempty"`
	FaceElem struct {
		Index int    `json:"index"`
		Data  string `json:"data,omitempty"`
	} `json:"faceElem,omitempty"`
	LocationElem struct {
		Description string  `json:"description,omitempty"`
		Longitude   float64 `json:"longitude"`
		Latitude    float64 `json:"latitude"`
	} `json:"locationElem,omitempty"`
	CustomElem struct {
		Data        string `json:"data,omitempty"`
		Description string `json:"description,omitempty"`
		Extension   string `json:"extension,omitempty"`
	} `json:"customElem,omitempty"`
	QuoteElem struct {
		Text         string     `json:"text,omitempty"`
		QuoteMessage *MsgStruct `json:"quoteMessage,omitempty"`
	} `json:"quoteElem,omitempty"`
	NotificationElem struct {
		Detail      string `json:"detail,omitempty"`
		DefaultTips string `json:"defaultTips,omitempty"`
	} `json:"notificationElem,omitempty"`
	MessageEntityElem struct {
		Text              string           `json:"text,omitempty"`
		MessageEntityList []*MessageEntity `json:"messageEntityList,omitempty"`
	} `json:"messageEntityElem,omitempty"`
	AttachedInfoElem AttachedInfoElem `json:"attachedInfoElem,omitempty"`
}
type AtInfo struct {
	AtUserID      string `json:"atUserID,omitempty"`
	GroupNickname string `json:"groupNickname,omitempty"`
}
type AttachedInfoElem struct {
	GroupHasReadInfo          GroupHasReadInfo `json:"groupHasReadInfo,omitempty"`
	IsPrivateChat             bool             `json:"isPrivateChat"`
	HasReadTime               int64            `json:"hasReadTime"`
	NotSenderNotificationPush bool             `json:"notSenderNotificationPush"`
	MessageEntityList         []*MessageEntity `json:"messageEntityList,omitempty"`
}
type MessageEntity struct {
	Type   string `json:"type,omitempty"`
	Offset int32  `json:"offset"`
	Length int32  `json:"length"`
	Url    string `json:"url,omitempty"`
	Info   string `json:"info,omitempty"`
}

// 群已读信息
type GroupHasReadInfo struct {
	// 已读人列表
	HasReadUserIDList []string `json:"hasReadUserIDList,omitempty"`
	HasReadCount      int32    `json:"hasReadCount"`
	GroupMemberCount  int32    `json:"groupMemberCount"`
}
type NewMsgList []*MsgStruct

// Implement the sort.Interface interface to get the number of elements method
func (n NewMsgList) Len() int {
	return len(n)
}

// Implement the sort.Interface interface comparison element method
func (n NewMsgList) Less(i, j int) bool {
	return n[i].SendTime < n[j].SendTime
}

// Implement the sort.Interface interface exchange element method
func (n NewMsgList) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

type IMConfig struct {
	Platform      int32  `json:"platform"`
	ApiAddr       string `json:"api_addr"`
	WsAddr        string `json:"ws_addr"`
	DataDir       string `json:"data_dir"`
	LogLevel      uint32 `json:"log_level"`
	ObjectStorage string `json:"object_storage"` //"cos"(default)  "oss"
}

var SvrConf IMConfig

// 会话消息同步对比结构？
type CmdNewMsgComeToConversation struct {
	MsgList     []*server_api_params.MsgData
	OperationID string
	SyncFlag    int
	// 服务端最大的seq
	MaxSeqOnSvr uint32
	// 本地最大的seq
	MaxSeqOnLocal uint32
	// 当前最大的seq
	CurrentMaxSeq uint32
	// 获取消息的顺序: 升序/ 降序
	PullMsgOrder int
}

type CmdPushMsgToMsgSync struct {
	Msg         *server_api_params.MsgData
	OperationID string
}

type CmdMaxSeqToMsgSync struct {
	MaxSeqOnSvr         uint32
	OperationID         string
	MinSeqOnSvr         uint32
	GroupID2MaxSeqOnSvr map[string]uint32
}
type CmdJoinedSuperGroup struct {
	OperationID string
}

type OANotificationElem struct {
	NotificationName    string `mapstructure:"notificationName" validate:"required"`
	NotificationFaceURL string `mapstructure:"notificationFaceURL" validate:"required"`
	NotificationType    int32  `mapstructure:"notificationType" validate:"required"`
	Text                string `mapstructure:"text" validate:"required"`
	Url                 string `mapstructure:"url"`
	MixType             int32  `mapstructure:"mixType"`
	Image               struct {
		SourceUrl   string `mapstructure:"sourceURL"`
		SnapshotUrl string `mapstructure:"snapshotURL"`
	} `mapstructure:"image"`
	Video struct {
		SourceUrl   string `mapstructure:"sourceURL"`
		SnapshotUrl string `mapstructure:"snapshotURL"`
		Duration    int64  `mapstructure:"duration"`
	} `mapstructure:"video"`
	File struct {
		SourceUrl string `mapstructure:"sourceURL"`
		FileName  string `mapstructure:"fileName"`
		FileSize  int64  `mapstructure:"fileSize"`
	} `mapstructure:"file"`
	Ex string `mapstructure:"ex"`
}
type MsgDeleteNotificationElem struct {
	GroupID     string   `json:"groupID"`
	IsAllDelete bool     `json:"isAllDelete"`
	SeqList     []string `json:"seqList"`
}
