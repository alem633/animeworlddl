package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"io"
	"net/http"

	"scraper/animeworld_dl/scraper"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	query := os.Args[1]

	searchResult := []scraper.Anime{}
	searchResult = scraper.SearchAnime(query)

	for i, v := range searchResult {
		fmt.Printf("[%d]: %s\n", i, v.Name)
	}

	fmt.Print("input: ")
	var userSelection int
	fmt.Scan(&userSelection)

	animeMainPageLink := searchResult[userSelection].Link
	epCount, downloadPageLink := scraper.GetDownloadInfo(animeMainPageLink)
	mainDlLink := scraper.GetDirectDlLink(downloadPageLink)

	fmt.Printf("\t[epCount]: %d\n", epCount)
	fmt.Printf("\t[downloadPageLink]: %s\n", downloadPageLink)
	fmt.Printf("\t[mainDlLink]: %s\n", mainDlLink)

	directDlLinks := getDirectDlLinksSlice(epCount, mainDlLink)

	filename := strings.Replace(searchResult[userSelection].Name, " ", "_", -1) 
	for i, link := range directDlLinks {
		fmt.Printf("[DL]: %s\n", link) 
		filepath := "out/" + filename + "_" + strconv.Itoa(i+1) + ".mp4"
		err := DownloadFile(filepath, link)
		if err != nil {
			fmt.Println("invalid")
			os.Exit(1)
		}
	}
}

func DownloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func getDirectDlLinksSlice(epCount int, mainDlLink string) (directDlLinks []string) {
	for i:=1;i<=epCount;i++ {
		formattedEpisodeDlLink := getFormattedEpisodeDlLink(i, epCount, mainDlLink)
		directDlLinks = append(directDlLinks, formattedEpisodeDlLink)
	}
	return directDlLinks
}

func getFormattedEpisodeDlLink(epNumber int, epCount int, mainDlLink string) (formattedEpisodeDlLink string) {
	var epNumber_f string  = epNumberFormatter(epNumber, epCount)
	ss := strings.Split(mainDlLink, "_Ep_")
	ss2 := strings.Split(ss[1], "_")

	prefix := ss[0] + "_Ep_"
	var suffix string
	for i:=1;i<len(ss2);i++ {
		suffix += "_" + ss2[i]
	}

	formattedEpisodeDlLink = prefix + epNumber_f + suffix
	return formattedEpisodeDlLink
}

func epNumberFormatter(epNumber int, epCount int) string {
	paddingWidth := len(strconv.Itoa(epCount))

	format := fmt.Sprintf("%%0%dd", paddingWidth)

	return fmt.Sprintf(format, epNumber)
}
