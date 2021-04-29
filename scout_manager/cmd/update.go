package cmd

import (
	"errors"
	"fmt"
	"github.com/kurolz/cobra"
	"github.com/toolkits/file"
	"io"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"../process"
)

var Update = &cobra.Command{
	Use:   "update",
	Short: "Update Service",
	RunE:  update,
	//Args:          cobra.MinimumNArgs(2),
	SilenceUsage:  true,
	SilenceErrors: true,
}

func update(c *cobra.Command, args []string) error {
	currentUser, err := user.Current()
	if err != nil {
		if err != nil {
			fmt.Printf("get user home dir failed. ")
			return err
		}
	}
	if runtime.GOOS == "linux" {
		tarName := file.Basename(args[0])
		tarFile := path.Join(currentUser.HomeDir, tarName)
		tarDirName := strings.Split(tarName, ".tar.gz")[0]

		if _, err := os.Stat(tarFile); !os.IsNotExist(err) {
			err = os.Remove(tarFile)
			if err != nil {
				fmt.Printf("remove file %s failed. ", tarFile)
				return err
			}
		}

		cmd := exec.Command("wget", args[0])
		cmd.Dir = currentUser.HomeDir
		err = cmd.Run()
		if err != nil {
			fmt.Printf("wget tar.gz file failed. ")
			return err
		}

		err = os.Chmod(tarFile, 0700)
		if err != nil {
			fmt.Printf("chmod tar file failed. ")
			return err
		}

		cmdTar := exec.Command("tar", "-zxvf", tarName)
		cmdTar.Dir = currentUser.HomeDir
		err = cmdTar.Run()
		if err != nil {
			fmt.Printf("tar unzip file failed. ")
			return err
		}

		execFile := path.Join(currentUser.HomeDir, tarDirName, "bin", "troop-scout")
		if _, err := os.Stat(execFile); os.IsNotExist(err) {
			return errors.New("exec file not exists in tar.gz")
		}

		cmdCp := exec.Command("cp", "-f", execFile, path.Join(args[1], "troop-scout"))
		err = cmdCp.Run()
		if err != nil {
			fmt.Printf("cp exec file failed. ")
			return err
		}

		if err := ExecRestart(); err != nil {
			fmt.Printf("restart service failed. ")
			return err
		}
	} else {
		req, _ := http.NewRequest("GET", args[0], nil)

		res, err := http.DefaultClient.Do(req)

		if err != nil {
			fmt.Print(err.Error())
			return err
		}

		if res.StatusCode == 401 {
			return errors.New("file exec file failed: invalid token")
		}

		fileP := path.Join(currentUser.HomeDir, file.Basename(args[0]))

		f, err := os.Create(fileP)

		if err != nil {
			fmt.Print(err.Error())
			f.Close()
			return err
		}
		_, err = io.Copy(f, res.Body)
		if err != nil {
			fmt.Print(err.Error())
			f.Close()
			return err
		}
		f.Close()

		//errRename := os.Rename(filepath.FromSlash(filepath.Join(args[1], "troop-scout.exe")), `troop-scout_wait_delete.exe`)
		//if errRename != nil {
		//	fmt.Print(errRename.Error())
		//	return errRename
		//}
		fileP = filepath.FromSlash(fileP)
		copyCmd := fmt.Sprintf("copy /Y %s %s", fileP, filepath.FromSlash(filepath.Join(args[1], "troop-scout.exe")))
		cmdL := fmt.Sprintln("cmd /c net stop troop-scout &&", copyCmd, "&& net start troop-scout")
		err = process.StartProcessAsCurrentUser(
			"C:\\Windows\\system32\\cmd.exe",
			cmdL,
			currentUser.HomeDir,
			true)
		if err != nil {
			return err
		}
		fmt.Print("successfully")
		//cmdL := exec.Command("cmd", "/c", "net stop troop-scout && " + copyCmd + " && net start troop-scout")
		//cmdL.Stdout = os.Stdout
		//cmdL.Stderr = os.Stderr

		//err = cmdL.Start()
		//if err != nil {
		//	return err
		//}

		//copyCmd := fmt.Sprintln("copy /Y", fileP, filepath.FromSlash(filepath.Join(args[1], "troop-scout.exe")))
		//log.Print(copyCmd)
		//cmdCp := exec.Command("cmd", "/c", copyCmd)
		//errCp := cmdCp.Run()
		//if errCp != nil {
		//	fmt.Printf("copy exec file failed. ")
		//	return errCp
		//}

		//err = utils.StartProcessAsCurrentUser(
		//	"C:\\Windows\\system32\\cmd.exe",
		//	"cmd /c net stop troop-scout && net start troop-scout",
		//	currentUser.HomeDir,
		//	true)
		//
		//if err != nil {
		//	return err
		//}

	}
	return nil
}
