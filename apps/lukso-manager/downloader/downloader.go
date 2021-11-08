package downloader

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"lukso/shared"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type GithubReleases []struct {
	URL             string    `json:"url"`
	AssetsURL       string    `json:"assets_url"`
	UploadURL       string    `json:"upload_url"`
	HTMLURL         string    `json:"html_url"`
	ID              int       `json:"id"`
	Author          Author    `json:"author"`
	NodeID          string    `json:"node_id"`
	TagName         string    `json:"tag_name"`
	TargetCommitish string    `json:"target_commitish"`
	Name            string    `json:"name"`
	Draft           bool      `json:"draft"`
	Prerelease      bool      `json:"prerelease"`
	CreatedAt       time.Time `json:"created_at"`
	PublishedAt     time.Time `json:"published_at"`
	Assets          []Assets  `json:"assets"`
	TarballURL      string    `json:"tarball_url"`
	ZipballURL      string    `json:"zipball_url"`
	Body            string    `json:"body"`
}
type Author struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}
type Uploader struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}
type Assets struct {
	URL                string    `json:"url"`
	ID                 int       `json:"id"`
	NodeID             string    `json:"node_id"`
	Name               string    `json:"name"`
	Label              string    `json:"label"`
	Uploader           Uploader  `json:"uploader"`
	ContentType        string    `json:"content_type"`
	State              string    `json:"state"`
	Size               int       `json:"size"`
	DownloadCount      int       `json:"download_count"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	BrowserDownloadURL string    `json:"browser_download_url"`
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func downloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

var ReleaseLocations = map[string]string{
	"pandora":      "https://api.github.com/repos/lukso-network/pandora-execution-engine/releases",
	"vanguard":     "https://api.github.com/repos/lukso-network/vanguard-consensus-engine/releases",
	"orchestrator": "https://api.github.com/repos/lukso-network/lukso-orchestrator/releases",
}

type updateRequestBody struct {
	Client  string
	Version string
	Url     string
}

func GetDownloadedVersions(w http.ResponseWriter, r *http.Request) {
	var DownloadedVerions = map[string][]string{}

	downloads := shared.UserHomeDir + "/.lukso/downloads"

	err := filepath.Walk(downloads,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				fmt.Println("path, info.Size()")
				fmt.Println(path, info.Size())
				pathParts := strings.Split(path, "downloads")
				fileNameParts := strings.Split(pathParts[1], "/")
				DownloadedVerions[fileNameParts[1]] = append(DownloadedVerions[fileNameParts[1]], fileNameParts[2])
			}

			return nil
		})
	if err != nil {
		log.Println(err)
	}

	jsonString, err := json.Marshal(DownloadedVerions)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonString)
}

func GetAvailableVersions(w http.ResponseWriter, r *http.Request) {
	confMap := map[string]map[string]string{}
	for client, url := range ReleaseLocations {
		r, err := http.Get(url + "?per_page=3")
		if err != nil {
			log.Fatalln("Request to "+url+" failed.", err)
		}

		confMap[client] = make(map[string]string)

		decoder := json.NewDecoder(r.Body)
		var releases GithubReleases

		err2 := decoder.Decode(&releases)
		if err2 != nil {
			log.Fatalln(err)
		}

		for _, v := range releases {
			a := getDownloadUrlFromAsset(client, v.Assets)
			if a != "" {
				confMap[client][v.TagName] = a
			}
		}
	}
	jsonString, err := json.Marshal(confMap)
	if err != nil {
		log.Fatalln(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonString)
}

func getDownloadUrlFromAsset(name string, assets []Assets) string {
	var downloadUrl string
	for i := range assets {
		if assets[i].Name == name+"-Linux-x86_64" {
			fmt.Println(assets[i].BrowserDownloadURL)
			downloadUrl = assets[i].BrowserDownloadURL
			break
		}
	}
	return downloadUrl
}

func DownloadClient(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t updateRequestBody
	err := decoder.Decode(&t)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	_, err = os.Stat(shared.BinaryDir)
	if err != nil {
		os.Mkdir(shared.BinaryDir, 0775)
	}

	clientFolder := shared.BinaryDir + "/" + t.Client + "/"

	_, err = os.Stat(clientFolder)
	if err != nil {
		os.Mkdir(clientFolder, 0775)
	}

	clientFolderWithVersion := shared.BinaryDir + "/" + t.Client + "/"

	_, err = os.Stat(clientFolderWithVersion)
	if err != nil {
		os.Mkdir(clientFolderWithVersion, 0775)
	}

	filePath := clientFolder + t.Version + "/" + t.Client
	fileUrl := t.Url

	err = downloadFile(filePath, fileUrl)
	if err != nil {
		return
	}

	_, err = os.Stat(filePath)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(filePath))
}
