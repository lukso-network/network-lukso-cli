export interface AvailableSoftwareBackendResponse {
  [key: string]: ReleaseBackendResponse;
}

export interface ReleaseBackendResponse {
  name: string;
  humanReadableName: string;
  downloadInfo: DownloadInfoPerTag;
}

export interface DownloadInfoPerTag {
  [key: string]: DownloadInfoBackendResponse;
}

export interface DownloadInfoBackendResponse {
  tag: string;
  downloadUrl: string;
}

export interface Releases {
  name: string;
  humanReadableName: string;
  downloadInfo: DownloadInfo[];
}

export interface DownloadInfo {
  tag: string;
  name: string;
  downloadUrl: string;
  isDownloading?: boolean;
}
