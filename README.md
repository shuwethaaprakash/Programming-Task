<!-- Improved compatibility of back to top link: See: https://github.com/othneildrew/Best-README-Template/pull/73 -->
<a name="readme-top"></a>

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/shuwethaaprakash/Programming-Task">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a>

<h3 align="center">Log File Parser</h3>

  <p align="center">
    Parse a log file containing HTTP requests and report on its contents. Namely:
    * The number of unique IP addresses
    * The top 3 most visted URLs
    * The top 3 most active IP addresses
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
    <li><a href="#assumptions">Assumptions</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project
Parse a log file containing HTTP requests and report on its contents. For a given file we want to know:
    * The number of unique IP addresses
    * The top 3 most visted URLs
    * The top 3 most active IP addresses

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

* Golang

<p align="right">(<a href="#readme-top">back to top</a>)</p>



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

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

Run the program with the following command:
```sh
   go run main.go <log_file>
```

The output should look as follows:

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- ASSUMPTIONS -->
## Assumptions

- The program is only required to take in one log file at a time
- All log files have the same structure for their requests i.e.
  `[IP ADDRESS] - [USER] [TIMESTAMP] "GET [REQUEST_URL] HTTP/1.1" [HTTP STATUS] 3574 "-" "[SEARCH ENGINES]"`
- The requests made are GET requests
- Any URLs that cause HTTP errors are still included
- The log file provided is in a readable format

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->
## Contact

Shuwethaa Prakash
Project Link: [https://github.com/shuwethaaprakash/Programming-Task](https://github.com/shuwethaaprakash/Programming-Task)

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/shuwethaaprakash/Programming-Task.svg?style=for-the-badge
[contributors-url]: https://github.com/shuwethaaprakash/Programming-Task/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/shuwethaaprakash/Programming-Task.svg?style=for-the-badge
[forks-url]: https://github.com/shuwethaaprakash/Programming-Task/network/members
[stars-shield]: https://img.shields.io/github/stars/shuwethaaprakash/Programming-Task.svg?style=for-the-badge
[stars-url]: https://github.com/shuwethaaprakash/Programming-Task/stargazers
[issues-shield]: https://img.shields.io/github/issues/shuwethaaprakash/Programming-Task.svg?style=for-the-badge
[issues-url]: https://github.com/shuwethaaprakash/Programming-Task/issues
[license-shield]: https://img.shields.io/github/license/shuwethaaprakash/Programming-Task.svg?style=for-the-badge
[license-url]: https://github.com/shuwethaaprakash/Programming-Task/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/linkedin_username
[product-screenshot]: images/screenshot.png
