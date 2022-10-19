package rest

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"os"
	"pikabu_api.go/libs"
	"pikabu_api.go/rest/structs"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

func pikabuClient(action string, data libs.Params) (*resty.Request, string) {
	body, _ := json.Marshal(pikabuPack(action, data))
	client := resty.New().R()

	client.SetBody(body).
		SetHeader("User-Agent", os.Getenv("PIKABU_AGENT")).
		SetHeader("Content-Type", "application/json; charset=utf-8")

	return client, fmt.Sprintf("%s/%s", os.Getenv("PIKABU_URL"), action)
}

func pikabuPack(action string, data libs.Params) libs.Params {
	var values []string

	data["new_sort"] = 1

	pikabuEncode(data, &values)

	sort.Strings(values)

	data["token"] = strconv.FormatInt(time.Now().Unix(), 10)
	data["hash"] = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%x", md5.Sum([]byte(strings.Join(
		[]string{os.Getenv("PIKABU_SALT"), action, data["token"].(string),
			strings.Join(values[:], ",")}, ","))))))
	data["id"] = "iws"

	return data
}

func pikabuEncode(data interface{}, result *[]string) {
	typeOf := reflect.TypeOf(data).Kind().String()

	if typeOf == "map" {
		for _, value := range data.(libs.Params) {
			pikabuEncode(value, result)
		}
	} else if typeOf == "slice" {
		for _, value := range data.([]interface{}) {
			pikabuEncode(value, result)
		}
	} else {
		*result = append(*result, fmt.Sprintf("%v", data))
	}
}

func PikabuFeed(cmd string, params ...libs.Params) structs.PikabuFeedComposite {
	var result structs.PikabuResponseFeed

	client, url := pikabuClient("feed", libs.MergeParams(libs.Params{"cmd": cmd, "page": 1}, params...))
	client.SetResult(&result).Post(url)

	return result.Response
}

func PikabuStory(id int, params ...libs.Params) structs.PikabuStoryComposite {
	var result structs.PikabuResponseStory

	client, url := pikabuClient("story.get", libs.MergeParams(libs.Params{"page": 1, "story_id": id}, params...))
	client.SetResult(&result).Post(url)

	return result.Response
}

func PikabuProfile(user string, params ...libs.Params) structs.PikabuProfile {
	var result structs.PikabuResponseProfile

	client, url := pikabuClient("user.profile.get", libs.MergeParams(libs.Params{"user_name": user}, params...))
	client.SetResult(&result).Post(url)

	return result.Response.User
}

func PikabuComments(id int, params ...libs.Params) structs.PikabuCommentsComposite {
	var result structs.PikabuResponseComments

	client, url := pikabuClient("story.get", libs.MergeParams(libs.Params{"page": 1, "story_id": id}, params...))
	client.SetResult(&result).Post(url)

	return result.Response
}

func PikabuSearch(params libs.Params) structs.PikabuFeedComposite {
	var result structs.PikabuResponseFeed

	client, url := pikabuClient("search", params)
	client.SetResult(&result).Post(url)

	return result.Response
}

func PikabuStoryProfile(user string, params libs.Params) structs.PikabuFeedComposite {
	var result structs.PikabuResponseFeed

	client, url := pikabuClient("story.profile.get", libs.MergeParams(libs.Params{"page": 1, "user_name": user}, params))
	client.SetResult(&result).Post(url)

	return result.Response
}
