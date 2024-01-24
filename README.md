# Kamel - iRacing IMSA Vintage AI roster generator

Herein lies some lashed up code to generate an AI roster for the iRacing IMSA Vintage series (previously known as Kamel)

The generated roster weights the skill levels of each AI driver based on their actual results over seasons in 2023. The roster
uses the names and any custom liveries of the drivers that completed in the Kamel series.

There is some randomization of Skill, Optimism, Smoothness and driver age.

## Quick start

Download the prebuilt Kamel.zip file from here,

https://drive.google.com/file/d/1cm6EUkNQs26TBzn3H_6hHr0KlqXxlOgT/view?usp=sharing


It's a 256Mb download, but it includes the top 100 Kamel GT drivers who have completed in the VCR and VCR Junior championships over the last three seasons 

All the paints, suits, helmets etc. are included.

Unzip the file above into Documents/iRacing/airosters/ and in there there is a roster.json file.

Now start an AI session, choose the roster in your folder and boom. 

![AI Roster](/screenshots/ai-race.png "AI Race")

In this example, I unzipped into `Documents/iRacing/airosters/Kamel` and looks a bit like this,

![AI folder](/screenshots/ai-folder.png "AI Folder")

The `car_spec_*` files are custom liveries for the drivers that were running the series at the time.

## Customizing

The zip file above will be out of date as drivers come and go, and liveries change.

There are few different ways to modify the roster and adjust skill levels, driver names etc.

### Via the iRacing UI

Either download the ZIP above, or use the `roster.json` from this repo and copy into a new folder in the `Documents/iRacing/airosters` directory. For example, `Documents/iRacing/airosters/imsa`

Using the iRacing UI go to **AI Racing** -> **Opponent Rosters** and you should see your **IMSA** roster.

https://www.iracing.com/ai-rosters/ explains how to use the UI to configure your AI drivers.

### Manually edits

Use your favourite text editor to make changes to the `roster.json` file.

Each of the AI attributes are explained in the section *How AI Driver Attributes Affect Racing*

### Liveries

Within `roster.json` we can specify a custom paint job.

```
...
      "driverName": "Ian Haycox",
      "carNumber": "8",
      "carDesign": "18,66d945,ed9321,000000",
      "carTgaName": "car_203536.tga",
...
```

When the AI race loads it uses the `car_203536.tga` file for the custom livery. This file should be in the same folder as the AI roster.

The file can be any name, but to tie up the liveries with the driver names I copied all the TGA files from the `Documents\iRacing\paint\audi90gto` and `Documents\iRacing\paint\nissangtpzxt` folders. These files are left behind by
[Trading Paints](https://www.tradingpaints.com/)

There is also an optional related `car_spec_203536` file used for the shiny bits.

So to add a new driver, e,g, 'Joe Slow', edit the roster to add their name, then find the livery files in the appropriate iRacing paint folder. Copy the TGA and the `carTgaName` attribute to the file name. Hint: The number is the users' iRacing customer id. You can find this by searching in the UI or looking in the Trading Paints or Crew Chief log files.

### Regenerate the roster

The generator `main.go` uses two files as input to determine AI skill levels and generate the `roster.json` file.

- drivers.csv - https://vcr.myleague.racing/seasons/60
- paints.json - Trading Paints

The `drivers.csv` file contains a list of results for the Nissan and Audi standing for three seasons. I grabbed this data from https://vcr.myleague.racing/seasons/60 and other seasons and just concatenated the files together.

The only thing we are really interested in is the Driver, Class and Points fields. Basically the higher the points the more likely is the driver to be rated as a skilful AI opponent.

The `paints.json` I concocted from the Trading Paints log files. This provides the default paint schemes for the Car, Suit and Helmet, if there is no TGA file.

Running

`go run main.go`

will generate a new `roster.json` file which you should copy to you iRacing airosters folder.

Oh, and BTW - Don't be offended by your Age or Skill Levels in the Roster - blame the random number generator.

## Warranty

No docs, warranty, or guarantees, expressed or implied.

