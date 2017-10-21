package helper

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"github.com/the-big-three/synchrogit/models"
	"io/ioutil"
	"strconv"
	"sync"
)

/**
opens and parses Config
*/
func ParseConfig() (models.SynchroGitSettings, error) {

	config := os.Getenv("CONFIG");
	synchroGitConfigFile, err := os.Open(config)
	settings := new(models.SynchroGitSettings)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while opening synchroGitSync.json : %s\n", err.Error())
		os.Exit(1)
	}

	err = json.NewDecoder(synchroGitConfigFile).Decode(&settings)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while parsing synchroGitSync.json : %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println("Successfully imported synchroGitSettings from synchroGitSync.json")
	fmt.Printf("Detected %d repos to sync\n", len(settings.Syncs))
	synchroGitConfigFile.Close()
	return *settings, err
}

func PullMirror(url string, targetDir string) {
	fmt.Printf("Cloning from %s into %s\n", url, targetDir)
	cmdName := "git"
	cmdArgs := []string{"clone", "--mirror", url, targetDir}
	cmd := exec.Command(cmdName, cmdArgs...)
	cmdOut, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "There was an error running git clone command, error %s, stout %s\n", err, string(cmdOut))
	}
}

func PushMirror(url string) {
	fmt.Printf("Pushing to %s\n", url)
	cmdName := "git"
	cmdArgs := []string{"push", "--mirror", url}
	cmd := exec.Command(cmdName, cmdArgs...)
	cmdOut, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "There was an error running git push command, error %s, stout %s\n", err, string(cmdOut))
	}
}

func ChangeDir(dir string) {
	fmt.Printf("Changing to directory %s\n", dir)
	err := os.Chdir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error changing to directory: %s, exiting\n", dir)
		os.Exit(-1)
	}
}

func SyncRepo(index int, element models.SynchroGitSync, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("\n------ Start mirroring %s to %s -------\n", element.Source.Url, element.Target.Url)
	// index is the index where we are
	// element is the element from synchroGitSettings.Syncs for where we are
	tmpGitRepoDir, err := ioutil.TempDir("", strconv.Itoa(index))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating tmp directory for repo: %s, exiting for error %s\n", element.Source.Url, err)
		os.Exit(-1)
	}
	PullMirror(element.Source.Url, tmpGitRepoDir)
	ChangeDir(tmpGitRepoDir)
	PushMirror(element.Target.Url)
	fmt.Printf("\n------ Finished mirroring %s to %s -------\n", element.Source.Url, element.Target.Url)
}

func CheckEnv(env string){
	config := os.Getenv("CONFIG");
	if(len(config) == 0){
		fmt.Println("env CONFIG not set, closing")
		os.Exit(1)
	}
}
