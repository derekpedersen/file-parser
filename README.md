# File Parser

This is my `File Parser` application that I wrote as part of the [interview process](https://github.com/serif-health/takehome) for [Serif Health](https://www.serifhealth.com/).

## Sample Output

If you want to skip the `build` & `run` phases and just view the output I generated in my latest run, you can view them in the [samples folder](./bin/samples).

## Run

As part of the application submission I've included a prebuilt version of the library. Typically I would not commit any build artifacts to the repositoy, but in this case the tradeoff between available time and setting up a distribution and consumption mechanism outside of the repository made this the best option.

| OS | File |
| --- | --- |
| Linux | `./bin/dist/windows/file-parser` |
| Mac | `./bin/dist/mac/file-parser` |
| Windows | `./bin/dist/windows/file-parser.exe` |

When the command is run from the root of this repository the generated CSVs are placed into the [bin/output](./bin/output/) folder.

## Build

This project was built in [golang](https://go.dev/doc/install) and utilizes a [Makefile](https://www.gnu.org/software/make/manual/make.html) to help simplify the build process.

If you have `golang` installed then just run `make run`.

## Time Break Down

I did spend more time than the project initially asked for and recommended, but not by much and I felt it was a worthy endeavor to help highlight my abilities so have provided the following breakdown. Each of the time quotes includes development/testing/bug-fixing.

| Task | Time | Notes |
|---|---|---|
| `models` | 5 - 10 minutes | Simple DTOs so not much time spent on them besides just looking at the available data and creating associated fields. |
| `utilities` | 5 minutes | Go does require that you need to write some boiler plate but as you can see it's not that complicated and is easily copy-pasta'd as needed. |
| `main` | 10 - 15 minutes | The `main` file was relatively quick to setup. The only major change I made from it's inception to it's delivery was how it handled the iterative processing of the data sources. Initially had it processing the files one at a time then realised I needed to process them both first in order cross polinate the sources and enrich the data before saving the CSVs. |
| `domain` | 65 - 70 minutes | Here is the meat of the application. |
| `Makefile` | 5 minutes | A small `Makefile` to help make running the application easier. |
| `README` | 10 - 15 minutes | Writting, editing, and re-reading. |
| `Total` | 100 - 125 minutes | Overall time estimate for this application. |

## Tradeoffs (Future Improvements)

With every application there is always a trade off between how much time we have and what we can deliver. Given infinite time we could deliver everything. But we don't have infinite time and we don't necessarily need everything to get started. Here are the tradeoffs I made when writting this application. 

### No Tests

Ideally now that I got the application working as intended I would write some tests to try ensure that we do not break the currently supported features when making modifications in the future. It's also important to take into account that the application is early in it's life at this point so we don't want to build out too much testing for it where we end up paiting ourselves into a corner so to speak and it becomes extremely tedious, or even impossible, to modify the application for a related but different use case.  

### Long Term Storage

At the moment I have it just saving the json/file that was returned when calling the URL. To increase the robustness of the application I would like to incorporate storing the raw and processed data in a Relational Database to make the results easily distributable as well as queryable.

### File Processing Iteration

This application needs a way to dynamically process a new file without rebuilding the application. Currently the URL hosting the json/files is hard coded as a variable, this can be improved by moving to a mechanism in which the URLS are supplied to the application at run time.

### Logging

There is some logging but it's mostly just around catching errors for development debugging purposes. Would not be a large effort to add more INFO/DEBUG logging throughout the application. 

### Validation

Knowing (like 98% sure) there are centralized sources upon which we can get a list of legitimate `Procedure Codes` I would like to add validation around the code being supplied is a valid one rather than relying on it just being present. 

### Deployment

Since this is a batch process that doesn't rely on instaneous information I assume this could be configured to be run upon a once daily schedule using a varierty of different scheduling tools available. This run rate will cut down on the costs required as we don't need to be constantly running the application and reaching out for new information. 

### API

Once we have this running as the ETL process we can stand up a seperate API to expose this data to the UI (or customers directly via API).

