package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"math/rand"
	"microshop/domain/entity"
	"microshop/domain/repository"
	"microshop/domain/repository/impl"
	"microshop/infrastructure/cache"
	"microshop/infrastructure/jwtMgr"
	"microshop/utils"
	"net/http"
	"net/url"
	"time"
)

type LoginService struct {
	LoginRepo impl.LoginRepo
}

func NewLoginService() LoginService {
	return LoginService{
		LoginRepo: repository.NewLoginRepo(),
	}
}

func (o LoginService) SendSms(ctx context.Context, mobile string) (entity.SmsResult, error) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	content := "【花样百货】尊敬的会员，欢迎使用花样百货，您的验证码为：" + vcode
	query := url.Values{}
	query.Add("u", cfg.SmsUser)
	query.Add("p", cfg.SmsPassWord)
	query.Add("m", mobile)
	query.Add("c", content)
	result, err := o.send2SmsSvr(cfg.SmsSendUrl, query)
	if err == nil {
		result.Vcode = vcode
		cache.MemCachePool.Set(mobile, vcode, 30 * 60)
		cache.MemCachePool.Set(mobile + "_last", struct{}{}, 60)
	}
	return result, err
}

func (o LoginService) send2SmsSvr(url string, query url.Values) (entity.SmsResult, error) {
	return entity.SmsResult{Code: http.StatusOK, Message: "0"}, nil
	req, err := http.PostForm(url, query)
	if err != nil {
		return entity.SmsResult{Code: req.StatusCode, Message: "请求失败"}, err
	}
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return entity.SmsResult{Code: req.StatusCode, Message: "读取失败"}, err
	}
	bodyString := string(body)
	if bodyString != "0" {
		errMsg := "短信发送错误码:" + bodyString
		return entity.SmsResult{Code: http.StatusInternalServerError, Message: errMsg}, errors.New(errMsg)
	}
	return entity.SmsResult{Code: req.StatusCode, Message: bodyString}, nil
}

func (o LoginService) Login(ctx context.Context, c *gin.Context, req entity.LoginReq) (string, *entity.UserInfo, error) {
	userInfo, err := o.LoginRepo.One(ctx, req.UserInfo.Mobile)
	if err != nil {
		return "", nil, err
	}
	now := time.Now().UnixNano()/1e6
	if userInfo == nil {
		userInfo = &entity.UserInfo{
			Username: req.UserInfo.Mobile,
			Password: "",
			RegisterTime: now,
			RegisterIp: c.ClientIP(),
			Mobile: req.UserInfo.Mobile,
			Avatar: "",
			Gender: 1,
			Nickname: req.UserInfo.Mobile,
			LastLoginTime: now,
			LastLoginIp: c.ClientIP(),
			Birthday: 0,
		}
		err = o.LoginRepo.Insert(ctx, *userInfo)
		if err != nil {
			return "", nil, err
		}
		userInfo, err = o.LoginRepo.One(ctx, req.UserInfo.Mobile)
		if err != nil || userInfo == nil {
			return "", nil, errors.New("after insert query not exsit")
		}
	}else{
		err = o.LoginRepo.Update(req.UserInfo.Mobile, map[string]interface{}{
			"last_login_ip": c.ClientIP(),
			"last_login_time": now,
		})
		userInfo.LastLoginIp = c.ClientIP()
		userInfo.LastLoginTime = now
		if err != nil {
			return "", nil, err
		}
	}
	token := jwtMgr.JwtMgr{}.Create(utils.Int2String(userInfo.Id))
	return token, userInfo, nil
}