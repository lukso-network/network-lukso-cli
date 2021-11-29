# This file is used for converting "--" flag prefix to "-"
# before being processed by the main script
# since "--" cannot be used for flags in Powershell

$InstallDir = $Env:APPDATA + "\LUKSO"

$cmd = "$args"
$cmd = $cmd.Replace(" --", " -")

Write-Output $cmd
powershell -command "$InstallDir\lukso.ps1 $cmd"