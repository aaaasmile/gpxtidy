# GpxTidy
Small tool to change gpx files. Created for removing the extension section and Metadata from a gpx file.

## Motivation
Importing gpx file with extensions and Metadata from Garmin Desktop App into the TomTom planning tool failed.
For example the file .\data\S1_12Ipertrail_2022-v221227.GPX.
So I created this tool to remove all extension sections and also changing metadata and gpx attributes to see if it works.
I do not known exactly what is wrong for TomTom because the importer doesn't shows any erros.
Simply does nothing. I suspect the Metadata  section or the extenstion section aren't recognized in TomTom.
To make the import working I have used a gpx section from a TomTom activity export (see .\data\reference_tomtom_export.gpx) 
and merged with all points from the Garmin File S1_12Ipertrail_2022-v221227.GPX into a new file (see .\data\redux_S1.gpx). 
The result is the file redux_S1.gpx and is imported in TomTom without any errors.
But at the end it seems that the problem is the filename extension. It should be .gpx instead of .GPX.


## Usage
Checkout this repository and build the project with golang (version used is 1.16.4)

    go build

Execute in powershell

    .\gpxtidy.exe -cmd remext -redfactor 2  -dir .\data -source S1_12Ipertrail_2022-v221227.GPX  -target redux_S1.gpx
