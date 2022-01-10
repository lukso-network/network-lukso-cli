package downloader

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"lukso/apps/lukso-manager/shared"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
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
	fmt.Println("Downloading: ", url)

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

type release struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type downloadInfo struct {
	Tag         string `json:"tag"`
	DownloadURL string `json:"downloadUrl"`
}

type clientInfo struct {
	Name              string                  `json:"name"`
	HumanReadableName string                  `json:"humanReadableName"`
	DownloadInfo      map[string]downloadInfo `json:"downloadInfo"`
}

var ReleaseLocations = map[string]release{
	"pandora": {
		Name: "Pandora",
		URL:  "https://api.github.com/repos/lukso-network/pandora-execution-engine/releases",
	},
	"vanguard": {
		Name: "Vanguard",
		URL:  "https://api.github.com/repos/lukso-network/vanguard-consensus-engine/releases",
	},
	"lukso-orchestrator": {
		Name: "Orchestrator",
		URL:  "https://api.github.com/repos/lukso-network/lukso-orchestrator/releases",
	},
	"lukso-deposit-cli": {
		Name: "Deposit CLI",
		URL:  "https://api.github.com/repos/lukso-network/network-deposit-cli/releases",
	},
	"lukso-validator": {
		Name: "Validator",
		URL:  "https://api.github.com/repos/lukso-network/vanguard-consensus-engine/releases",
	},
}

type updateRequestBody struct {
	Client  string `json:"client"`
	Version string `json:"version"`
	Url     string `json:"url"`
}

func GetDownloadedVersions() (map[string][]string, error) {
	var DownloadedVerions = map[string][]string{}

	downloads := shared.BinaryDir

	_, err1 := os.Stat(downloads)
	if err1 != nil {
		os.MkdirAll(downloads, 0775)
	}

	err := filepath.Walk(downloads,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				pathParts := strings.Split(path, "binaries")
				fileNameParts := strings.Split(pathParts[1], "/")
				DownloadedVerions[fileNameParts[1]] = append(DownloadedVerions[fileNameParts[1]], fileNameParts[2])
			}

			return nil
		})

	if err != nil {
		return DownloadedVerions, errors.New("file system error")
	}

	return DownloadedVerions, nil
}

func GetDownloadedVersionsEndpoint(w http.ResponseWriter, r *http.Request) {

	DownloadedVerions, err := GetDownloadedVersions()

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(nil)
		return
	}

	jsonString, err := json.Marshal(DownloadedVerions)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonString)
}

func GetAvailableVersions() (map[string]clientInfo, error) {
	confMap := map[string]clientInfo{}
	for client, release := range ReleaseLocations {
		r, err := http.Get(release.URL + "?per_page=20")
		if err != nil {
			fmt.Println("Request to "+release.URL+" failed.", err)
			return confMap, errors.New("Request Failed")
		}

		confMap[client] = clientInfo{
			Name:              client,
			HumanReadableName: release.Name,
			DownloadInfo:      make(map[string]downloadInfo),
		}

		if r.StatusCode == http.StatusForbidden {
			return confMap, errors.New("Github API Rate Limit Exceeded")
		}

		if r.StatusCode != http.StatusOK {
			log.Fatal(err.Error())
			return confMap, errors.New("HTTP Request failed")
		}

		decoder := json.NewDecoder(r.Body)
		var releases GithubReleases

		decodeError := decoder.Decode(&releases)
		if decodeError != nil {
			log.Fatalln(decodeError)
			log.Fatal(decodeError.Error())
			return confMap, errors.New("Cannot decode JSON API")
		}

		for _, v := range releases {
			assetURL := getDownloadUrlFromAsset(client, v.Assets)
			if assetURL != "" {
				confMap[client].DownloadInfo[v.TagName] = downloadInfo{
					Tag:         v.TagName,
					DownloadURL: assetURL,
				}
			}
		}

	}
	return confMap, nil
}

func GetAvailableVersionsEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	confMap, err := GetAvailableVersions()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	jsonString, err := json.Marshal(confMap)

	if err != nil {
		log.Fatalln(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("invalid json"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonString)
}

func getDownloadUrlFromAsset(name string, assets []Assets) string {
	var downloadUrl string
	os_type := strings.Title(runtime.GOOS)

	for i := range assets {
		if assets[i].Name == name+"-"+os_type+"-x86_64" {
			downloadUrl = assets[i].BrowserDownloadURL
			break
		}
	}
	return downloadUrl
}

func DownloadClient(client string, version string) {
	clientFolder := shared.BinaryDir + client + "/"
	clientFolderWithVersion := clientFolder + version

	createDirIfNotExists(shared.BinaryDir)
	createDirIfNotExists(clientFolder)
	createDirIfNotExists(clientFolderWithVersion)

	availableVersions, err := GetAvailableVersions()

	filePath := clientFolder + version + "/" + client
	fileUrl := availableVersions[client].DownloadInfo[version].DownloadURL

	err = downloadFile(filePath, fileUrl)
	if err != nil {
		log.Fatal("File " + fileUrl + " could not be downloaded: " + err.Error())
		return
	}

	_, err = os.Stat(filePath)
	if err != nil {
		log.Fatal([]byte("File " + fileUrl + " not found"))
		return
	}

	os.Chmod(filePath, 0775)
}

func DownloadClientEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var t updateRequestBody
	err := decoder.Decode(&t)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	clientFolder := shared.BinaryDir + t.Client + "/"
	clientFolderWithVersion := clientFolder + t.Version

	createDirIfNotExists(shared.BinaryDir)
	createDirIfNotExists(clientFolder)
	createDirIfNotExists(clientFolderWithVersion)

	filePath := clientFolder + t.Version + "/" + t.Client
	fileUrl := t.Url

	err = downloadFile(filePath, fileUrl)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("File " + fileUrl + " could not be downloaded: " + err.Error()))
		return
	}

	_, err = os.Stat(filePath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("File " + fileUrl + " not found"))
		return
	}

	os.Chmod(filePath, 0775)

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode("Download Successful"); err != nil {
		panic(err)
	}
}

func DownloadClientBinary(client string, tag_version string, url string) {
	clientFolder := shared.BinaryDir + client + "/"
	clientFolderWithVersion := clientFolder + tag_version

	createDirIfNotExists(shared.BinaryDir)
	createDirIfNotExists(clientFolder)
	createDirIfNotExists(clientFolderWithVersion)

	filePath := clientFolder + tag_version + "/" + client
	urlWithTagSet := strings.Replace(url, "_TAG_", tag_version, 1)
	fileUrl := strings.Replace(urlWithTagSet, "_OS_TYPE_", strings.Title(runtime.GOOS)+"-x86_64", 1)

	err := downloadFile(filePath, fileUrl)
	if err != nil {
		log.Fatal(err.Error())
		log.Fatal("Failed to download" + fileUrl)
		return
	}

	_, err = os.Stat(filePath)
	if err != nil {
		return
	}

	os.Chmod(filePath, 0775)
}

func createDirIfNotExists(folder string) {
	_, err := os.Stat(folder)
	if err != nil {
		os.MkdirAll(folder, 0775)
	}
}

func DownloadConfigFiles(network string) (err error) {
	CDN := "https://storage.googleapis.com/l15-cdn/networks/" + network
	folder := shared.NetworkDir + network + "/config"
	createDirIfNotExists(folder)

	dlError := downloadFile(folder+"/network-config.yaml", CDN+"/network-config.yaml?ignoreCache=1")
	if dlError != nil {
		return
	}
	dlError1 := downloadFile(folder+"/pandora-genesis.json", CDN+"/pandora-genesis.json?ignoreCache=1")
	if dlError1 != nil {
		return
	}
	dlError2 := downloadFile(folder+"/vanguard-genesis.ssz", CDN+"/vanguard-genesis.ssz?ignoreCache=1")
	if dlError2 != nil {
		return
	}
	dlError3 := downloadFile(folder+"/vanguard-config.yaml", CDN+"/vanguard-config.yaml?ignoreCache=1")
	if dlError3 != nil {
		return
	}
	dlError4 := downloadFile(folder+"/pandora-nodes.json", CDN+"/pandora-nodes.json?ignoreCache=1")
	if dlError4 != nil {
		return
	}
	return
}
