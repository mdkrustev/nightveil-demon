hdiutil create -volname "NightVeilDemon" \
  -srcfolder release \
  -ov -format UDZO \
  NightVeilDemon.dmg


MAC

go build -o NightVeilDemon


hdiutil create -volname "NightVeilDemon" \
  -srcfolder release \
  -ov -format UDZO \
  NightVeilDemon.dmg


go build -o NightVeilDemon


go run main.go


Windows

go build -ldflags="-H windowsgui" -o NightVeilDemon.exe