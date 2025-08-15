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

const outputDir = "out/"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("missing search query")
		os.Exit(1)
	}
	makeOutputDir()
	query := os.Args[1]

	searchResult := []scraper.Anime{}
	searchResult = scraper.SearchAnime(query)

	for i, v := range searchResult {
		fmt.Printf("[%d]: %s\n", i, v.Name)
	}

	fmt.Print("input: ")
	var userSelection int
	_, err := fmt.Scan(&userSelection)
	if err != nil || userSelection < 0 || userSelection >= len(searchResult) {
		fmt.Println("Invalid input")
		os.Exit(1)
	}

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
		filepath := outputDir + filename + "_" + strconv.Itoa(i+1) + ".mp4"
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

func makeOutputDir() {
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		fmt.Printf("Directory '%s' not found, running mkdir.\n", outputDir)
		err := os.Mkdir(outputDir, 0755)
		if err != nil {
			fmt.Printf("Error during os.Mkdir: %s\n", err)
			return
		}
		fmt.Printf("Directory '%s' successfully made.\n", outputDir)
	} else {
		fmt.Printf("Directory '%s' found.\n", outputDir)
	}
}
