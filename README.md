# Weather CLI Application

This is a simple command-line application written in Go programming language to fetch and display weather information.

## Usage

You can run the application in the terminal with the following command:

```shell
weather [City Name]
By default, weather information is fetched for the city "Ankara," but you can specify a different city name if desired.

Dependencies
This application uses the following dependencies:

github.com/fatih/color: Used for colorful text output.
Installation
Clone this project:

shell
Copy code
git clone [project_url]
Navigate to the project directory:

shell
Copy code
cd [project_directory]
To build the application:

shell
Copy code
go build -o weather
Run the application:

shell
Copy code
weather [City Name]
Example Output
The application provides current weather information for the specified city and forecasts for the upcoming hours.

Example output:

mathematica
Copy code
Ankara, Turkey: 15°C, Partly cloudy
10:00 - 17°C, 10%, Clear
11:00 - 19°C, 5%, Partly cloudy
12:00 - 22°C, 0%, Sunny
License
This project is licensed under the MIT License - see the LICENSE file for details.
