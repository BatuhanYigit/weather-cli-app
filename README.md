# Weather CLI Application

This is a simple command-line application written in Go programming language to fetch and display weather information.

## Usage

You can run the application in the terminal with the following command:

```shell

weather [City Name]

```

By default, weather information is fetched for the city "Ankara," but you can specify a different city name if desired.

##Dependencies
1. This application uses the following dependencies:
 github.com/fatih/color: Used for colorful text output.

#Installation
1. Clone this project:

```bash
git clone [project_url]
```
2. Navigate to the project directory:
```bash
cd [project_directory]
```
3. To build the application:

```bash
go build -o weather
```

4. Run the application:

```bash
weather [City Name]
```
##Example Output
The application provides current weather information for the specified city and forecasts for the upcoming hours.

Ankara, Turkey: 15°C, Partly cloudy
10:00 - 17°C, 10%, Clear
11:00 - 19°C, 5%, Partly cloudy
12:00 - 22°C, 0%, Sunny

