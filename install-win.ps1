$Network = "l15";
$InstallDir = $Env:APPDATA+"\LUKSO";

Function download ($url, $dst) {
    $client = New-Object System.Net.WebClient
    $client.DownloadFile($url, $dst)
}

Function download_network_config ($network) {
    $CDN = "https://storage.googleapis.com/l15-cdn/networks/"+$network
    $TARGET = $InstallDir+"\networks\"+$network+"\config"
    New-Item -ItemType Directory -Force -Path $TARGET
    download $CDN"/network-config.yaml?ignoreCache=1" $TARGET"\network-config.yaml"
    download $CDN"/network-config.yaml?ignoreCache=1" $TARGET"\network-config.yaml"
}


New-Item -ItemType Directory -Force -Path $InstallDir
New-Item -ItemType Directory -Force -Path $InstallDir/tmp
New-Item -ItemType Directory -Force -Path $InstallDir/binaries
New-Item -ItemType Directory -Force -Path $InstallDir/networks
New-Item -ItemType Directory -Force -Path $InstallDir/globalPath

download_network_config("l15-prod")
download_network_config("l15-staging")
download_network_config("l15-dev")

Invoke-WebRequest -Uri http://192.168.0.112:8080/script -OutFile $InstallDir\lukso.ps1
if (Test-Path "$InstallDir\globalPath\lukso") {
    rm "$InstallDir\globalPath\lukso"
}

Write-Output "powershell.exe -File $InstallDir\lukso.ps1 %*" | Out-File -Encoding ASCII -FilePath "$InstallDir\globalPath\lukso.bat"

lukso bind-binaries `
-orchestrator v0.1.0-rc.1 `
-pandora v0.1.0-rc.1 `
-vanguard v0.1.0-rc.1 `
-validator v0.1.0-rc.1 `
-eth2stats v0.1.0-develop
