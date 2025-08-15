# AnimeWorld Downloader

A command-line tool written in Go to search for and download entire anime series from AnimeWorld. It uses the [Gocolly](https://github.com/gocolly/colly) web scraping framework to efficiently navigate the site and retrieve download links.

## ⚠️ Disclaimer

This project is for educational purposes only. Downloading copyrighted material may be illegal in your country. The developer of this project does not condone piracy and is not responsible for any misuse of this tool. 

Web scrapers are fragile and can break if the target website's structure changes. This tool may require updates to remain functional.

## How It Works

The tool follows a multi-step scraping process:

1.  **Search**: Takes the user's query and scrapes the AnimeWorld search results page to find matching anime titles and links.
2.  **Episode Discovery**: After the user selects an anime, the tool visits the anime's main page. It scrapes this page to find two key pieces of information: the total episode count and the link to the main download page.
3.  **Link Generation**: The scraper navigates to the download page to extract a direct download link for a single episode. This link serves as a template. The tool then programmatically generates a full list of direct download links for all episodes by correctly formatting and padding the episode numbers.
4.  **Download**: The program iterates through the generated list of links, downloading each episode as an `.mp4` file into the local `out/` directory.

## Installation

1.  **Clone the repository:**
    ```sh
    git clone https://github.com/alem633/animeworlddl.git
    cd animeworlddl
    ```

2.  **Install dependencies:**
    The project uses Go Modules. The dependencies (like Gocolly) will be downloaded automatically when you build or run the project. You can also fetch them manually:
    ```sh
    go mod tidy
    ```

3.  **Build the executable:**
    ```sh
    go build .
    ```
    This will create an executable file named `animeworld-dl` (or `animeworld-dl.exe` on Windows).

## Usage

Run the program from your terminal, providing the name of the anime you want to download as an argument. Make sure to enclose multi-word titles in quotes.

```sh
./animeworld-dl "search query"
```

### Example

```sh
# 1. Run the command with your search query
$ ./animeworld-dl "Jujutsu Kaisen"

# 2. The program presents the search results
[0]: Jujutsu Kaisen (TV)
[1]: Jujutsu Kaisen 0 (Movie)

# 3. Enter the number corresponding to your choice
input: 0

# 4. The program scrapes for download info and begins downloading
	[epCount]: 24
	[downloadPageLink]: https://www.animeworld.tv/play/jujutsu-kaisen.FLk2u/AYb-C8
	[mainDlLink]: https://download.animeworld.tv/files/Jujutsu Kaisen/Ep_01_SUB_ITA_1080p_BD.mp4
[DL]: https://download.animeworld.tv/files/Jujutsu Kaisen/Ep_01_SUB_ITA_1080p_BD.mp4
...downloading episode 1...
[DL]: https://download.animeworld.tv/files/Jujutsu Kaisen/Ep_02_SUB_ITA_1080p_BD.mp4
...downloading episode 2...
... and so on ...
```

## To-Do / Future Improvements

-   [ ] Add support for concurrent downloads to speed up the process.
-   [ ] Implement a progress bar for downloads.
-   [ ] Allow users to select a specific range of episodes to download (e.g., episodes 5-10).
-   [ ] Add functionality to resume interrupted downloads.
-   [ ] Improve error handling for network issues or changes in website structure.
