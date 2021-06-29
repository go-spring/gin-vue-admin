package utils

import (
	"bytes"
	"fmt"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/dchest/captcha"
	"github.com/go-spring/spring-boot"
)

func init() {
	SpringBoot.RegisterBean(new(GinCaptchaService))
}

type CaptchaConfig struct {
	KeyLong   int `value:"${captcha.key-long}"`
	ImgWidth  int `value:"${captcha.img-width}"`
	ImgHeight int `value:"${captcha.img-height}"`
}

type GinCaptchaService struct {
	CaptchaConfig CaptchaConfig
}

// 这里需要自行实现captcha 的gin模式
func (service *GinCaptchaService) GinCaptchaServeHTTP(w http.ResponseWriter, r *http.Request) {
	dir, file := path.Split(r.URL.Path)
	ext := path.Ext(file)
	id := file[:len(file)-len(ext)]
	if ext == "" || id == "" {
		http.NotFound(w, r)
		return
	}
	fmt.Println("reload : " + r.FormValue("reload"))
	if r.FormValue("reload") != "" {
		captcha.Reload(id)
	}
	lang := strings.ToLower(r.FormValue("lang"))
	download := path.Base(dir) == "download"
	if service.Serve(w, r, id, ext, lang, download, service.CaptchaConfig.ImgWidth, service.CaptchaConfig.ImgHeight) == captcha.ErrNotFound {
		http.NotFound(w, r)
	}
}

func (service *GinCaptchaService) Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		_ = captcha.WriteImage(&content, id, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		_ = captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}
