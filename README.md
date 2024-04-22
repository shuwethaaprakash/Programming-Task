<!-- Improved compatibility of back to top link: See: https://github.com/othneildrew/Best-README-Template/pull/73 -->
<a name="readme-top"></a>

<!-- PROJECT LOGO -->
<br />

<h3 align="center">Log File Parser</h3>

  <p align="center">
    Parse a log file containing HTTP requests and report on its contents.
    <br />
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#testing">Testing</a></li>
    <li><a href="#assumptions">Assumptions</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project
Parse a log file containing HTTP requests and report on its contents. 

For a given file we want to know:
- The number of unique IP addresses
- The top 3 most visted URLs
- The top 3 most active IP addresses


### Built With

* Golang


<!-- GETTING STARTED -->
## Getting Started

This is how to set up the project locally on your device. 

### Prerequisites

Must have most recent stable version of golang _(at this current time: go1.22.1)_ installed on local device.
* Golang - please refer to the [documentation](https://go.dev/doc/install) online.

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/shuwethaaprakash/Programming-Task.git
   ```


<!-- USAGE EXAMPLES -->
## Usage

Run the program with the following command:
```sh
   go run cmd/main/main.go <log_file>
```
Or you can run it from the ```cmd``` or ```main``` directory and adjust the location of the ```main.go``` file accordingly.

The output should look as follows (if used with ```demo/programming-task-example-data.log```):

![output](/demo/output.png)

<!-- TESTING -->
## Testing

The program has separate tests for each file within the internal package. You can test the program with the following command:
```sh
   go test internal/test/<package>_test.go
```
where package is, for example, `validate` or `count` etc.

_Or_ you can change location to within a given test directory (such as `internal/file/test`) and run :
```sh
   go test <package>_test.go
``` 

<!-- ASSUMPTIONS -->
## Assumptions

- The program is only required to take in one log file at a time
- All log files have the same structure for their requests </br>
  i.e. `[IP ADDRESS] - [USER] [TIMESTAMP] "GET [REQUEST_URL] HTTP/1.1" [HTTP STATUS] 3574 "-" "[SEARCH ENGINES]"`
- Any URLs that cause HTTP errors are still included
- The log file provided is in a readable format

<!-- CONTACT -->
## Contact

Shuwethaa Prakash

Project Link: [https://github.com/shuwethaaprakash/Programming-Task](https://github.com/shuwethaaprakash/Programming-Task)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
