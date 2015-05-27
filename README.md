#Android vibrator for Go

#Description
Port of Android C code (vibrator/vibrator.c) to Golang

#Installation
    go get github.com/MarinX/vibrator
#Notes
* Requires root
* Compile GOOS=linux GOARCH=arm go build

#Example
    package main
    
    import (
	    "fmt"
	    "github.com/MarinX/vibrator"
    )

    func main() {

	    //Check if Android vibrator hardware exists
	    if _, err := vibrator.VibratorExists(); err != nil {
		    fmt.Println(err)
		    return
	    }

	    //Turn on Android vibrator for 1 second
	    if _, err := vibrator.VibratorON(1000); err != nil {
		    fmt.Println(err)
		    return
	    }

	    //Turn off Android vibrator
	    if _, err := vibrator.VibratorOFF(); err != nil {
		    fmt.Println(err)
		    return
	    }
    }

#Execute on Android
###Compile:
    GOOS=linux GOARCH=arm go build
###Push to android
    adb push myprogram /data/local/tmp
###Execute as root
    adb shell
    su
    cd /data/local/tmp
    chmod 777 myprogram
    ./myprogram

#License
This library is under the MIT License
#Author
Marin Basic 
