@echo off

git add .

:enter_message
set /p "comment=Enter Commit Message: "

echo Commit With Message: "%comment%" ?
CHOICE /C YRC /M "Press Y for Yes, R for Redo or C for Cancel."
if %ERRORLEVEL% == 2 goto enter_message
if %ERRORLEVEL% == 3 exit

git commit -m "%comment%"

git remote add origin https://github.com/overrnet/yamlconfig >nul

git push --force -u origin master >nul

::for /f "tokens=* delims=" %a in ('git tag -l') do git tag -d %a

echo current version is
git fetch
git describe --tags

:enter_version
set /p "version=Enter New Version: "

echo Add Tag Version: "%version%" ?
CHOICE /C YRC /M "Press Y for Yes, R for Redo or C for Cancel."

if %ERRORLEVEL% == 2 goto enter_version
if %ERRORLEVEL% == 3 exit

git tag %version%

git push origin --tags

echo.

pause