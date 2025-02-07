package service

import (
	"github.com/goTouch/TicTok_SimpleVersion/dao"
	"github.com/goTouch/TicTok_SimpleVersion/domain"
	"github.com/goTouch/TicTok_SimpleVersion/util"
	"log"
)

/**
该文件用于：
	向数据库中插入记录条数
	倒叙返回指定id的作者的所有作品，用于个人界面展示
*/
var staticRooterUrl = util.StaticRooterPrefix

// InsertVideos 向数据库插入视频信息
func InsertVideos(videoName string, title string, coverName string, userId int64, coverGenerateStatus bool) error {
	video := domain.Video{
		Id:       0,
		AuthorId: userId,
		Title:    title,
		PlayUrl:  staticRooterUrl + "static/videos" + videoName,
		CoverUrl: staticRooterUrl + "static/covers" + coverName,
		//CreatTime: time.Time{}, // TODO 这个字段 为什么没赋值呢？去研究一下
	}
	// 若生成封面失败，视频的封面地址会被替换为默认封面的地址
	if !coverGenerateStatus {
		video.CoverUrl = "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg"
	}
	//插入数据库
	if err := dao.DB.Create(video).Error; err != nil {
		log.Print("向video数据库中插入数据失败！")
		log.Println(err)
		return err
	}
	//没出错
	return nil
}

// QueryAuthorPublishedVideo 查询用户发布的视频，以展示在个人列表中
func QueryAuthorPublishedVideo(authorIdInt64 int64) (VideoList []domain.Video) {
	dao.DB.Model(&domain.Video{}).
		Where("author_id = ?", authorIdInt64).
		Order("creat_time desc"). //该字段加了索引
		Find(&VideoList)
	return
}
