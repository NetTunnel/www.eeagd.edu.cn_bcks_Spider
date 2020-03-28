@echo off
SET CGO_ENABLED=0
SET GOARCH=386
SET GOOS=windows
SET DIST_FILE_NAME=Spider.exe
goto :do_build

Rem for linux.
:os_linux
  SET GOOS=linux
  SET DIST_FILE_NAME=Spider


Rem do build task.
:do_build
  echo build for %GOOS%^<%GOARCH%^>
  go build -ldflags "-s -w" -o %DIST_FILE_NAME%
  if %ERRORLEVEL% NEQ 0 (
    pause
    exit
  )
  if "%GOOS%" == "windows" goto :os_linux
  echo build complete
  pause