package swc

import (
	"fmt"
	"net/url"
	"path"
	"strconv"

	jsoniter "github.com/json-iterator/go"
	"github.com/levigross/grequests"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type SeaweedClient struct {
	Master string
}

type AssignResp struct {
	Count     int    `json:"count"`
	Fid       string `json:"fid"`
	URL       string `json:"url"`
	Publicurl string `json:"publicUrl"`
}

type UploadResp struct {
	Size uint64 `json:"size"`
}

func (swc SeaweedClient) Assign(count int) (AssignResp, error) {
	var ret AssignResp
	ro := &grequests.RequestOptions{
		Params: map[string]string{"count": strconv.Itoa(count)},
	}
	u, err := url.Parse(swc.Master)
	if err != nil {
		logrus.WithError(err).Error()
		return ret, err
	}
	u.Path = path.Join("dir", "assign")
	resp, err := grequests.Get(u.String(), ro)
	if err != nil {
		logrus.WithError(err).Error()
		return ret, err
	}
	defer resp.Close()
	jsoniter.UnmarshalFromString(resp.String(), &ret)
	return ret, nil
}

func (swc SeaweedClient) Upload(volume, fid string, filename string) (publicURL, f string, size uint64, err error) {
	var ret UploadResp
	fd, err := grequests.FileUploadFromDisk(filename)
	if err != nil {
		logrus.WithError(err).Error()
		return ``, ``, uint64(0), err
	}

	uploadURL := fmt.Sprintf(`http://%s/%s`, volume, fid)
	ro := &grequests.RequestOptions{
		Files: fd,
	}
	resp, err := grequests.Post(uploadURL, ro)
	if err != nil {
		logrus.WithError(err).Error()
		return ``, ``, uint64(0), err
	}
	defer resp.Close()

	jsoniter.UnmarshalFromString(resp.String(), &ret)
	return volume, fid, ret.Size, nil
}

func (swc SeaweedClient) Delete(volume, fid string) error {
	deleteURL := fmt.Sprintf(`%s%s`, volume, fid)
	resp, err := grequests.Delete(deleteURL, nil)
	if err != nil {
		logrus.WithError(err).Error()
		return err
	}
	defer resp.Close()
	return nil
}

func (swc SeaweedClient) DownloadAndUpload(originURL string) (publicURL, fid string, size uint64, err error) {
	resp, err := grequests.Get(originURL, nil)
	if err != nil {
		logrus.WithError(err).Error()
		return ``, ``, uint64(0), err
	}
	defer resp.Close()
	dest := path.Join(`/tmp`, uuid.NewV4().String())
	if err := resp.DownloadToFile(dest); err != nil {
		logrus.WithError(err).Error()
		return ``, ``, uint64(0), err
	}
	return swc.UploadSimplely(dest)
}

func (swc SeaweedClient) UploadSimplely(filename string) (publicURL, fid string, size uint64, err error) {
	assign, err := swc.Assign(1)
	if err != nil {
		logrus.WithError(err).Error()
		return ``, ``, uint64(0), err
	}
	logrus.WithField("fid", assign.Fid).WithField("Publicurl", assign.Publicurl).Info()
	return swc.Upload(assign.Publicurl, assign.Fid, filename)
}
