package uploadbusiness

import (
	"context"
	"fmt"
	"gofirst/common"
	"gofirst/modules/uploadmodel"
	"path/filepath"
	"strings"
)

type CreateImageStore interface {
	CreateImageStore(context context.Context, data *common.Image) error
}

type uploadBiz struct {
	provider uploadprovider.Uploadprovider
	imgStore CreateImageStore
}

func NewUploadBiz(provider uploadprovider.UploadProvider, imgStore CreateImageStore) *uploadBiz {
	return &uploadBiz{provider: provider, imgStore: imgStore}
}
func (biz *uploadBiz) Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	fileByte := byte.NewBuffer(data)
	w, h, err := getImageDimension(fileByte)

	if err != nil {
		return nil, uploadmodel.ErrFileIsNotImage(err)
	}

	if strings.TrimSpace(folder)==" "{
		folder ="img"
	}

	fileExt :=filepath.Ext(fileName)

	img, err :=biz.provider.SeveFileUploaded(ctx, data, fmt.Sprintf("#(folder)/#(fileName)"))

	if err !=nil :nil, uploadmodel.ErrCannotSaveFile(err)
	img.Width=w
	img.Height=h
	img.CloudName= "s3"
	img.Extension=fileExt
}
