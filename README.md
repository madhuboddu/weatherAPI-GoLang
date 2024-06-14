# Weather API GoLang

This project is a GoLang application that fetches weather data for a given Canadian location using a weather API.



## Project Structure

- `main/`: Contains the main Go source code files.
- `app.env`: Environment configuration file.
- `.vscode/`: VSCode configuration files.
- `.github/` and `github/`: GitHub workflows and other configurations.
- `.git/`: Git repository metadata.
- `go.mod`: Go module file.
- `go.sum`: Go dependencies file.

## Prerequisites

- Go 1.16+ installed
- Docker (optional, for containerized deployment)

## Setup

1. **Clone the repository:**

   ```sh
   git clone https://github.com/madhuboddu/weatherAPI-GoLang.git
   cd weatherAPI-GoLang


2. **Install dependencies:**
    Assuming you have GO installed, run : go mod tidy


3. **Configure environment variables:**

    Edit the app.env file to set your configuration variables:
    
    BaseURL=http://api.weatherapi.com/v1/current.json
    APIKey=a489e9b203534843ba000645242604

## Building and Running

Run the application:
    go run main.go


### Scope of improvements

1. Planing to build a UI for this weather API.
2. Expand the zipcode support to get weather accross the world.


### Authors

Madhusudhan Boddu - Initial work
Manohar Reddy     - Reviewer

