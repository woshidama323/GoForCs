package main
// This file is: web/main_dev.go
// It will recursively monitor any path we are interested in and re-start the web server
// on any file changes
//
// web/main.go is where my web server code resides. You should change the code
// to your run the command you use for your web server in main() 

import (
    "os/exec"
    "path/filepath"
    "os"
    "crypto/md5"
    "io"
    "encoding/hex"
    "time"
    "fmt"
)

// We will store all MD5 hashes of files we are interested in, in this variable
var fileHashes map[string]string

// Command to start the web server
var cmd *exec.Cmd

// Path we want to monitor for file changes
var pathToMonitor string

// fileMd5 calculates the md5 hash of a file
// Source obtained from http://www.mrwaggel.be/post/generate-md5-hash-of-a-file/
func fileMd5(filePath string) (string, error) {
    //Initialize variable returnMD5String now in case an error has to be returned
    var returnMD5String string

    //Open the passed argument and check for any error
    file, err := os.Open(filePath)
    if err != nil {
        return returnMD5String, err
    }

    //Tell the program to call the following function when the current function returns
    defer file.Close()

    //Open a new hash interface to write to
    hash := md5.New()

    //Copy the file in the hash interface and check for any error
    if _, err := io.Copy(hash, file); err != nil {
        return returnMD5String, err
    }

    //Get the 16 bytes hash
    hashInBytes := hash.Sum(nil)[:16]

    //Convert the bytes to a string
    returnMD5String = hex.EncodeToString(hashInBytes)

    return returnMD5String, nil

}

// fileWatcher monitors files and restarts the web server if any file changes
func fileWatcher() {
    for {
        filepath.Walk(pathToMonitor, func(path string, f os.FileInfo, err error) error {
            fileHash, err := fileMd5(path)
            if err != nil {
                //panic(`Could not calculate hash for ` + path)
            }

            if _, ok := fileHashes[path]; !ok {
                fileHashes[path], _ = fileMd5(path)

            } else if fileHashes[path] != fileHash {
                fileHashes[path] = fileHash

                fmt.Println(`file changed`, path, ` . Restarting web server`)
                cmd.Process.Kill()
                cmd.Run()
            }

            return nil
        })

        time.Sleep(100)
    }
}

func testwatchfile() {
    pathToMonitor = "./"

    // First we get MD5 hashes of all the files we want to monitor
    fileHashes = make(map[string]string)
    filepath.Walk(pathToMonitor, func(path string, f os.FileInfo, err error) error {
        fileHashes[path], _ = fileMd5(path)


        return nil
    })



    // Start a file watcher go rountine that will monitor the files for
    // any changes
    go fileWatcher()

    // Run the web server
    fmt.Println(`Started server`)
    // cmd = exec.Command(`go`, `run`, `web/main.go`)
    // cmd.Run()

    // Create a channel and wait on it. This is here so the main thread
    // exit
    doneChannel := make(chan bool)
    _ = <- doneChannel
}