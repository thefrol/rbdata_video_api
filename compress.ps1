get-childitem -Path . -Exclude .*,*.zip -Recurse | Compress-Archive -DestinationPath archive.zip -Force